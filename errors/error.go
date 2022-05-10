package errors

import (
	"fmt"
	"github.com/soaradmin/goutils/internal"
	"strings"
)

func SafelyError(err error, defaultTip string) error {
	if err != nil {
		return err
	}

	return fmt.Errorf(defaultTip)
}

func SafeErrorDesc(err error, defaultTip string) string {
	if err != nil {
		return err.Error()
	}

	return defaultTip
}

//RichError print the caller
func RichError(format string, a ...interface{}) error {
	var (
		file, pos = internal.Caller()
		ss        = strings.Split(file, "/")
	)
	return fmt.Errorf(fmt.Sprintf(`%s:%d, `, ss[len(ss)-1], pos) + fmt.Sprintf(format, a...))
}