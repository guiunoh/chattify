package subscribe

import (
	"log"
	"publisher/pkg/ulid"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Route(r gin.IRoutes)
	post(c *gin.Context)
	Delete(c *gin.Context)
}

func NewHandler(u Usecase, p Presenter) Handler {
	return &handler{
		usecase:   u,
		presenter: p,
	}
}

type handler struct {
	usecase   Usecase
	presenter Presenter
}

func (h handler) Route(r gin.IRoutes) {
	r.POST("/subscribes", h.post)
	r.DELETE("/subscribes", h.Delete)

}

func (h handler) post(c *gin.Context) {
	var input struct {
		ConnectionID string `form:"connectionID" binding:"required"`
		Topic        string `form:"topic" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	connectionID, err := ulid.ParseID(input.ConnectionID)
	if err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	log.Println("input:", input)
	if err := h.usecase.Subscribe(c, connectionID, input.Topic); err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}

	h.presenter.Created(c)
}

func (h handler) Delete(c *gin.Context) {
	var input struct {
		ConnectionID string `form:"connectionID" binding:"required"`
		Topic        string `form:"topic"`
	}

	if err := c.ShouldBind(&input); err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	connectionID, err := ulid.ParseID(input.ConnectionID)
	if err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	log.Println("input:", input)
	if err := h.usecase.UnSubscribe(c, connectionID, input.Topic); err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}

	h.presenter.DeleteOneOK(c)
}
