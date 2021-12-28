// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/panforyou/seller-finding/go/ent/website"
)

// Website is the model entity for the Website schema.
type Website struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// URLTemplate holds the value of the "url_template" field.
	// クローリングするURLのテンプレート
	URLTemplate string `json:"url_template,omitempty"`
	// AllowDomains holds the value of the "allow_domains" field.
	// クローリングする際に許容するドメイン
	AllowDomains []string `json:"allow_domains,omitempty"`
	// LatestVisitURL holds the value of the "latest_visit_url" field.
	// 最新のクローリングしたURL（このサイトまできたら処理を中断する）
	LatestVisitURL string `json:"latest_visit_url,omitempty"`
	// MaxPage holds the value of the "max_page" field.
	// クローリングするページの最大数
	MaxPage int `json:"max_page,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Website) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case website.FieldAllowDomains:
			values[i] = new([]byte)
		case website.FieldID, website.FieldMaxPage:
			values[i] = new(sql.NullInt64)
		case website.FieldKey, website.FieldURLTemplate, website.FieldLatestVisitURL:
			values[i] = new(sql.NullString)
		case website.FieldCreatedAt, website.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Website", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Website fields.
func (w *Website) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case website.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case website.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				w.CreatedAt = value.Time
			}
		case website.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				w.UpdatedAt = value.Time
			}
		case website.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				w.Key = value.String
			}
		case website.FieldURLTemplate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url_template", values[i])
			} else if value.Valid {
				w.URLTemplate = value.String
			}
		case website.FieldAllowDomains:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field allow_domains", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &w.AllowDomains); err != nil {
					return fmt.Errorf("unmarshal field allow_domains: %w", err)
				}
			}
		case website.FieldLatestVisitURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field latest_visit_url", values[i])
			} else if value.Valid {
				w.LatestVisitURL = value.String
			}
		case website.FieldMaxPage:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_page", values[i])
			} else if value.Valid {
				w.MaxPage = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Website.
// Note that you need to call Website.Unwrap() before calling this method if this Website
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Website) Update() *WebsiteUpdateOne {
	return (&WebsiteClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Website entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Website) Unwrap() *Website {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Website is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Website) String() string {
	var builder strings.Builder
	builder.WriteString("Website(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(w.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(w.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", key=")
	builder.WriteString(w.Key)
	builder.WriteString(", url_template=")
	builder.WriteString(w.URLTemplate)
	builder.WriteString(", allow_domains=")
	builder.WriteString(fmt.Sprintf("%v", w.AllowDomains))
	builder.WriteString(", latest_visit_url=")
	builder.WriteString(w.LatestVisitURL)
	builder.WriteString(", max_page=")
	builder.WriteString(fmt.Sprintf("%v", w.MaxPage))
	builder.WriteByte(')')
	return builder.String()
}

// Websites is a parsable slice of Website.
type Websites []*Website

func (w Websites) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}