package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("userID"),
		field.String("email"),
		field.String("firstName"),
		field.String("lastName"),
		field.String("password").Sensitive(),
		field.Bool("resetPassword"),
		field.String("verificationCode").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organisation", Organisation.Type).
			Ref("users").
			Unique(),
	}
}
