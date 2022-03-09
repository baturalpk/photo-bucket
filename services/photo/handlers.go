package photo

import (
	"net/http"

	"github.com/baturalpk/photo-bucket/services/photo/contracts"
	"github.com/baturalpk/photo-bucket/utils/validator"
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

func (h handlers) PostNewPhoto(c *fiber.Ctx) error { // TODO: ! Handle image data extraction explicitly, set format, etc...
	in := new(contracts.NewPhoto)
	if err := c.BodyParser(in); err != nil {
		return err
	}
	if err := validator.GetValidatorInstance().Struct(in); err != nil {
		return err
	}
	_, err := h.repo.Create(c.UserContext(), *in)
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
	id := c.Params("id")
	oid, err := uuid.Parse(id)
	if err != nil {
		return c.SendStatus(http.StatusFound)
	}
	photos, err := h.repo.GetByOwner(c.UserContext(), oid)
	if err != nil {
		return err
	}
	return c.JSON(photos)
}
