package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// LanguageDetection holds the schema definition for the LanguageDetection entity.
type LanguageDetection struct {
	ent.Schema
}

// Fields of the LanguageDetection.
func (LanguageDetection) Fields() []ent.Field {
	return []ent.Field{
		field.String("query"),
		field.JSON("detected_languages", map[string]interface{}{}),
		field.Int64("duration"),
		field.String("model_name").MaxLen(50),
		field.Int("input_token"),
		field.Int("output_token"),
		field.Int("cached_token"),
		field.String("owner_id"),
		field.Float("cost"),
	}
}
