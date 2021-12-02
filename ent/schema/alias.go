package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid"
)

// Alias holds the schema definition for the Alias entity.
type Alias struct {
	ent.Schema
}

// Fields of the Alias.
func (Alias) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.String("short"),
		field.String("long"),
	}
}

// Edges of the Alias.
func (Alias) Edges() []ent.Edge {
	return nil
}
