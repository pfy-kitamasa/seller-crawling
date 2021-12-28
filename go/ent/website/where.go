// Code generated by entc, DO NOT EDIT.

package website

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/panforyou/seller-finding/go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// URLTemplate applies equality check predicate on the "url_template" field. It's identical to URLTemplateEQ.
func URLTemplate(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURLTemplate), v))
	})
}

// LatestVisitURL applies equality check predicate on the "latest_visit_url" field. It's identical to LatestVisitURLEQ.
func LatestVisitURL(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatestVisitURL), v))
	})
}

// MaxPage applies equality check predicate on the "max_page" field. It's identical to MaxPageEQ.
func MaxPage(v int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxPage), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Website {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Website(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Website {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Website(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Website {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Website(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Website {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Website(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKey), v))
	})
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.Website {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Website(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKey), v...))
	})
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.Website {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Website(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKey), v...))
	})
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKey), v))
	})
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKey), v))
	})
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKey), v))
	})
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKey), v))
	})
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldKey), v))
	})
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldKey), v))
	})
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldKey), v))
	})
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldKey), v))
	})
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldKey), v))
	})
}

// URLTemplateEQ applies the EQ predicate on the "url_template" field.
func URLTemplateEQ(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURLTemplate), v))
	})
}

// URLTemplateNEQ applies the NEQ predicate on the "url_template" field.
func URLTemplateNEQ(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURLTemplate), v))
	})
}

// URLTemplateIn applies the In predicate on the "url_template" field.
func URLTemplateIn(vs ...string) predicate.Website {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Website(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURLTemplate), v...))
	})
}

// URLTemplateNotIn applies the NotIn predicate on the "url_template" field.
func URLTemplateNotIn(vs ...string) predicate.Website {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Website(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURLTemplate), v...))
	})
}

// URLTemplateGT applies the GT predicate on the "url_template" field.
func URLTemplateGT(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURLTemplate), v))
	})
}

// URLTemplateGTE applies the GTE predicate on the "url_template" field.
func URLTemplateGTE(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURLTemplate), v))
	})
}

// URLTemplateLT applies the LT predicate on the "url_template" field.
func URLTemplateLT(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURLTemplate), v))
	})
}

// URLTemplateLTE applies the LTE predicate on the "url_template" field.
func URLTemplateLTE(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURLTemplate), v))
	})
}

// URLTemplateContains applies the Contains predicate on the "url_template" field.
func URLTemplateContains(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURLTemplate), v))
	})
}

// URLTemplateHasPrefix applies the HasPrefix predicate on the "url_template" field.
func URLTemplateHasPrefix(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURLTemplate), v))
	})
}

// URLTemplateHasSuffix applies the HasSuffix predicate on the "url_template" field.
func URLTemplateHasSuffix(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURLTemplate), v))
	})
}

// URLTemplateEqualFold applies the EqualFold predicate on the "url_template" field.
func URLTemplateEqualFold(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURLTemplate), v))
	})
}

// URLTemplateContainsFold applies the ContainsFold predicate on the "url_template" field.
func URLTemplateContainsFold(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURLTemplate), v))
	})
}

// AllowDomainsIsNil applies the IsNil predicate on the "allow_domains" field.
func AllowDomainsIsNil() predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAllowDomains)))
	})
}

// AllowDomainsNotNil applies the NotNil predicate on the "allow_domains" field.
func AllowDomainsNotNil() predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAllowDomains)))
	})
}

// LatestVisitURLEQ applies the EQ predicate on the "latest_visit_url" field.
func LatestVisitURLEQ(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatestVisitURL), v))
	})
}

// LatestVisitURLNEQ applies the NEQ predicate on the "latest_visit_url" field.
func LatestVisitURLNEQ(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLatestVisitURL), v))
	})
}

// LatestVisitURLIn applies the In predicate on the "latest_visit_url" field.
func LatestVisitURLIn(vs ...string) predicate.Website {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Website(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLatestVisitURL), v...))
	})
}

// LatestVisitURLNotIn applies the NotIn predicate on the "latest_visit_url" field.
func LatestVisitURLNotIn(vs ...string) predicate.Website {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Website(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLatestVisitURL), v...))
	})
}

// LatestVisitURLGT applies the GT predicate on the "latest_visit_url" field.
func LatestVisitURLGT(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLatestVisitURL), v))
	})
}

// LatestVisitURLGTE applies the GTE predicate on the "latest_visit_url" field.
func LatestVisitURLGTE(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLatestVisitURL), v))
	})
}

// LatestVisitURLLT applies the LT predicate on the "latest_visit_url" field.
func LatestVisitURLLT(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLatestVisitURL), v))
	})
}

// LatestVisitURLLTE applies the LTE predicate on the "latest_visit_url" field.
func LatestVisitURLLTE(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLatestVisitURL), v))
	})
}

// LatestVisitURLContains applies the Contains predicate on the "latest_visit_url" field.
func LatestVisitURLContains(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLatestVisitURL), v))
	})
}

// LatestVisitURLHasPrefix applies the HasPrefix predicate on the "latest_visit_url" field.
func LatestVisitURLHasPrefix(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLatestVisitURL), v))
	})
}

// LatestVisitURLHasSuffix applies the HasSuffix predicate on the "latest_visit_url" field.
func LatestVisitURLHasSuffix(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLatestVisitURL), v))
	})
}

// LatestVisitURLIsNil applies the IsNil predicate on the "latest_visit_url" field.
func LatestVisitURLIsNil() predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLatestVisitURL)))
	})
}

// LatestVisitURLNotNil applies the NotNil predicate on the "latest_visit_url" field.
func LatestVisitURLNotNil() predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLatestVisitURL)))
	})
}

// LatestVisitURLEqualFold applies the EqualFold predicate on the "latest_visit_url" field.
func LatestVisitURLEqualFold(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLatestVisitURL), v))
	})
}

// LatestVisitURLContainsFold applies the ContainsFold predicate on the "latest_visit_url" field.
func LatestVisitURLContainsFold(v string) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLatestVisitURL), v))
	})
}

// MaxPageEQ applies the EQ predicate on the "max_page" field.
func MaxPageEQ(v int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxPage), v))
	})
}

// MaxPageNEQ applies the NEQ predicate on the "max_page" field.
func MaxPageNEQ(v int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaxPage), v))
	})
}

// MaxPageIn applies the In predicate on the "max_page" field.
func MaxPageIn(vs ...int) predicate.Website {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Website(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMaxPage), v...))
	})
}

// MaxPageNotIn applies the NotIn predicate on the "max_page" field.
func MaxPageNotIn(vs ...int) predicate.Website {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Website(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMaxPage), v...))
	})
}

// MaxPageGT applies the GT predicate on the "max_page" field.
func MaxPageGT(v int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaxPage), v))
	})
}

// MaxPageGTE applies the GTE predicate on the "max_page" field.
func MaxPageGTE(v int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaxPage), v))
	})
}

// MaxPageLT applies the LT predicate on the "max_page" field.
func MaxPageLT(v int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaxPage), v))
	})
}

// MaxPageLTE applies the LTE predicate on the "max_page" field.
func MaxPageLTE(v int) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaxPage), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Website) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Website) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
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
func Not(p predicate.Website) predicate.Website {
	return predicate.Website(func(s *sql.Selector) {
		p(s.Not())
	})
}
