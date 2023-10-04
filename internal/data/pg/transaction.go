package pg

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/rarimo/broadcaster-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (q TransactionQ) Select(ctx context.Context) ([]data.Transaction, error) {
	stmt := squirrel.Select(colsTransaction).From("transactions")
	var result []data.Transaction
	err := q.db.SelectContext(ctx, &result, stmt)
	return result, errors.Wrap(err, "failed to exec stmt")
}
