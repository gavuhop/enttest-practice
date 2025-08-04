package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Translation holds the schema definition for the Translation entity.
type Translation struct {
	ent.Schema
}

// Fields of the Translation.
func (Translation) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
		field.JSON("choices", []map[string]interface{}{}),
		field.String("external_id").Unique(),
		field.String("language_pair"),
		field.String("owner_id"),
	}
}

// Edges of the Translation.
func (Translation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transcription_stats", TranscriptionStats.Type),
	}
}
