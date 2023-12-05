package broadcaster

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/distributed_lab/logan/v3"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types"
	client "github.com/cosmos/cosmos-sdk/types/tx"
	txclient "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
	"github.com/rarimo/broadcaster-svc/internal/config"
	"github.com/rarimo/broadcaster-svc/internal/data"
	"github.com/rarimo/rarimo-core/app"
	ethermint "github.com/rarimo/rarimo-core/ethermint/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

const (
	SuccessTxCode = 0
)

type broadcaster struct {
	log      *logan.Entry
	key      *config.KeyConf
	txConfig sdkclient.TxConfig
	txclient txclient.ServiceClient
	auth     authtypes.QueryClient
	txQ      data.TransactionQ
}

func Run(ctx context.Context, cfg config.Config) {
	log := cfg.Log().WithField("service", "broadcaster")

	svc := &broadcaster{
		log:      log,
		key:      cfg.Key(),
		txclient: client.NewServiceClient(cfg.Cosmos()),
		auth:     authtypes.NewQueryClient(cfg.Cosmos()),
		txQ:      cfg.Storage().TransactionQ(),
		txConfig: tx.NewTxConfig(
			codec.NewProtoCodec(codectypes.NewInterfaceRegistry()),
			[]signing.SignMode{signing.SignMode_SIGN_MODE_DIRECT},
		),
	}

	running.WithBackOff(ctx, svc.log, "run", svc.runOnceIndexing, 5*time.Second, 5*time.Second, 5*time.Second)
}

func (t *broadcaster) runOnceIndexing(ctx context.Context) error {
	txs, err := t.txQ.Select(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to select txs")
	}
	if len(txs) == 0 {
		t.log.Info("No txs to broadcast")
		return nil
	}

	t.log.Info("Got txs to broadcast", logan.F{
		"count": len(txs),
	})

	for _, txRaw := range txs {
		log := t.log.WithField("tx", txRaw.Data)

		tx, err := t.genTx(ctx, 0, 0, txRaw.Data)
		if err != nil {
			return errors.Wrap(err, "failed to generate tx")
		}

		gasUsed, err := t.simulateTx(ctx, tx)
		if err != nil {
			log.WithError(err).Error("Failed to simulate tx")

			if err = t.deleteTx(ctx, log, &txRaw); err != nil {
				return err
			}
		}

		gasLimit := ApproximateGasLimit(gasUsed)
		feeAmount := GetFeeAmount(gasLimit)

		t.log.Debugf("Gas limit: %d; Fee: %durmo", gasLimit, feeAmount)

		tx, err = t.genTx(ctx, gasLimit, feeAmount, txRaw.Data)
		if err != nil {
			return errors.Wrap(err, "failed to generate tx after simulation")
		}

		if err = t.broadcastTx(ctx, tx); err != nil {
			log.WithError(err).Error("Failed to broadcast tx")
		}

		if err = t.deleteTx(ctx, log, &txRaw); err != nil {
			return err
		}
	}

	t.log.Info("Finished broadcasting txs")
	return nil
}

func (t *broadcaster) deleteTx(ctx context.Context, log *logan.Entry, txRaw *data.Transaction) error {
	err := t.txQ.DeleteCtx(ctx, txRaw)
	if err != nil {
		log.WithError(err).Error("Failed to delete tx")
		return errors.Wrap(err, "failed to delete tx")
	}

	return nil
}

func (t *broadcaster) genTx(ctx context.Context, gasLimit, feeAmount uint64, rawTx []byte) ([]byte, error) {
	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	decodedTx, err := encCfg.TxConfig.TxDecoder()(rawTx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode tx")
	}

	builder, err := t.txConfig.WrapTxBuilder(decodedTx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to wrap tx with builder")
	}

	builder.SetGasLimit(gasLimit)
	builder.SetFeeAmount(types.Coins{types.NewInt64Coin(t.key.CoinName, int64(feeAmount))})

	accountResp, err := t.auth.Account(ctx, &authtypes.QueryAccountRequest{Address: t.key.SenderAddress})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get validator account")
	}

	account := ethermint.EthAccount{}
	err = account.Unmarshal(accountResp.Account.Value)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal validator account")
	}

	err = builder.SetSignatures(signing.SignatureV2{
		PubKey: t.key.Sender.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  t.txConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: account.Sequence,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to set signatures")
	}

	signerData := xauthsigning.SignerData{
		ChainID:       t.key.ChainId,
		AccountNumber: account.AccountNumber,
		Sequence:      account.Sequence,
	}

	sigV2, err := clienttx.SignWithPrivKey(
		t.txConfig.SignModeHandler().DefaultMode(), signerData,
		builder, t.key.Sender, t.txConfig, account.Sequence,
	)

	err = builder.SetSignatures(sigV2)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set signatures V2")
	}

	return t.txConfig.TxEncoder()(builder.GetTx())
}

func (t *broadcaster) simulateTx(ctx context.Context, tx []byte) (gasUsed uint64, err error) {
	simRes, err := t.txclient.Simulate(
		ctx,
		&client.SimulateRequest{
			Tx:      nil,
			TxBytes: tx,
		},
	)

	if err != nil {
		return 0, err
	}

	t.log.Debugf("Gas wanted estimation: %d; Gas used estimation: %d", simRes.GasInfo.GasWanted, simRes.GasInfo.GasUsed)
	return simRes.GasInfo.GasUsed, nil
}

func (t *broadcaster) broadcastTx(ctx context.Context, tx []byte) error {
	grpcRes, err := t.txclient.BroadcastTx(
		ctx,
		&client.BroadcastTxRequest{
			Mode:    client.BroadcastMode_BROADCAST_MODE_BLOCK,
			TxBytes: tx,
		},
	)
	if err != nil {
		return err
	}

	t.log.Debugf("Submitted transaction to the core: %s", grpcRes.TxResponse.TxHash)

	if grpcRes.TxResponse.Code != SuccessTxCode {
		return errors.New(fmt.Sprintf("Got error code: %d, info: %s, log: %s", grpcRes.TxResponse.Code, grpcRes.TxResponse.Info, grpcRes.TxResponse.RawLog))
	}

	return nil
}
