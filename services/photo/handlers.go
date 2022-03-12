package photo

import (
	"fmt"
	"image"
	"math"
	"net/http"

	"github.com/baturalpk/photo-bucket/services/photo/contracts"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type handlers struct {
	repo *repository
}

func NewHandlers(photoRepo *repository) *handlers {
	return &handlers{
		repo: photoRepo,
	}
}

func (h handlers) PostNewPhoto(c *fiber.Ctx) error {
	in := new(contracts.NewPhoto)
	if err := c.BodyParser(in); err != nil {
		return err
	}
	in.OwnerID = c.UserContext().Value("id").(uuid.UUID)
	photo, err := c.FormFile("photo")
	if err != nil {
		return err
	}
	if photo.Size > (30 << 20) { // 30 MB upload limit
		return c.Status(http.StatusRequestEntityTooLarge).SendString(
			fmt.Sprintf("photo size must not exceed %d MB",
				(30<<20)/int(math.Pow(10, 6))),
		)
	}
	photoFile, err := photo.Open()
	if err != nil {
		return err
	}
	img, format, err := image.Decode(photoFile)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("corrupted photo file")
	}
	in.Data = img
	in.Format = format
	_, err = h.repo.Create(c.UserContext(), *in)
	return err
}

func (h handlers) GetOneByID(c *fiber.Ctx) error {
	id := c.Params("id")
	pid, err := uuid.Parse(id)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}
	photo, err := h.repo.GetByID(c.UserContext(), pid)
	if err != nil {
		return err
	}
	return c.JSON(photo)
}

func (h handlers) GetAllByOwnerID(c *fiber.Ctx) error {
	oid := c.UserContext().Value("id").(uuid.UUID)
	photos, err := h.repo.GetByOwner(c.UserContext(), oid)
	if err != nil {
		return err
	}
	return c.JSON(photos)
}
