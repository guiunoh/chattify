package thumbnail

import (
	"context"
	"image"
	"thumbnailer/pkg/resizer"
	"thumbnailer/pkg/ulid"
)

func NewInteractor(r Repository, resizer resizer.ImageResizer) Usecase {
	return &interactor{r, resizer}
}

type interactor struct {
	repo    Repository
	resizer resizer.ImageResizer
}

func (i interactor) CreateThumbnail(c context.Context, src image.Image, rate Rate) (*Thumbnail, error) {
	resized, imageType, err := i.resizer.Resize(src, rate.Value())
	if err != nil {
		return nil, err
	}

	entity := NewThumbnail(resized, imageType)
	if err := i.repo.Save(c, entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (i interactor) GetThumbnail(c context.Context, id ulid.ID) (*Thumbnail, error) {
	entity, err := i.repo.FindByID(c, id)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
