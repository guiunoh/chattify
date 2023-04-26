package subscribe

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Presenter interface {
	GetOneOK(c *gin.Context, data interface{})
	GetManyOK(c *gin.Context, data interface{})
	DeleteOneOK(c *gin.Context)
	DeleteManyOK(c *gin.Context)
	Created(c *gin.Context)
	NoContent(c *gin.Context)
	BadRequest(c *gin.Context, err error)
	InternalServerError(c *gin.Context, err error)
}

func NewPresenter() Presenter {
	return &presenter{}
}

type presenter struct {
}

func (p presenter) GetOneOK(c *gin.Context, data interface{}) {
	//TODO implement me
	panic("implement me")
}

func (p presenter) GetManyOK(c *gin.Context, data interface{}) {
	//TODO implement me
	panic("implement me")
}

func (p presenter) DeleteOneOK(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func (p presenter) DeleteManyOK(c *gin.Context) {
	c.Status(http.StatusNoContent)
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
