package core

import "github.com/louisevanderlith/husk/validation"

type Clothing struct {
	Code        string
	Brand       string
	Description string
	Type        string
	Size        string
	Colour      string
	Material    string
	Weight      int //grams
}

func (c Clothing) Valid() error {
	return validation.Struct(c)
}
