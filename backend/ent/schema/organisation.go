package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Organisation holds the schema definition for the Organisation entity.
type Organisation struct {
	ent.Schema
}

func (Organisation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Organisation.
func (Organisation) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Organisation.
func (Organisation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
