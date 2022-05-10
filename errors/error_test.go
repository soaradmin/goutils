package errors

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	fmt.Print(RichError("hellowrold"))
}
