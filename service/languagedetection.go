package service

import (
	"context"
	"ent/ent"
	"ent/ent/languagedetection"
)

// CreateLanguageDetection thêm mới một bản ghi LanguageDetection
func CreateLanguageDetection(ctx context.Context, client *ent.Client, query string, detectedLanguages map[string]interface{}, duration int64, modelName string, inputToken, outputToken, cachedToken int, ownerID string, cost float64) (*ent.LanguageDetection, error) {
	return client.LanguageDetection.
		Create().
		SetQuery(query).
		SetDetectedLanguages(detectedLanguages).
		SetDuration(duration).
		SetModelName(modelName).
		SetInputToken(inputToken).
		SetOutputToken(outputToken).
		SetCachedToken(cachedToken).
		SetOwnerID(ownerID).
		SetCost(cost).
		Save(ctx)
}

// GetLanguageDetectionByID đọc một bản ghi LanguageDetection theo ID
func GetLanguageDetectionByID(ctx context.Context, client *ent.Client, id int) (*ent.LanguageDetection, error) {
	return client.LanguageDetection.
		Query().
		Where(languagedetection.IDEQ(id)).
		Only(ctx)
}
