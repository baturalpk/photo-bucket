package photo

import (
	"context"
	"fmt"

	"github.com/baturalpk/photo-bucket/clients/s3client"
	"github.com/baturalpk/photo-bucket/ent"
	"github.com/baturalpk/photo-bucket/ent/photosmetadata"
	"github.com/baturalpk/photo-bucket/services/photo/contracts"
	"github.com/google/uuid"
)

type repository struct {
	s3   s3client.S3
	meta *ent.Client
}

func NewRepository(metadataDB *ent.Client, mainStorage s3client.S3) *repository {
	return &repository{
		s3:   mainStorage,
		meta: metadataDB,
	}
}

func (r repository) Create(ctx context.Context, newp contracts.NewPhoto) (uuid.UUID, error) {
	pid := uuid.New()
	url := fmt.Sprintf("/%s/%s.%s", newp.OwnerID, pid.String(), newp.Format)

	if err := r.s3.UploadPhoto(url, newp.Data); err != nil {
		return pid, err
	}

	metad := r.meta.PhotosMetadata.Create()
	metad.
		SetID(pid).
		SetOwnerID(newp.OwnerID).
		SetTags(newp.Tags).
		SetDescription(newp.Description).
		SetWidth(newp.Data.Bounds().Dx()).
		SetHeight(newp.Data.Bounds().Dy()).
		SetImageFormat(photosmetadata.ImageFormat(newp.Format)).
		SetRelativeURL(url)

	if _, err := metad.Save(ctx); err != nil {
		return pid, err
	}
	return pid, nil
}

func (r repository) GetByID(ctx context.Context, id uuid.UUID) (*ent.PhotosMetadata, error) {
	return r.meta.PhotosMetadata.Get(ctx, id)
}

func (r repository) GetByOwner(ctx context.Context, oid uuid.UUID) ([]*ent.PhotosMetadata, error) {
	return r.meta.PhotosMetadata.Query().Where(photosmetadata.OwnerID(oid)).All(ctx)
}
