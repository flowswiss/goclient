package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type ImageService struct {
	client *core.Client
}

func NewImageService(client *core.Client) *ImageService {
	return &ImageService{client: client}
}

func (i ImageService) List(ctx context.Context, cursor core.Cursor) (list common.List[Image], err error) {
	list.Pagination, err = i.client.List(ctx, getImagesPath(), cursor, &list.Items)
	return
}

type ImageGetReq struct {
	ID uint `json:"-"`
}

func (i ImageService) Get(ctx context.Context, req ImageGetReq) (image Image, err error) {
	err = i.client.Get(ctx, getSpecificImagePath(req.ID), &image)
	return
}

const imagesSegment = "/v4/entities/compute/images"

func getImagesPath() string {
	return imagesSegment
}

func getSpecificImagePath(id uint) string {
	return core.Join(imagesSegment, id)
}
