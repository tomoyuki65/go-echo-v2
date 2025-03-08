package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("uid").
			NotEmpty().
			Unique(),
		field.String("last_name").
			NotEmpty(),
		field.String("first_name").
			NotEmpty(),
		field.String("email").
			NotEmpty().
			Unique(),
		field.Time("created_at").
			SchemaType(map[string]string{
				"postgres": "timestamptz",
			}).
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			SchemaType(map[string]string{
				"postgres": "timestamptz",
			}).
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			SchemaType(map[string]string{
				"postgres": "timestamptz",
			}).
			Nillable().
			Optional(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("deleted_at"),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
