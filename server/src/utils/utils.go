package utils

import (
	"encoding/json"
)

func JSONMarshalIndent(v interface{}) string {
	s, _ := json.MarshalIndent(v, "", "\t")
	return string(s)
}

func JSONMarshal(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

func JSONUnmarshal[T any](body []byte) (T, error) {
	var resBody T

	if err := json.Unmarshal(body, &resBody); err != nil {
		return resBody, err
	} else {
		return resBody, nil
	}
}

func StringPtr(s string) *string {
	return &s
}
