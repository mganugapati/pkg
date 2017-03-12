package validator

import (
	"errors"
	"regexp"
)

type Validator interface {
	Validate() error
}

// ValidatorFunc represents a function which implements Validator interface
type ValidatorFunc func() error

func (v ValidatorFunc) Validate() error { return v() }

// ValidatorList holds n Validators, and implements Validator interface too
type ValidatorList []Validator

func (v ValidatorList) Validate() error {
	for _, x := range v {
		if err := x.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// AndRequire appends a boolValidator
func (v *ValidatorList) AndRequire(val bool, msg string) *ValidatorList {
	(*v) = append((*v), boolValidator{val: val, msg: msg})
	return v
}

// And appends a Validator
func (v *ValidatorList) And(validator Validator) *ValidatorList {
	(*v) = append((*v), validator)
	return v
}

// boolValidator implements Validator which used to valite a boolean value should be true
type boolValidator struct {
	val bool
	msg string
}

func (v boolValidator) Validate() error {
	if !v.val {
		return errors.New(v.msg)
	}
	return nil
}

// Require returns a ValidatorList which holds a boolValidator created by (val,msg)
func Require(val bool, msg string) *ValidatorList {
	v := new(ValidatorList)
	return v.AndRequire(val, msg)
}

var emailRegexp = regexp.MustCompilePOSIX("^[A-Za-z0-9\\u4e00-\\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")

func IsValidEmail(email string) bool {
	return emailRegexp.MatchString(email)
}
