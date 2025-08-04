package service

import (
	"context"
	"ent/ent"
	"ent/ent/translation"
)

// CreateTranslation thêm mới một bản ghi Translation
func CreateTranslation(ctx context.Context, client *ent.Client, content, externalID, languagePair, ownerID string, choices []map[string]interface{}) (*ent.Translation, error) {
	return client.Translation.
		Create().
		SetContent(content).
		SetExternalID(externalID).
		SetLanguagePair(languagePair).
		SetOwnerID(ownerID).
		SetChoices(choices).
		Save(ctx)
}

// GetTranslationByID đọc một bản ghi Translation theo ID
func GetTranslationByID(ctx context.Context, client *ent.Client, id int) (*ent.Translation, error) {
	return client.Translation.
		Query().
		Where(translation.IDEQ(id)).
		Only(ctx)
}
