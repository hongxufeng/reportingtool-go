package function

import (
	"errors"
	"strings"
)

func VerifyInjection(str string) error {
	if strings.Index(str, "'") > -1 || strings.Index(str, "\"") > -1 || strings.Index(str, ";") > -1 {
		return errors.New("Your string has a special character")
	}
	return nil
}
