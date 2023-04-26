package message

import (
	"log"
	"publisher/pkg/ulid"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Route(r gin.IRoutes)
	GetOne(c *gin.Context)
	GetMany(c *gin.Context)
	Post(c *gin.Context)
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
	r.GET("/messages/:id", h.GetOne)
	r.GET("/messages", h.GetMany)
	r.POST("/messages", h.Post)
}

func (h handler) GetOne(c *gin.Context) {
	var input struct {
		ID    string `uri:"id" binding:"required"`
		Topic string `form:"topic" binding:"required"`
	}

	_ = c.ShouldBindUri(&input)
	if err := c.ShouldBind(&input); err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	id, err := ulid.ParseID(input.ID)
	if err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	log.Println("GetOne", input)

	messages, err := h.usecase.GetMessageOne(c.Request.Context(), id, input.Topic)
	if err != nil {
		h.presenter.InternalServerError(c, err)
	}

	h.presenter.GetOnOK(c, messages)
}

func (h handler) GetMany(c *gin.Context) {
	var input struct {
		LastID  string `form:"lastID" binding:"required"`
		Topic   string `form:"topic" binding:"required"`
		Limit   int    `form:"limit" binding:"required"`
		Forward bool   `form:"forward"`
	}

	_ = c.ShouldBindUri(&input)
	if err := c.ShouldBind(&input); err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	lastID, err := ulid.ParseID(input.LastID)
	if err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	topic := input.Topic
	limit := input.Limit
	forward := input.Forward

	log.Println("getMany", input)

	messages, err := h.usecase.GetMessageMany(c.Request.Context(), lastID, topic, limit, forward)
	if err != nil {
		h.presenter.InternalServerError(c, err)
	}

	h.presenter.GetManyOK(c, messages)
}

func (h handler) Post(c *gin.Context) {
	var input struct {
		Topic   string `json:"topic" binding:"required"`
		Content string `json:"content" binding:"required"`
		Author  string `json:"author" binding:"required"`
	}

	_ = c.ShouldBindUri(&input)
	if err := c.ShouldBind(&input); err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	res, err := h.usecase.CreateMessage(c.Request.Context(), input.Topic, input.Content, input.Author)
	if err != nil {
		h.presenter.InternalServerError(c, err)
	}

	h.presenter.Created(c, res.ID.String(), res.ExpiryAt)
}
