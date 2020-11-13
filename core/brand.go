package core

import "github.com/louisevanderlith/husk/validation"

type Brand struct {
	Name string
}

func (b Brand) Valid() error {
	return validation.Struct(b)
}
