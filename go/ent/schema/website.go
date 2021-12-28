package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Website holds the schema definition for the Website entity.
type Website struct {
	ent.Schema
}

// Fields of the Website.
func (Website) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Unique(),
		field.String("url_template").Unique().Comment("クローリングするURLのテンプレート"),
		field.Strings("allow_domains").Optional().Comment("クローリングする際に許容するドメイン"),
		field.String("latest_visit_url").Optional().Comment("最新のクローリングしたURL（このサイトまできたら処理を中断する）"),
		field.Int("max_page").Default(1000).Comment("クローリングするページの最大数"),
	}
}

func (Website) Edges() []ent.Edge {
	return nil
}

func (Website) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
