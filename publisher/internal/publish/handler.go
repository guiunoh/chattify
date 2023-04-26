package publish

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Route(r gin.IRoutes)
	post(c *gin.Context)
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
	r.POST("/publish", h.post)
}

func (h handler) post(c *gin.Context) {
	var input struct {
		Topic   string `json:"topic" binding:"required"`
		Content string `json:"content" binding:"required"`
		Sender  string `json:"sender" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	log.Println("input:", input)

	if err := h.usecase.PostToTopic(c.Request.Context(), input.Topic, input.Sender, input.Content); err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}

	h.presenter.Created(c)
}
