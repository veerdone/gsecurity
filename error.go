package gsecurity

import (
	"errors"
	"fmt"
)

var (
	ErrBeReplace = errors.New("has been replaced")
	ErrNotLogin  = errors.New("not login")
	abnormalMap  = map[int64]error{
		BeReplace: ErrBeReplace,
	}
)

// isValidId check id is valid
func isValidId(id int64) bool {
	return checkValidId(id) == nil
}

// checkValidId check id is valid, if it is not valid, return error
func checkValidId(id int64) error {
	return abnormalMap[id]
}

type ErrDisable struct {
	level   int64
	service string
}

func (e ErrDisable) Error() string {
	return fmt.Sprintf("disable level=%d, service=%s", e.level, e.service)
}

func NewErrDisable(level int64, service string) error {
	return ErrDisable{level: level, service: service}
}
