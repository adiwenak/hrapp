package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.String("username"),
		field.String("mobileNumber"),
		field.String("firstName"),
		field.String("lastName"),
		field.String("password").Sensitive(),
		field.Bool("needPasswordReset"),
		field.String("verificationCode").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organisation", Organisation.Type).
			Ref("users").
			Unique().
			Required(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email", "username", "mobileNumber").Unique(),
	}
}
