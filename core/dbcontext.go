package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
)

type WearContext interface {
	GetClothing(key hsk.Key) (Clothing, error)
	FindClothing(page, size int) (records.Page, error)
	CreateClothing(obj Clothing) (hsk.Key, error)
	UpdateClothing(key hsk.Key, obj Clothing) error
}

func CreateContext() WearContext {
	ctx = context{Clothing: husk.NewTable(Clothing{})}

	return ctx
}

func Shutdown() {
	ctx.Clothing.Save()
}

func Context() WearContext {
	return ctx
}

type context struct {
	Clothing husk.Table
}

var ctx context

func (c context) GetClothing(key hsk.Key) (Clothing, error) {
	rec, err := c.Clothing.FindByKey(key)

	if err != nil {
		return Clothing{}, err
	}

	return rec.GetValue().(Clothing), nil
}

func (c context) FindClothing(page, size int) (records.Page, error) {
	return c.Clothing.Find(page, size, op.Everything())
}

func (c context) CreateClothing(obj Clothing) (hsk.Key, error) {
	return c.Clothing.Create(obj)
}

func (c context) UpdateClothing(key hsk.Key, obj Clothing) error {
	return c.Clothing.Update(key, obj)
}
