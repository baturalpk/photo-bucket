package s3client

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"testing"

	"github.com/baturalpk/photo-bucket/test/data"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
)

var (
	uniqueSubDir = uuid.New().String()
)

func TestInitConnection(t *testing.T) {
	if err := InitConnection(); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestUploadPhoto(t *testing.T) {
	s3 := S3{}
	file, err := os.Open(filepath.Join(data.PhotosTestDir, data.TestPhotoName))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err := s3.UploadPhoto(fmt.Sprintf("test/%s/%s", uniqueSubDir, data.TestPhotoName), img); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestRetrievePhoto(t *testing.T) {
	s3 := S3{}
	img, err := s3.RetrievePhoto(fmt.Sprintf("test/%s/%s", uniqueSubDir, data.TestPhotoName))
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	file, err := os.Create(filepath.Join(data.PhotosTestDir, fmt.Sprintf("%s_result_%s.jpg", data.TestPhotoName, uniqueSubDir)))
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if err = jpeg.Encode(file, img, &jpeg.Options{Quality: PhotosQuality}); err != nil {
		t.Error(err)
		t.Fail()
	}
}
