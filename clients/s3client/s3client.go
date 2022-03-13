package s3client

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/baturalpk/photo-bucket/config"
)

const (
	PhotosSupportedFormat = "jpeg"
	PhotosQuality         = 80
)

var (
	s3Session    *session.Session
	photosBucket string
	originServer string
)

type S3 struct{}

func upload(bucket, key string, buffer []byte) error {
	c := s3.New(s3Session)
	_, err := c.PutObject(&s3.PutObjectInput{
		Body:          bytes.NewReader(buffer),
		Bucket:        aws.String(bucket),
		Key:           aws.String(key),
		ContentLength: aws.Int64(int64(len(buffer))),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})
	return err
}

func retrieve(bucket, key string) (io.ReadCloser, error) {
	c := s3.New(s3Session)
	if res, err := c.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}); err != nil {
		return nil, err
	} else {
		return res.Body, nil
	}
}

func (S3) UploadPhoto(path string, image image.Image) error {
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, image, &jpeg.Options{Quality: PhotosQuality}); err != nil {
		return err
	}
	return upload(photosBucket, path, buf.Bytes())
}

func (S3) RetrievePhoto(path string) (image.Image, error) {
	r, err := retrieve(photosBucket, path)
	if err != nil {
		return nil, err
	}

	img, err := jpeg.Decode(r)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func InitConnection() error {
	switch config.Get().Env {
	case "prod":
		photosBucket = config.Get().S3.BucketNames.Photos.Prod
	case "dev":
		photosBucket = config.Get().S3.BucketNames.Photos.Dev
	case "test":
		photosBucket = config.Get().S3.BucketNames.Photos.Test
	default:
		return errors.New("GO_ENV is not defined properly")
	}

	originServer = fmt.Sprintf("%s/%s", config.Get().S3.Endpoint, photosBucket)
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(config.Get().S3.Credentials.Id, config.Get().S3.Credentials.Secret, ""),
		Endpoint:         aws.String(config.Get().S3.Endpoint),
		Region:           aws.String(config.Get().S3.Region),
		S3ForcePathStyle: aws.Bool(true),
	}

	log.Println("trying to establish S3 connection...")
	if newSession, err := session.NewSession(s3Config); err != nil {
		return err
	} else {
		s3Session = newSession
	}

	log.Println("successfully created a new S3 session")
	return nil
}

func (S3) OriginServer() string {
	return originServer
}
