package thumbnail

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Presenter interface {
	GetOnOK(c *gin.Context, t Thumbnail)
	GetManyOK(c *gin.Context, t Thumbnail)
	Created(c *gin.Context, t Thumbnail)
	BadRequest(c *gin.Context, err error)
	InternalServerError(c *gin.Context, err error)
}

func NewPresenter() Presenter {
	return &presenter{}
}

type presenter struct {
}

func (p presenter) GetOnOK(c *gin.Context, t Thumbnail) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"result": t,
		},
	)
}

func (p presenter) GetManyOK(c *gin.Context, t Thumbnail) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"result": t,
		},
	)
}

func (p presenter) Created(c *gin.Context, t Thumbnail) {
	if t.Expiry.After(time.Now()) {
		maxAge := int(time.Until(t.Expiry).Seconds())
		c.Writer.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", maxAge))
	}

	url := fmt.Sprintf("%s/%s", c.Request.RequestURI, t.ID)
	c.Writer.Header().Set("Location", url)

	c.JSON(
		http.StatusCreated,
		gin.H{
			"id": t.ID,
		},
	)
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
