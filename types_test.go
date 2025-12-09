package conv

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

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
