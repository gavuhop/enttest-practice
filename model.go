package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type BaseType struct {
	ID        uint      `gorm:"primaryKey;type:int unsigned"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Choices struct {
	Confidence float64                `json:"confidence"`
	Content    string                 `json:"content"`
	Extras     map[string]interface{} `json:"extras"`
}

type ChoicesList []Choices

// Value - Marshal to JSON for writing to DB
func (c ChoicesList) Value() (driver.Value, error) {
	return json.Marshal(c)
}
func (c *ChoicesList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("cannot convert %T to ChoicesList", value)
	}
	if err := json.Unmarshal(bytes, c); err != nil {
		return err
	}

	// Fix nil Extras
	for i := range *c {
		if (*c)[i].Extras == nil {
			(*c)[i].Extras = map[string]interface{}{}
		}
	}
	return nil
}

// Translation represents the translations table
type Translation struct {
	BaseType
	Content      string
	Choices      ChoicesList `gorm:"type:json;not null" json:"choices"`
	ExternalID   string      `gorm:"size:255;uniqueIndex" json:"external_id"`
	LanguagePair string      `gorm:"size:255;not null" json:"language_pair"`
	OwnerID      string      `gorm:"size:255" json:"owner_id"`

	TranscriptionStats []TranscriptionStats `gorm:"foreignKey:TranslationID;constraint:OnDelete:CASCADE" json:"transcription_stats,omitempty"`
}

// TranscriptionStats represents the transcription_stats table
type TranscriptionStats struct {
	BaseType
	Duration      int64   `gorm:"not null" json:"duration"`
	InputToken    int     `gorm:"not null" json:"input_token"`
	OutputToken   int     `gorm:"not null" json:"output_token"`
	ModelName     string  `gorm:"type:varchar(50);not null" json:"model_name"`
	Cost          float64 `gorm:"type:decimal(20,10)" json:"cost"`
	OwnerID       string  `gorm:"size:255" json:"owner_id"`
	TranslationID uint    `gorm:"not null;index;type:int unsigned" json:"translation_id"`
}

// DetectedLanguagesMap represents a map of language codes to certainty values
type DetectedLanguagesMap map[string]float32

// Value - Marshal to JSON for writing to DB
func (d DetectedLanguagesMap) Value() (driver.Value, error) {
	return json.Marshal(d)
}

// Scan - Unmarshal from JSON when reading from DB
func (d *DetectedLanguagesMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("cannot convert %T to DetectedLanguagesMap", value)
	}
	return json.Unmarshal(bytes, d)
}

type LanguageDetection struct {
	BaseType
	Query             string               `gorm:"type:text;not null" json:"query"`
	DetectedLanguages DetectedLanguagesMap `gorm:"type:json" json:"detected_languages"`
	Duration          int64                `gorm:"not null" json:"duration"`
	ModelName         string               `gorm:"type:varchar(50);not null" json:"model_name"`
	InputToken        int                  `gorm:"not null" json:"input_token"`
	OutputToken       int                  `gorm:"not null" json:"output_token"`
	CachedToken       int                  `gorm:"not null" json:"cached_token"`
	OwnerID           string               `gorm:"size:255" json:"owner_id"`
	Cost              float64              `gorm:"type:decimal(30,20)" json:"cost"`
}
