package main

import (
	"log"
	"notifier/cmd/notifier/bootstrap"
	_gin "notifier/infrastructure/gin"
	_redis "notifier/infrastructure/redis"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	rdb := _redis.NewClient(bootstrap.Config.RDB)
	defer rdb.Close()

	engine := _gin.NewEngine()
	engine.Use(_gin.MiddlewarePing("/ping"))

	bootstrap.Setup(engine, rdb)

	_gin.Serve(engine, bootstrap.Config.Service.Port)
}
