package gotest

import (
	"errors"
	"fmt"
)

// MakeByErrorsNew 方法
func MakeByErrorsNew() error {
	return errors.New("new error")
}

// MakeByFmtErrorf 方法
func MakeByFmtErrorf() error {
	return fmt.Errorf("new error")
}
