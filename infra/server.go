package infra

import (
	"context"
	"github.com/developerserasiautoraya/fms-dm-attendance/controller"
)

func RunServer() error {
	ctx := context.Background()

	gService := controller.NewAbsenceServiceServer()
	go func() {
		_ = RunREST(ctx)
	}()

	return RunGRPC(ctx, gService)
}
