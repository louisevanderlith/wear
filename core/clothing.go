package core

import "github.com/louisevanderlith/husk/validation"

type Clothing struct {
	Type   string
	Size   string
	Colour string
}

func (c Clothing) Valid() error {
	return validation.Struct(c)
}
