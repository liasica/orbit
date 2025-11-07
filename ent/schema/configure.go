package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/bytedance/sonic"
)

// Configure holds the schema definition for the Configure entity.
type Configure struct {
	ent.Schema
}

func (Configure) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "configure"},
		entsql.WithComments(true),
	}
}

// Fields of the Configure.
func (Configure) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("key").Values("yunxiao", "gitlab_merge_targets").Comment("配置键"),
		field.JSON("data", sonic.NoCopyRawMessage{}).Comment("配置值"),
	}
}

// Edges of the Configure.
func (Configure) Edges() []ent.Edge {
	return nil
}

func (Configure) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key").Unique(),
	}
}
