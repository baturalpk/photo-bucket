// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/baturalpk/photo-bucket/ent/photosmetadata"
	"github.com/baturalpk/photo-bucket/ent/profile"
	"github.com/google/uuid"
)

// PhotosMetadata is the model entity for the PhotosMetadata schema.
type PhotosMetadata struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID uuid.UUID `json:"owner_id,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags []string `json:"tags,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Width holds the value of the "width" field.
	Width int `json:"width,omitempty"`
	// Height holds the value of the "height" field.
	Height int `json:"height,omitempty"`
	// ImageFormat holds the value of the "image_format" field.
	ImageFormat photosmetadata.ImageFormat `json:"image_format,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// UploadedAt holds the value of the "uploaded_at" field.
	UploadedAt time.Time `json:"uploaded_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PhotosMetadataQuery when eager-loading is set.
	Edges PhotosMetadataEdges `json:"edges"`
}

// PhotosMetadataEdges holds the relations/edges for other nodes in the graph.
type PhotosMetadataEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Profile `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PhotosMetadataEdges) OwnerOrErr() (*Profile, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: profile.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PhotosMetadata) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case photosmetadata.FieldTags:
			values[i] = new([]byte)
		case photosmetadata.FieldWidth, photosmetadata.FieldHeight:
			values[i] = new(sql.NullInt64)
		case photosmetadata.FieldDescription, photosmetadata.FieldImageFormat, photosmetadata.FieldURL:
			values[i] = new(sql.NullString)
		case photosmetadata.FieldUploadedAt:
			values[i] = new(sql.NullTime)
		case photosmetadata.FieldID, photosmetadata.FieldOwnerID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PhotosMetadata", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PhotosMetadata fields.
func (pm *PhotosMetadata) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case photosmetadata.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pm.ID = *value
			}
		case photosmetadata.FieldOwnerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value != nil {
				pm.OwnerID = *value
			}
		case photosmetadata.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pm.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case photosmetadata.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pm.Description = value.String
			}
		case photosmetadata.FieldWidth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field width", values[i])
			} else if value.Valid {
				pm.Width = int(value.Int64)
			}
		case photosmetadata.FieldHeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				pm.Height = int(value.Int64)
			}
		case photosmetadata.FieldImageFormat:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_format", values[i])
			} else if value.Valid {
				pm.ImageFormat = photosmetadata.ImageFormat(value.String)
			}
		case photosmetadata.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				pm.URL = value.String
			}
		case photosmetadata.FieldUploadedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field uploaded_at", values[i])
			} else if value.Valid {
				pm.UploadedAt = value.Time
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the PhotosMetadata entity.
func (pm *PhotosMetadata) QueryOwner() *ProfileQuery {
	return (&PhotosMetadataClient{config: pm.config}).QueryOwner(pm)
}

// Update returns a builder for updating this PhotosMetadata.
// Note that you need to call PhotosMetadata.Unwrap() before calling this method if this PhotosMetadata
// was returned from a transaction, and the transaction was committed or rolled back.
func (pm *PhotosMetadata) Update() *PhotosMetadataUpdateOne {
	return (&PhotosMetadataClient{config: pm.config}).UpdateOne(pm)
}

// Unwrap unwraps the PhotosMetadata entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pm *PhotosMetadata) Unwrap() *PhotosMetadata {
	tx, ok := pm.config.driver.(*txDriver)
	if !ok {
		panic("ent: PhotosMetadata is not a transactional entity")
	}
	pm.config.driver = tx.drv
	return pm
}

// String implements the fmt.Stringer.
func (pm *PhotosMetadata) String() string {
	var builder strings.Builder
	builder.WriteString("PhotosMetadata(")
	builder.WriteString(fmt.Sprintf("id=%v", pm.ID))
	builder.WriteString(", owner_id=")
	builder.WriteString(fmt.Sprintf("%v", pm.OwnerID))
	builder.WriteString(", tags=")
	builder.WriteString(fmt.Sprintf("%v", pm.Tags))
	builder.WriteString(", description=")
	builder.WriteString(pm.Description)
	builder.WriteString(", width=")
	builder.WriteString(fmt.Sprintf("%v", pm.Width))
	builder.WriteString(", height=")
	builder.WriteString(fmt.Sprintf("%v", pm.Height))
	builder.WriteString(", image_format=")
	builder.WriteString(fmt.Sprintf("%v", pm.ImageFormat))
	builder.WriteString(", url=")
	builder.WriteString(pm.URL)
	builder.WriteString(", uploaded_at=")
	builder.WriteString(pm.UploadedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PhotosMetadataSlice is a parsable slice of PhotosMetadata.
type PhotosMetadataSlice []*PhotosMetadata

func (pm PhotosMetadataSlice) config(cfg config) {
	for _i := range pm {
		pm[_i].config = cfg
	}
}
