package infra

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/wulungtry/techtalk/controller"
)

func RunServer(db *gorm.DB) error {
	ctx := context.Background()

	gService := controller.NewPersonalServiceServer(db)
	go func() {
		_ = RunREST(ctx)
	}()

	return RunGRPC(ctx, gService)
}
