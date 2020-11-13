package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
)

type Clothing struct {
	Code        string `hsk:"size(64)"`
	Brand       hsk.Key
	Description string
	Type        hsk.Key
	Size        string
	Colour      string
	Material    string
	Weight      int //grams
}

func (c Clothing) Valid() error {
	return validation.Struct(c)
}
