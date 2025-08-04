package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TranscriptionStats holds the schema definition for the TranscriptionStats entity.
type TranscriptionStats struct {
	ent.Schema
}

// Fields of the TranscriptionStats.
func (TranscriptionStats) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("duration"),
		field.Int("input_token"),
		field.Int("output_token"),
		field.String("model_name").MaxLen(50),
		field.Float("cost"),
		field.String("owner_id"),
		field.Int("translation_id"),
	}
}

// Edges of the TranscriptionStats.
func (TranscriptionStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("translation", Translation.Type).Ref("transcription_stats").Field("translation_id").Unique().Required(),
	}
}
