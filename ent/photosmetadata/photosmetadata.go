// Code generated by entc, DO NOT EDIT.

package photosmetadata

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the photosmetadata type in the database.
	Label = "photos_metadata"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldWidth holds the string denoting the width field in the database.
	FieldWidth = "width"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldImageFormat holds the string denoting the image_format field in the database.
	FieldImageFormat = "image_format"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldUploadedAt holds the string denoting the uploaded_at field in the database.
	FieldUploadedAt = "uploaded_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the photosmetadata in the database.
	Table = "photos_metadata"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "photos_metadata"
	// OwnerInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	OwnerInverseTable = "profiles"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
)

// Columns holds all SQL columns for photosmetadata fields.
var Columns = []string{
	FieldID,
	FieldOwnerID,
	FieldTags,
	FieldDescription,
	FieldWidth,
	FieldHeight,
	FieldImageFormat,
	FieldURL,
	FieldUploadedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// WidthValidator is a validator for the "width" field. It is called by the builders before save.
	WidthValidator func(int) error
	// HeightValidator is a validator for the "height" field. It is called by the builders before save.
	HeightValidator func(int) error
	// DefaultUploadedAt holds the default value on creation for the "uploaded_at" field.
	DefaultUploadedAt func() time.Time
)

// ImageFormat defines the type for the "image_format" enum field.
type ImageFormat string

// ImageFormat values.
const (
	ImageFormatJpeg ImageFormat = "jpeg"
)

func (_if ImageFormat) String() string {
	return string(_if)
}

// ImageFormatValidator is a validator for the "image_format" field enum values. It is called by the builders before save.
func ImageFormatValidator(_if ImageFormat) error {
	switch _if {
	case ImageFormatJpeg:
		return nil
	default:
		return fmt.Errorf("photosmetadata: invalid enum value for image_format field: %q", _if)
	}
}
