package photo

import (
	"context"
	"fmt"

	"github.com/baturalpk/photo-bucket/clients/s3client"
	"github.com/baturalpk/photo-bucket/ent"
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

func (r repository) Create(ctx context.Context, newp contracts.NewPhoto) error {
	pid := uuid.New()
	if err := r.s3.UploadPhoto(
		fmt.Sprintf("/%s/%s.%s", newp.OwnerID, pid.String(), newp.Format),
		newp.Data,
	); err != nil {
		return err
	}

	// TODO: ent. /create hook for uploading metadata
	return nil
}
