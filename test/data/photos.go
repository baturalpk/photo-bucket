package data

import (
	"path/filepath"

	"github.com/baturalpk/photo-bucket/config"
)

const TestPhotoName = "1.jpg"

var PhotosTestDir = filepath.Join(config.TestDirPath(), "data", "photos")
