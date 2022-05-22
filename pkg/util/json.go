package util

import "encoding/json"

func IsJsonString(str string) bool {
	return json.Valid([]byte(str))
}
