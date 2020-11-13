package core

import "github.com/louisevanderlith/husk/validation"

type Type struct {
	Name  string //Shoes, Shorts, etc
	Sizes []Size
}

func (t Type) Valid() error {
	return validation.Struct(t)
}
