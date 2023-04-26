package thumbnail

import (
	"image"
	"mime/multipart"
	"thumbnailer/pkg/ulid"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Route(r gin.IRoutes)
	GetOne(c *gin.Context)
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
	r.GET("/thumbnails/:id", h.GetOne)
	r.POST("/thumbnails", h.post)
}

func (h handler) GetOne(c *gin.Context) {
	var input struct {
		ID string `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&input); err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	id, err := ulid.ParseID(input.ID)
	if err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}

	output, err := h.usecase.GetThumbnail(c.Request.Context(), id)
	if err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}

	h.presenter.GetOnOK(c, *output)
}

func (h handler) post(c *gin.Context) {
	var input struct {
		Rate string                `form:"rate" binding:"required"`
		File *multipart.FileHeader `form:"file" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	rate, err := StringToRate(input.Rate)
	if err != nil {
		h.presenter.BadRequest(c, err)
		return
	}

	file := input.File
	imageFile, err := file.Open()
	if err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}
	defer imageFile.Close()

	src, _, err := image.Decode(imageFile)
	if err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}

	output, err := h.usecase.CreateThumbnail(c.Request.Context(), src, rate)
	if err != nil {
		h.presenter.InternalServerError(c, err)
		return
	}

	h.presenter.Created(c, *output)
}
