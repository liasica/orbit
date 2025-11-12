package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Repository holds the schema definition for the Repository entity.
type Repository struct {
	ent.Schema
}

// Annotations of the Repository.
func (Repository) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "repository"},
		entsql.WithComments(true),
	}
}

// Fields of the Repository.
func (Repository) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("path").Unique().Comment("仓库路径"),
	}
}

// Edges of the Repository.
func (Repository) Edges() []ent.Edge {
	return nil
}
