package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
		entsql.WithComments(true),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("lark_user_id").Unique().Comment("飞书用户ID"),
		field.String("lark_union_id").Comment("飞书 UnionId"),
		field.String("lark_open_id").Comment("飞书 OpenId"),
		field.String("yunxiao_user_id").Comment("云效用户ID"),
		field.String("gitlab_username").Comment("GitLab 用户名"),
		field.String("gitlab_email").Comment("GitLab 用户邮箱"),
		field.String("name").Comment("用户名称"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("lark_union_id"),
		index.Fields("lark_open_id"),
		index.Fields("yunxiao_user_id"),
		index.Fields("gitlab_username"),
		index.Fields("gitlab_email"),
		index.Fields("name"),
	}
}
