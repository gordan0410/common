package helper

import (
	"encoding/json"
)

func JsonReMarshal(v any, dst any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &dst)
	if err != nil {
		return err
	}
	return nil
}

func JsonMarshal(v any) string {
	jsonBytes, _ := json.Marshal(v)
	return string(jsonBytes)
}

func IsJSONObject(data json.RawMessage) bool {
	if len(data) == 0 {
		return false
	}

	if !json.Valid(data) {
		return false
	}

	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return false
	}

	_, ok := v.(map[string]any)

	return ok
}

func JsonMarshalBytes(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		return []byte("")
	}
	return b
}
