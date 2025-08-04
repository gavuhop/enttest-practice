package schema

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Choices struct {
	Confidence float64                `json:"confidence"`
	Content    string                 `json:"content"`
	Extras     map[string]interface{} `json:"extras"`
}

type ChoicesList []Choices

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
	for i := range *c {
		if (*c)[i].Extras == nil {
			(*c)[i].Extras = map[string]interface{}{}
		}
	}
	return nil
}

func (c ChoicesList) MarshalJSON() ([]byte, error) {
	return json.Marshal([]Choices(c))
}

func (c *ChoicesList) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*[]Choices)(c))
}

type DetectedLanguagesMap map[string]float32

func (d DetectedLanguagesMap) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *DetectedLanguagesMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("cannot convert %T to DetectedLanguagesMap", value)
	}
	return json.Unmarshal(bytes, d)
}

func (d DetectedLanguagesMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]float32(d))
}

func (d *DetectedLanguagesMap) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*map[string]float32)(d))
}
