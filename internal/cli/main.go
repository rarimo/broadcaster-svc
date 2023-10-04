package cli

import (
	"context"
	"github.com/rarimo/broadcaster-svc/grpc"
	"github.com/rarimo/broadcaster-svc/internal/services/broadcaster"
	"gitlab.com/distributed_lab/logan/v3"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gitlab.com/distributed_lab/logan/v3/errors"

	"github.com/alecthomas/kingpin"
	"github.com/rarimo/broadcaster-svc/internal/config"
	"gitlab.com/distributed_lab/kit/kv"
)

func Run(args []string) bool {
	log := logan.New()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("app panicked")
		}
	}()

	app := kingpin.New("broadcaster-svc", "")

	runCmd := app.Command("run", "run command")
	allCmd := runCmd.Command("all", "run all services (broadcaster, grpc server)")

	apiCmd := runCmd.Command("api", "run service") // you can insert custom help

	runBroadcaster := runCmd.Command("broadcaster", "")

	migrateCmd := app.Command("migrate", "migrate command")
	migrateUpCmd := migrateCmd.Command("up", "migrate db up")
	migrateDownCmd := migrateCmd.Command("down", "migrate db down")

	cmd, err := app.Parse(args[1:])
	if err != nil {
		logan.New().WithError(err).Fatal("failed to parse arguments")
		return false
	}

	ctx, cancel := context.WithCancel(context.Background())
	cfg := config.New(kv.MustFromEnv())

	var wg sync.WaitGroup

	run := func(f func(ctx context.Context, cfg config.Config)) {
		wg.Add(1)
		go func() {
			defer wg.Done()

			defer func() {
				if r := recover(); r != nil {
					log.WithRecover(r).Fatal("one of the services panicked")
				}
			}()

			f(ctx, cfg)
		}()
	}

	runAll := func() {
		cfg.Log().Info("starting all services")
		run(broadcaster.Run)
		run(grpc.RunAPI)
	}

	switch cmd {
	case allCmd.FullCommand():
		runAll()
	case apiCmd.FullCommand():
		run(grpc.RunAPI)
	case runBroadcaster.FullCommand():
		run(broadcaster.Run)
	case migrateUpCmd.FullCommand():
		if err := MigrateUp(cfg); err != nil {
			panic(errors.Wrap(err, "failed to migrate up"))
		}
	case migrateDownCmd.FullCommand():
		if err := MigrateDown(cfg); err != nil {
			panic(errors.Wrap(err, "failed to migrate down"))
		}
	default:
		panic(errors.From(errors.New("unknown command"), logan.F{
			"raw_command": cmd,
		}))
	}

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	wgch := make(chan struct{})
	go func() {
		wg.Wait()
		close(wgch)
	}()

	select {
	case <-wgch:
		cfg.Log().Warn("all services stopped")
		cancel()
	case <-gracefulStop:
		cfg.Log().Info("received signal to stop")
		cancel()
		<-wgch
	}

	return true
}
