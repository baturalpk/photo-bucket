package api

import (
	"fmt"
	"log"
	"math"
	"net"

	"github.com/baturalpk/photo-bucket/api/middlewares"
	"github.com/baturalpk/photo-bucket/clients/entclient"
	"github.com/baturalpk/photo-bucket/clients/s3client"
	"github.com/baturalpk/photo-bucket/services/photo"
	"github.com/baturalpk/photo-bucket/services/profile"
	"github.com/gofiber/fiber/v2"
)

func isValidPort(port int) bool {
	return port >= 0 && port <= math.MaxUint16
}

func registerHandlers(app *fiber.App) {
	// Photos
	photoHandlers := photo.NewHandlers(photo.NewRepository(entclient.Client(), s3client.S3{}))
	photos := app.Group("/photos")
	photos.Get("/:id", photoHandlers.GetOneByID)
	photos.Get("/owner/:id", middlewares.AuthorizationGuard, photoHandlers.GetAllByOwnerID)
	photos.Post("/", middlewares.AuthorizationGuard, photoHandlers.PostNewPhoto)

	// Profiles
	profileHandlers := profile.NewHandlers(profile.NewRepository(entclient.Client()))
	profiles := app.Group("/profiles")
	profiles.Get("/:username", profileHandlers.GetByUsername)
	profiles.Post("/signup/email", profileHandlers.SignUpViaEmail)
	profiles.Post("/signup/phone", profileHandlers.SignUpViaPhone)
	profiles.Post("/signin", profileHandlers.SignIn)
}

func Serve(port ...int) error {
	app := fiber.New()
	registerHandlers(app)

	var address = ":0"
	if len(port) > 0 && isValidPort(port[0]) {
		address = fmt.Sprintf(":%d", port[0])
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	log.Printf("Server is listening at: %s\n", listener.Addr().String())
	return app.Listener(listener)
}
