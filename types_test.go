package conv

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestRecursiveValue(t *testing.T) {
	assertions := assert.New(t)
	data := map[string]any{
		"content": map[string]any{
			"lastVersion": map[string]any{
				"attributes": map[string]any{
					"Poss": "value123",
					"nested": map[string]any{
						"deeper": "deepValue",
					},
					"list": []any{"a", "b", "c"},
				},
			},
		},
	}

	// Тестируем
	tests := []struct {
		path    string
		wantErr bool
		result  any
	}{
		{"content.lastVersion.attributes.Poss", false, "value123"},
		{"content.lastVersion.attributes.nested.deeper", false, "deepValue"},
		{"content.lastVersion.attributes.list.1", false, "b"}, // доступ к элементу массива
		{"content.lastVersion.attributes.missing", true, nil},
		{"content.invalid.nested", true, nil},
		{"", false, data}, // возвращает весь data
	}

	for _, test := range tests {
		val, err := RecursiveValue(data, test.path)
		fmt.Printf("Path: %s\n", test.path)
		if err != nil {
			assertions.Error(err)
			fmt.Printf("  Error: %v\n", err)
		} else {
			assertions.Equal(test.result, val)
			fmt.Printf("  Value: %v (type: %T)\n", val, val)
		}
		fmt.Println()
	}

	// Пример с массивом
	arrayData := map[string]any{
		"users": []any{
			map[string]any{"name": "Alice", "age": 30},
			map[string]any{"name": "Bob", "age": 25},
		},
	}

	val, err := RecursiveValue(arrayData, "users.1.name")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("User 1 name: %v\n", val) // "Bob"
	}
}
func TestString(t *testing.T) {
	assertions := assert.New(t)
	num := int64(12)
	val := String(&num)
	assertions.Equal("12", val)
	val = String(num)
	assertions.Equal("12", val)
}
func TestPtr(t *testing.T) {
	fmt.Println("TestTo")
	assertions := assert.New(t)
	v := "hello world"
	assertions.Equal(&v, Ptr[string](v))
}
func TestOrig(t *testing.T) {
	assertions := assert.New(t)
	v := "hello world"
	assertions.Equal(Orig(&v), v)
	assertions.Equal(Orig[string](nil), "")
}
func TestTo(t *testing.T) {
	fmt.Println("TestTo")
	assertions := assert.New(t)
	assertions.Equal("hello world", To[string]("hello world"))
	assertions.Equal("1", To[string](1))
	assertions.Equal("true", To[string](true))
	assertions.Equal("235", To[string](235))
	assertions.Equal("235.44", To[string](235.44))
}
func TestErr(t *testing.T) {
	fmt.Println("TestErr")
	assertions := assert.New(t)
	assertions.Equal(uint64(0), Er(strconv.ParseUint("hello world", 10, 64)))
}
