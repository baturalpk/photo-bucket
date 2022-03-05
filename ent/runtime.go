// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/baturalpk/photo-bucket/ent/photosmetadata"
	"github.com/baturalpk/photo-bucket/ent/profile"
	"github.com/baturalpk/photo-bucket/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	photosmetadataFields := schema.PhotosMetadata{}.Fields()
	_ = photosmetadataFields
	// photosmetadataDescWidth is the schema descriptor for width field.
	photosmetadataDescWidth := photosmetadataFields[4].Descriptor()
	// photosmetadata.WidthValidator is a validator for the "width" field. It is called by the builders before save.
	photosmetadata.WidthValidator = photosmetadataDescWidth.Validators[0].(func(int) error)
	// photosmetadataDescHeight is the schema descriptor for height field.
	photosmetadataDescHeight := photosmetadataFields[5].Descriptor()
	// photosmetadata.HeightValidator is a validator for the "height" field. It is called by the builders before save.
	photosmetadata.HeightValidator = photosmetadataDescHeight.Validators[0].(func(int) error)
	// photosmetadataDescUploadedAt is the schema descriptor for uploaded_at field.
	photosmetadataDescUploadedAt := photosmetadataFields[8].Descriptor()
	// photosmetadata.DefaultUploadedAt holds the default value on creation for the uploaded_at field.
	photosmetadata.DefaultUploadedAt = photosmetadataDescUploadedAt.Default.(func() time.Time)
	profileFields := schema.Profile{}.Fields()
	_ = profileFields
	// profileDescUsername is the schema descriptor for username field.
	profileDescUsername := profileFields[1].Descriptor()
	// profile.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	profile.UsernameValidator = profileDescUsername.Validators[0].(func(string) error)
	// profileDescPasswordHash is the schema descriptor for password_hash field.
	profileDescPasswordHash := profileFields[2].Descriptor()
	// profile.PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	profile.PasswordHashValidator = func() func(string) error {
		validators := profileDescPasswordHash.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password_hash string) error {
			for _, fn := range fns {
				if err := fn(password_hash); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// profileDescName is the schema descriptor for name field.
	profileDescName := profileFields[4].Descriptor()
	// profile.DefaultName holds the default value on creation for the name field.
	profile.DefaultName = profileDescName.Default.(string)
	// profileDescBiography is the schema descriptor for biography field.
	profileDescBiography := profileFields[5].Descriptor()
	// profile.DefaultBiography holds the default value on creation for the biography field.
	profile.DefaultBiography = profileDescBiography.Default.(string)
	// profileDescIsEmailVerified is the schema descriptor for is_email_verified field.
	profileDescIsEmailVerified := profileFields[8].Descriptor()
	// profile.DefaultIsEmailVerified holds the default value on creation for the is_email_verified field.
	profile.DefaultIsEmailVerified = profileDescIsEmailVerified.Default.(bool)
	// profileDescIsPhoneVerified is the schema descriptor for is_phone_verified field.
	profileDescIsPhoneVerified := profileFields[9].Descriptor()
	// profile.DefaultIsPhoneVerified holds the default value on creation for the is_phone_verified field.
	profile.DefaultIsPhoneVerified = profileDescIsPhoneVerified.Default.(bool)
	// profileDescCreatedAt is the schema descriptor for created_at field.
	profileDescCreatedAt := profileFields[10].Descriptor()
	// profile.DefaultCreatedAt holds the default value on creation for the created_at field.
	profile.DefaultCreatedAt = profileDescCreatedAt.Default.(func() time.Time)
	// profileDescID is the schema descriptor for id field.
	profileDescID := profileFields[0].Descriptor()
	// profile.DefaultID holds the default value on creation for the id field.
	profile.DefaultID = profileDescID.Default.(func() uuid.UUID)
}
