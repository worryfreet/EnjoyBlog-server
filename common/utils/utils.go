package utils

import (
	"encoding/json"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"strconv"
)

func NewObjID() string {
	paramId := int64(0)
	snowId, err := snowflake.NewSnowflake(paramId, paramId)
	if err != nil {
		return ""
	}
	return strconv.FormatInt(snowId.NextVal(), 10)
}

func FillModel(dest, model interface{}) error {
	modelBytes, err := json.Marshal(model)
	if err != nil {
		return err
	}
	return json.Unmarshal(modelBytes, dest)
}

// Wait For Updating...
