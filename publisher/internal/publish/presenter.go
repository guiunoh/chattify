package publish

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Presenter interface {
	Created(c *gin.Context)
	BadRequest(c *gin.Context, err error)
	InternalServerError(c *gin.Context, err error)
}

func NewPresenter() Presenter {
	return &presenter{}
}

type presenter struct {
}

func (p presenter) Created(c *gin.Context) {
	c.Status(http.StatusCreated)
}

func (p presenter) NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func (p presenter) BadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(
		http.StatusBadRequest,
		gin.H{
			"err": err.Error(),
		},
	)
}

func (p presenter) InternalServerError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		gin.H{
			"err": err.Error(),
		},
	)
}
