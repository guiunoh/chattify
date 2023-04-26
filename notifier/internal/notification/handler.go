package notification

import (
	"log"
	"notifier/internal/connection/adaptor"

	"github.com/gin-gonic/gin"
)

func NewHandler(hub adaptor.Hub) *Handler {
	return &Handler{
		hub:       hub,
		presenter: NewPresenter(),
	}
}

type Handler struct {
	hub       adaptor.Hub
	presenter Presenter
}

func (h Handler) Route(r gin.IRoutes) {
	r.POST("/notification", h.Post)
}

func (h Handler) Post(c *gin.Context) {
	var input struct {
		ConnectionID string `json:"connectionID" binding:"required"`
		Payload      string `json:"payload" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		h.presenter.BadRequest(c, err)
		return
	}

	connector := h.hub.GetConnector(input.ConnectionID)
	if connector == nil {
		h.presenter.BadRequest(c, ErrNotFound)
		return
	}
	connector.SendMessage(input.Payload)
	h.presenter.Created(c)
}
