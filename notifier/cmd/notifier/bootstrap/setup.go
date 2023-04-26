package bootstrap

import (
	"notifier/internal/connection"
	"notifier/internal/connection/adaptor"
	"notifier/internal/notification"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Setup(r gin.IRouter, rdb *redis.Client) {
	v1 := r.Group("/api/v1")
	hub := adaptor.NewHub()
	setupConnection(v1, hub, rdb)
	setupNotification(v1, hub)

}

func setupConnection(r gin.IRouter, hub adaptor.Hub, rdb *redis.Client) {
	repository := connection.NewRepository(rdb, "connection")
	usecase := connection.NewInteractor(repository)
	connection.NewHandler(hub, usecase).Route(r)
}

func setupNotification(r gin.IRouter, hub adaptor.Hub) {
	notification.NewHandler(hub).Route(r)
}
