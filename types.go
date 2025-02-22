package conv

import (
	"strconv"
	"strings"
)

func Bool(v any) bool {
	if v == nil {
		return false
	}
	switch v.(type) {
	case bool:
		return v.(bool)
	case int:
		return v.(int) != 0
	case string:
		return strings.ToLower(v.(string)) == "true"
	}
	return false
}

func Int(v any) int {
	if v == nil {
		return 0
	}
	switch v.(type) {
	case bool:
		if v.(bool) {
			return 1
		}
		return 0
	case int:
		return v.(int)
	case string:
		return First(strconv.Atoi(v.(string)))
	}
	return 0
}
func String(v any) string {
	if v == nil {
		return ""
	}
	switch v.(type) {
	case string:
		return v.(string)
	case bool:
		if v.(bool) {
			return "true"
		}
		return "false"
	case int:
		return strconv.Itoa(v.(int))
	case uint64:
		return strconv.FormatUint(v.(uint64), 10)
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	default:
		return ""
	}
}
func First[T any](t T, _ any) T {
	return t
}
