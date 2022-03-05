// Code generated by entc, DO NOT EDIT.

package photosmetadata

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/baturalpk/photo-bucket/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v uuid.UUID) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwnerID), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Width applies equality check predicate on the "width" field. It's identical to WidthEQ.
func Width(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWidth), v))
	})
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeight), v))
	})
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// UploadedAt applies equality check predicate on the "uploaded_at" field. It's identical to UploadedAtEQ.
func UploadedAt(v time.Time) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUploadedAt), v))
	})
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v uuid.UUID) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOwnerID), v))
	})
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v uuid.UUID) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOwnerID), v))
	})
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...uuid.UUID) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOwnerID), v...))
	})
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...uuid.UUID) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOwnerID), v...))
	})
}

// OwnerIDIsNil applies the IsNil predicate on the "owner_id" field.
func OwnerIDIsNil() predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOwnerID)))
	})
}

// OwnerIDNotNil applies the NotNil predicate on the "owner_id" field.
func OwnerIDNotNil() predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOwnerID)))
	})
}

// TagsIsNil applies the IsNil predicate on the "tags" field.
func TagsIsNil() predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTags)))
	})
}

// TagsNotNil applies the NotNil predicate on the "tags" field.
func TagsNotNil() predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTags)))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// WidthEQ applies the EQ predicate on the "width" field.
func WidthEQ(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWidth), v))
	})
}

// WidthNEQ applies the NEQ predicate on the "width" field.
func WidthNEQ(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWidth), v))
	})
}

// WidthIn applies the In predicate on the "width" field.
func WidthIn(vs ...int) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWidth), v...))
	})
}

// WidthNotIn applies the NotIn predicate on the "width" field.
func WidthNotIn(vs ...int) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWidth), v...))
	})
}

// WidthGT applies the GT predicate on the "width" field.
func WidthGT(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWidth), v))
	})
}

// WidthGTE applies the GTE predicate on the "width" field.
func WidthGTE(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWidth), v))
	})
}

// WidthLT applies the LT predicate on the "width" field.
func WidthLT(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWidth), v))
	})
}

// WidthLTE applies the LTE predicate on the "width" field.
func WidthLTE(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWidth), v))
	})
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeight), v))
	})
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHeight), v))
	})
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...int) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHeight), v...))
	})
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...int) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHeight), v...))
	})
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHeight), v))
	})
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHeight), v))
	})
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHeight), v))
	})
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v int) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHeight), v))
	})
}

// ImageFormatEQ applies the EQ predicate on the "image_format" field.
func ImageFormatEQ(v ImageFormat) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImageFormat), v))
	})
}

// ImageFormatNEQ applies the NEQ predicate on the "image_format" field.
func ImageFormatNEQ(v ImageFormat) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImageFormat), v))
	})
}

// ImageFormatIn applies the In predicate on the "image_format" field.
func ImageFormatIn(vs ...ImageFormat) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImageFormat), v...))
	})
}

// ImageFormatNotIn applies the NotIn predicate on the "image_format" field.
func ImageFormatNotIn(vs ...ImageFormat) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImageFormat), v...))
	})
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURL), v))
	})
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURL), v...))
	})
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURL), v...))
	})
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURL), v))
	})
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURL), v))
	})
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURL), v))
	})
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURL), v))
	})
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURL), v))
	})
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURL), v))
	})
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURL), v))
	})
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURL), v))
	})
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURL), v))
	})
}

// UploadedAtEQ applies the EQ predicate on the "uploaded_at" field.
func UploadedAtEQ(v time.Time) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUploadedAt), v))
	})
}

// UploadedAtNEQ applies the NEQ predicate on the "uploaded_at" field.
func UploadedAtNEQ(v time.Time) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUploadedAt), v))
	})
}

// UploadedAtIn applies the In predicate on the "uploaded_at" field.
func UploadedAtIn(vs ...time.Time) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUploadedAt), v...))
	})
}

// UploadedAtNotIn applies the NotIn predicate on the "uploaded_at" field.
func UploadedAtNotIn(vs ...time.Time) predicate.PhotosMetadata {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUploadedAt), v...))
	})
}

// UploadedAtGT applies the GT predicate on the "uploaded_at" field.
func UploadedAtGT(v time.Time) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUploadedAt), v))
	})
}

// UploadedAtGTE applies the GTE predicate on the "uploaded_at" field.
func UploadedAtGTE(v time.Time) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUploadedAt), v))
	})
}

// UploadedAtLT applies the LT predicate on the "uploaded_at" field.
func UploadedAtLT(v time.Time) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUploadedAt), v))
	})
}

// UploadedAtLTE applies the LTE predicate on the "uploaded_at" field.
func UploadedAtLTE(v time.Time) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUploadedAt), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Profile) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PhotosMetadata) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PhotosMetadata) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PhotosMetadata) predicate.PhotosMetadata {
	return predicate.PhotosMetadata(func(s *sql.Selector) {
		p(s.Not())
	})
}
