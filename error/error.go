package error

import "fmt"

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
