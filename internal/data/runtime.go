package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidRuntimeFormat = errors.New("invalid runtime format")
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	out := fmt.Sprintf("%d min", r)

	quoted := strconv.Quote(out)
	return []byte(quoted), nil
}

func (r *Runtime) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	splitted := strings.Split(unquoted, " ")

	if len(splitted) != 2 || splitted[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(splitted[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
