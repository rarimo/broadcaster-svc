package grpc

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"net"

	"google.golang.org/grpc"
)

type ServeConfig interface {
	Listener() net.Listener
	Log() *logan.Entry
}

func serve(ctx context.Context, server *grpc.Server, serveConfig ServeConfig) {
	done := make(chan struct{})

	go func() {
		defer close(done)
		defer func() {
			if r := recover(); r != nil {
				serveConfig.Log().WithRecover(r).Fatal("grpc service panicked")
			}
		}()

		err := server.Serve(serveConfig.Listener())
		if err == grpc.ErrServerStopped {
			serveConfig.Log().Info("stopped accepting new connections")
			return
		}

		if err != nil {
			serveConfig.Log().WithError(err).Error("server died")
		}
	}()

	select {
	case <-done:
		// error is already logged
		return
	case <-ctx.Done():
		serveConfig.Log().Info("received signal to shutdown gracefully")
		server.GracefulStop()
		serveConfig.Log().Info("shutdown gracefully")
	}
}
