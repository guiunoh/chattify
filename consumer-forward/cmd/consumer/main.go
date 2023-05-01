package main

import (
	"consumer-forward/cmd/consumer/bootstrap"
	"consumer-forward/pkg/httpclient"
	"context"
	"log"
	"os/signal"
	"syscall"

	_redis "consumer-forward/infrastructure/redis"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	rdb := _redis.NewClient(bootstrap.Config.RDB)
	defer rdb.Close()
	client := httpclient.DefaultHttpClient()
	bootstrap.Setup(rdb, client, bootstrap.Config.Forward.Endpoint, bootstrap.Config.Service.Channel)
	shutdown()
}

func shutdown() {
	log.Println("start consumer:", bootstrap.Config.Service.Channel)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	log.Println("exiting")
}
