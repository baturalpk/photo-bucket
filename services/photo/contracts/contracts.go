package contracts

import (
	"image"

	"github.com/google/uuid"
)

type NewPhoto struct {
	OwnerID     uuid.UUID `json:"owner_id" validate:"required"`
	Tags        []string  `json:"tags"`
	Description string    `json:"description"`
	Data        image.Image
	Format      string
}
