package infra

import (
	"context"
	"github.com/wulungtry/techtalk/controller"
)

func RunServer() error {
	ctx := context.Background()

	gService := controller.NewPersonalServiceServer()
	go func() {
		_ = RunREST(ctx)
	}()

	return RunGRPC(ctx, gService)
}
