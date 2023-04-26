package bootstrap

import (
	"thumbnailer/internal/thumbnail"
	"thumbnailer/pkg/resizer"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Setup(r gin.IRouter, rdb *redis.Client) {
	v1 := r.Group("/api/v1")
	setupThumbnail(v1, rdb)
}

func setupThumbnail(r gin.IRouter, rdb *redis.Client) {
	repository := thumbnail.NewRepository(rdb)
	usecase := thumbnail.NewInteractor(repository, resizer.NewImageResizer())
	presenter := thumbnail.NewPresenter()
	thumbnail.NewHandler(usecase, presenter).Route(r)
}
