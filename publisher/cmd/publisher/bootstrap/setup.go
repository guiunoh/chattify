package bootstrap

import (
	"publisher/internal/message"
	"publisher/internal/publish"
	"publisher/internal/subscribe"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Setup(r gin.IRouter, db *gorm.DB, rdb *redis.Client) {
	v1 := r.Group("/api/v1")
	setupMessage(v1, db)
	setupSubscribe(v1, db)
	setupPublish(v1, db, rdb)
}

func setupMessage(r gin.IRouter, db *gorm.DB) {
	repository := message.NewRepository(db)
	usecase := message.NewInteractor(repository)
	presenter := message.NewPresenter()
	message.NewHandler(usecase, presenter).Route(r)
}

func setupSubscribe(r gin.IRouter, db *gorm.DB) {
	repository := subscribe.NewRepository(db)
	usecase := subscribe.NewInteractor(repository)
	presenter := subscribe.NewPresenter()
	subscribe.NewHandler(usecase, presenter).Route(r)
}

func setupPublish(r gin.IRouter, db *gorm.DB, rdb *redis.Client) {
	reader := subscribe.NewInteractor(subscribe.NewRepository(db))
	writer := message.NewInteractor(message.NewRepository(db))

	repository := publish.NewRepository(rdb, "publish-channel")
	usecase := publish.NewInteractor(repository, reader, writer)
	presenter := publish.NewPresenter()
	publish.NewHandler(usecase, presenter).Route(r)
}
