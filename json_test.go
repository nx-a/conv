package conv

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_JSON(t *testing.T) {
	assertions := assert.New(t)
	val := JSON(struct {
		Id   uint8  `json:"id"`
		Name string `json:"name"`
	}{
		Id:   1,
		Name: "test",
	})
	fmt.Println(val)
	assertions.NotNil(val)
	assertions.Equal("{\"id\":1,\"name\":\"test\"}", val, "JSON conversion failed")
	if val != "{\"id\":1,\"name\":\"test\"}" {
		t.Errorf("JSON conversion failed")
	}
}
