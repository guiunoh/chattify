package message

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Presenter interface {
	GetOnOK(c *gin.Context, data interface{})
	GetManyOK(c *gin.Context, data interface{})
	Created(c *gin.Context, id string, expiry time.Time)
	NoContent(c *gin.Context)
	BadRequest(c *gin.Context, err error)
	InternalServerError(c *gin.Context, err error)
}

func NewPresenter() Presenter {
	return &presenter{}
}

type presenter struct {
}

func (p presenter) GetOnOK(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"data": data,
		},
	)
}

func (p presenter) GetManyOK(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"data": data,
		},
	)
}

func (p presenter) Created(c *gin.Context, id string, expiry time.Time) {
	if expiry.After(time.Now()) {
		maxAge := int(time.Until(expiry).Seconds())
		c.Writer.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", maxAge))
	}

	url := fmt.Sprintf("%s/%s", c.Request.RequestURI, id)
	c.Writer.Header().Set("Location", url)

	c.JSON(
		http.StatusCreated,
		gin.H{
			"id": id,
		},
	)
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
