package validator

import (
	"regexp"
	"slices"
)

var (
	EmailPattern = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) Set(key, msg string) {
	if _, ok := v.Errors[key]; !ok {
		v.Errors[key] = msg
	}
}

func (v *Validator) Check(ok bool, key, msg string) {
	if !ok {
		v.Set(key, msg)
	}
}

func PermittedValue[T comparable](value T, permitted ...T) bool {
	return slices.Contains(permitted, value)
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func Unique[T comparable](values []T) bool {
	unique := make(map[T]bool)

	for _, v := range values {
		unique[v] = true
	}

	return len(unique) == len(values)
}
