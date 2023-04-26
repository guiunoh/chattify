package bootstrap

import (
	"consumer/internal/forwarder"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func Setup(rdb *redis.Client, httpclient *http.Client, url, channel string) {
	setupForwarder(rdb, httpclient, url, channel)
}

func setupForwarder(rdb *redis.Client, httpclient *http.Client, url, channel string) {
	gateway := forwarder.NewGateway(httpclient, url)
	usecase := forwarder.NewUsecase(gateway)
	handler := forwarder.NewHandler(usecase, rdb, httpclient, channel)
	go handler.RunLoop()
}
