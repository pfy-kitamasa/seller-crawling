package schema

import (
	"entgo.io/ent/schema/index"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Seller holds the schema definition for the Seller entity.
type Seller struct {
	ent.Schema
}

// Fields of the Seller.
func (Seller) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").Unique().Comment("クローリングしたURL"),
		field.String("key").Unique().Comment("クローリングしたサイトで一意となるキー"),
		field.String("name").Optional().Comment("店名"),
		field.String("tel").Optional().Comment("電話番号"),
		field.String("prefecture").Optional().Comment("都道府県"),
		field.String("address").Optional().Comment("住所"),
		field.Float("latitude").Optional().Comment("緯度"),
		field.Float("longitude").Optional().Comment("経度"),
		field.Time("opened_at").Optional().Comment("オープン日"),
		field.Bool("exported").Comment("出力済み"),
		field.Time("exported_at").Optional().Comment("出力日"),
		field.Bool("disabled").Comment("不要フラグ"),
	}
}

func (Seller) Edges() []ent.Edge {
	return nil
}

func (Seller) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("prefecture", "opened_at"),
	}
}

func (Seller) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
