package infra

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"github.com/wulungtry/techtalk/api/proto/person"
	"github.com/wulungtry/techtalk/common"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunREST(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	restPort := viper.GetString("rest.port")
	grpcAddr := viper.GetString("grpc.addr")
	grpcPort := viper.GetString("grpc.port")
	endpoint := fmt.Sprintf("%s:%s", grpcAddr, grpcPort)

	if err := personal.RegisterPersonalServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatalf(common.FailToStartServer("HTTP Gateway"))
	}

	server := &http.Server{
		Addr:    ":" + restPort,
		Handler: mux,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			server.Shutdown(ctx)
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		_ = server.Shutdown(ctx)
	}()

	log.Println(common.StartRESTServer())
	return server.ListenAndServe()
}
