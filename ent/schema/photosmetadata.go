package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PhotosMetadata holds the schema definition for the PhotosMetadata entity.
type PhotosMetadata struct {
	ent.Schema
}

// Fields of the PhotosMetadata.
func (PhotosMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique().Immutable(),
		field.UUID("owner_id", uuid.UUID{}).Optional(),
		field.Strings("tags").Optional(),
		field.String("description").Optional(),
		field.Int("width").NonNegative().Immutable(),
		field.Int("height").NonNegative().Immutable(),
		field.Enum("image_format").Values("jpeg").Immutable(),
		// TODO: Add field 'origin_server' to indicate where relative_url is exactly pointing.
		// 		 e.g., "s3.amazonaws.com", or "cdn.example.com/images", ...
		//		 ! Load from repository configurations & Set() while you are creating a new photo.
		field.String("relative_url").Immutable(),
		field.Time("uploaded_at").Default(func() time.Time { return time.Now() }),
	}
}

// Edges of the PhotosMetadata.
func (PhotosMetadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Profile.Type).
			Ref("photos").
			Unique().
			Field("owner_id"),
	}
}
