package conv

import (
	"encoding/json"
	"io"
)

func Parse[T any](jsonData io.Reader) (T, error) {
	decoder := json.NewDecoder(jsonData)
	var dt T
	err := decoder.Decode(&dt)
	return dt, err
}
func JSON(data any) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
