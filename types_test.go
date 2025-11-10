package conv

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println("TestString")
	assertions := assert.New(t)
	assertions.Equal("hello world", String("hello world"))
	assertions.Equal("1", String(1))
	assertions.Equal("true", String(true))
	assertions.Equal("235", String(235))
	assertions.Equal("235.44", String(235.44))
}
