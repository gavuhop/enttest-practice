package service

import (
	"context"
	"ent/ent"
	"ent/ent/transcriptionstats"
)

// CreateTranscriptionStats thêm mới một bản ghi TranscriptionStats
func CreateTranscriptionStats(ctx context.Context, client *ent.Client, duration int64, inputToken, outputToken int, modelName, ownerID string, cost float64, translationID int) (*ent.TranscriptionStats, error) {
	return client.TranscriptionStats.
		Create().
		SetDuration(duration).
		SetInputToken(inputToken).
		SetOutputToken(outputToken).
		SetModelName(modelName).
		SetOwnerID(ownerID).
		SetCost(cost).
		SetTranslationID(translationID).
		Save(ctx)
}

// GetTranscriptionStatsByID đọc một bản ghi TranscriptionStats theo ID
func GetTranscriptionStatsByID(ctx context.Context, client *ent.Client, id int) (*ent.TranscriptionStats, error) {
	return client.TranscriptionStats.
		Query().
		Where(transcriptionstats.IDEQ(id)).
		Only(ctx)
}
