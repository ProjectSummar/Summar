package utils

import "encoding/json"

func JSONMarshal(v interface{}) string {
	s, _ := json.MarshalIndent(v, "", "\t")
	return string(s)
}
