package conv

import (
	"errors"
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

// RecursiveValue принемает на вход рекурсивный point, и строку, раскладывающую рекурсию и возвращающую значение
//
//nolint:unused
func RecursiveValue(data any, path string) (any, error) {
	if data == nil {
		return nil, errors.New("data is nil")
	}

	if path == "" {
		return data, nil
	}

	keys := strings.Split(path, ".")
	var current any = data

	for i, key := range keys {
		switch v := current.(type) {
		case map[string]any:
			val, exists := v[key]
			if !exists {
				return nil, fmt.Errorf("key '%s' not found at path '%s'", key, strings.Join(keys[:i+1], "."))
			}
			current = val

		case []any:
			// Если ключ - числовой индекс для массива
			index := -1
			_, err := fmt.Sscanf(key, "%d", &index)
			if err != nil {
				return nil, fmt.Errorf("expected array index, got '%s' at path '%s'", key, strings.Join(keys[:i+1], "."))
			}

			if index < 0 || index >= len(v) {
				return nil, fmt.Errorf("array index %d out of bounds at path '%s'", index, strings.Join(keys[:i+1], "."))
			}
			current = v[index]

		default:
			return nil, fmt.Errorf("cannot access key '%s' on non-map/non-array type %T at path '%s'", key, current, strings.Join(keys[:i+1], "."))
		}
	}

	return current, nil
}

// To is ext func
//
//nolint:unused
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
	case int32:
		return any(Int32(v)).(T)
	case int64:
		return any(Int64(v)).(T)
	case bool:
		return any(Bool(v)).(T)
	case uint:
		return any(uint(Uint(v))).(T)
	case uint32:
		return any(uint32(Uint(v))).(T)
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
func Int64(v any) int64 {
	if v == nil {
		return 0
	}
	switch kv := v.(type) {
	case bool:
		if kv {
			return 1
		}
		return 0
	case int:
		return int64(kv)
	case int32:
		return int64(kv)
	case int64:
		return kv
	case string:
		return First(strconv.ParseInt(kv, 10, 64))
	}
	return 0
}
func Int32(v any) int32 {
	if v == nil {
		return 0
	}
	switch kv := v.(type) {
	case bool:
		if kv {
			return 1
		}
		return 0
	case int:
		return int32(kv)
	case int32:
		return kv
	case int64:
		return int32(kv)
	case string:
		return int32(First(strconv.ParseInt(kv, 10, 32)))
	}
	return 0
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

// Ptr convert elements to any types and get point
//
//nolint:unused
func Ptr[T any](v any) *T {
	var t T = To[T](v)
	return &t
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
	case int32:
		return strconv.FormatInt(int64(v.(int32)), 10)
	case int64:
		return strconv.FormatInt(v.(int64), 10)
	case uint64:
		return strconv.FormatUint(v.(uint64), 10)
	case float32:
		return strconv.FormatFloat(float64(v.(float32)), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case time.Time:
		return (v.(time.Time)).Format(time.RFC3339)
	default:
		return fmt.Sprintf("%v", v)
	}
}
func First[T any](t T, _ any) T {
	return t
}

// Er is ext func
//
//nolint:unused
func Er[T any](t T, err error) T {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		_, _ = fmt.Fprintf(os.Stderr, "%s:%d: %v\n", file, line, err)
	}
	return t
}
