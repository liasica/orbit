package schema

import (
	"time"

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
		field.Enum("type").Values("underReview", "job", "reviewed").Comment("消息类型"),
		field.String("workitem_id").Optional().Nillable().Comment("工作项ID"),
		field.JSON("varaibales", sonic.NoCopyRawMessage{}).Comment("消息变量"),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}

func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("workitem_id"),
		index.Fields("created_at"),
	}
}
