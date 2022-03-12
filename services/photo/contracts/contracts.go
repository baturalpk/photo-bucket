package contracts

import (
	"github.com/google/uuid"
	"image"
)

type NewPhoto struct {
	OwnerID     uuid.UUID   `json:"owner_id"`
	Tags        []string    `json:"tags"`
	Description string      `json:"description"`
	Data        image.Image `json:"data"`
	Format      string      `json:"format"`
}
