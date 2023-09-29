package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TN_BBS holds the schema definition for the TN_BBS entity.
type TN_BBS struct {
	ent.Schema
}

// Fields of the TN_BBS.
func (TN_BBS) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StorageKey("idx").Unique(),
		field.String("user_id"),
		field.String("user_name"),
		field.String("title"),
		field.String("content"),
		field.String("delect_yn"),
		field.String("reg_date"),
		field.String("udt_date"),
	}
}

// Edges of the TN_BBS.
func (TN_BBS) Edges() []ent.Edge {
	return nil
}
