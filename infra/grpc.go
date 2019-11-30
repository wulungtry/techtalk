package infra

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/wulungtry/techtalk/api/proto/person"
	"github.com/wulungtry/techtalk/common"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

type Handler struct {
	Person personal.PersonalServiceServer
}

func RunGRPC(ctx context.Context, pers personal.PersonalServiceServer) error {
	port := viper.GetString("grpc.port")
	address := fmt.Sprintf(":%s", port)
	listen, err := net.Listen(common.TCP, address)

	if err != nil {
		return nil
	}

	server := grpc.NewServer()
	personal.RegisterPersonalServiceServer(server, pers)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println(common.StartGRPCServer())
	return server.Serve(listen)
}
