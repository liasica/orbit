package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/bytedance/sonic"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Annotations of the Message.
func (Message) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "message"},
		entsql.WithComments(true),
	}
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("message_id").Unique().Comment("消息ID"),
		field.String("workitem_id").Optional().Nillable().Comment("工作项ID"),
		field.JSON("varaibales", sonic.NoCopyRawMessage{}).Comment("消息变量"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}

func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("workitem_id"),
	}
}
