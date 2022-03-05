package photo

import (
	"context"
	"image"
	"os"
	"path/filepath"
	"testing"

	"github.com/baturalpk/photo-bucket/clients/entclient"
	"github.com/baturalpk/photo-bucket/clients/s3client"
	"github.com/baturalpk/photo-bucket/services/photo/contracts"
	"github.com/baturalpk/photo-bucket/services/profile"
	"github.com/baturalpk/photo-bucket/test/data"
)

func beforeTest(t *testing.T) {
	if err1, err2 := entclient.InitConnection(), s3client.InitConnection(); err1 != nil || err2 != nil {
		t.Error(err1, err2)
		t.FailNow()
	}
}

func TestCreate(t *testing.T) {
	beforeTest(t)
	profileRepo := profile.NewRepository(entclient.Client())
	p, err := profileRepo.GetByUsername(context.Background(), data.TestUsers()[0].Username)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	file, err := os.Open(filepath.Join(data.PhotosTestDir, data.TestPhotoName))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer file.Close()
	img, format, err := image.Decode(file)

	repo := NewRepository(entclient.Client(), s3client.S3{})
	repo.Create(context.Background(), contracts.NewPhoto{
		OwnerID:     p.ID,
		Tags:        []string{"nature", "forest", "hut"},
		Description: "this is an photo description",
		Data:        img,
		Format:      format,
	})
}
