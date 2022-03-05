package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("username").NotEmpty().Unique(),
		field.String("password_hash").NotEmpty().Sensitive().MinLen(8),
		field.String("picture_url").Optional(),
		field.String("name").Default(""),
		field.String("biography").Default(""),
		field.String("email").Unique().Optional(),
		field.String("phone").Unique().Optional(),
		field.Bool("is_email_verified").Default(false),
		field.Bool("is_phone_verified").Default(false),
		field.Time("created_at").Default(func() time.Time { return time.Now() }),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("photos", PhotosMetadata.Type),
	}
}
