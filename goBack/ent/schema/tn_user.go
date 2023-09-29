package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TN_USER holds the schema definition for the TN_USER entity.
type TN_USER struct {
	ent.Schema
}

// Fields of the TN_USER.
func (TN_USER) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StorageKey("idx").Unique(),
		field.String("user_id"),
		field.String("user_name"),
		field.String("password"),
		field.String("delect_yn"),
		field.String("reg_date"),
		field.String("udt_date"),
	}
}

// Edges of the TN_USER.
func (TN_USER) Edges() []ent.Edge {
	return nil
}
