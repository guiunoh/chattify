package main

import (
	"log"
	"thumbnailer/cmd/thumbnailer/bootstrap"
	_gin "thumbnailer/infrastructure/gin"
	_redis "thumbnailer/infrastructure/redis"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	rdb := _redis.NewClient(bootstrap.Config.RDB)
	defer rdb.Close()

	engine := _gin.NewEngine()
	engine.MaxMultipartMemory = 4 << 20 // 4 MiB
	engine.Use(_gin.MiddlewarePing("/ping"))

	bootstrap.Setup(engine, rdb)

	_gin.Serve(engine, bootstrap.Config.Service.Port)
}
