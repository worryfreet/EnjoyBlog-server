package utils

import (
	"encoding/json"
	"github.com/google/uuid"
	"strings"
)

func NewUUID() string {
	u, _ := uuid.NewUUID()
	return strings.ReplaceAll(u.String(), "-", "")
}

func FillModel(dest, model interface{}) error {
	modelBytes, err := json.Marshal(model)
	if err != nil {
		return err
	}
	return json.Unmarshal(modelBytes, dest)
}
