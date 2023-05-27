package utils

import "encoding/json"

func JSONMarshalIndent(v interface{}) string {
	s, _ := json.MarshalIndent(v, "", "\t")
	return string(s)
}

func JSONMarshal(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

func StringPtr(s string) *string {
	return &s
}
