package conv

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func isStruct(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Struct
}
func To[T any](v any) T {
	var t T
	if v == nil {
		return t
	}
	switch any(t).(type) {
	case string:
		return any(String(v)).(T)
	case int:
		return any(Int(v)).(T)
	case bool:
		return any(Bool(v)).(T)
	case uint:
		return any(Uint(v)).(T)
	case uint64:
		return any(Uint(v)).(T)
	case float32:
		return any(Float(v)).(T)
	case float64:
		return any(Float(v)).(T)
	case []byte:
		return any([]byte(String(v))).(T)
	case time.Time:
		return any(Time(v)).(T)
	default:
		reader, ok := v.(io.Reader)
		if isStruct(v) && ok {
			parse, err := Parse[T](reader)
			if err != nil {
				return t
			}
			return parse
		}
	}
	return t
}
func Time(el any) time.Time {
	var t time.Time
	if el == nil {
		return time.Now()
	}
	switch v := el.(type) {
	case string:
		p, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			p, err = time.Parse("2006-01-02T15:04:05", v)
			if err != nil {
				return t
			}
			return p
		}
		return p
	}
	return t
}
func Float(el any) float64 {
	if el == nil {
		return 0
	}
	switch v := el.(type) {
	case int:
		return float64(v)
	case uint:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0
		}
		return f
	}
	return 0
}
func Uint(el any) uint64 {
	if el == nil {
		return 0
	}
	switch val := el.(type) {
	case string:
		return First(strconv.ParseUint(val, 10, 64))
	case bool:
		if el.(bool) {
			return 1
		}
		return 0
	case int:
		return uint64(val)
	case int64:
		return uint64(val)
	case uint64:
		return val
	case float32:
		return uint64(val)
	case float64:
		return uint64(val)
	default:
		return 0
	}
}
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

func Er[T any](t T, err error) T {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Fprintf(os.Stderr, "%s:%d: %v\n", file, line, err)
	}
	return t
}
