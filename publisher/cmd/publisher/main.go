package main

import (
	"log"
	"publisher/cmd/publisher/bootstrap"
	_gin "publisher/infrastructure/gin"
	_mysql "publisher/infrastructure/mysql"
	_redis "publisher/infrastructure/redis"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	db := _mysql.NewDB(bootstrap.Config.DB)
	defer _mysql.CloseSQL(db)

	rdb := _redis.NewClient(bootstrap.Config.RDB)
	defer rdb.Close()

	engine := _gin.NewEngine()
	engine.Use(_gin.MiddlewarePing("/ping"))

	bootstrap.Setup(engine, db, rdb)

	_gin.Serve(engine, bootstrap.Config.Service.Port)
}
