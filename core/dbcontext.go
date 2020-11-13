package core

import (
	"encoding/json"
	"fmt"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/collections"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
	"os"
	"reflect"
)

type WearContext interface {
	GetClothing(k hsk.Key) (Clothing, error)
	FindClothing(page, size int) (records.Page, error)
	CreateClothing(obj Clothing) (hsk.Key, error)
	UpdateClothing(k hsk.Key, obj Clothing) error
	GetTypeSizes(k hsk.Key) (map[string]string, error)
	GetBrand(k hsk.Key) (Brand, error)
	FindBrands(page, size int) (records.Page, error)
	CreateBrand(obj Brand) (hsk.Key, error)
	UpdateBrand(k hsk.Key, obj Brand) error
	GetType(k hsk.Key) (Type, error)
	FindTypes(page, size int) (records.Page, error)
	CreateType(obj Type) (hsk.Key, error)
	UpdateType(k hsk.Key, obj Type) error
}

func CreateContext() WearContext {
	ctx = context{
		Clothing: husk.NewTable(Clothing{}),
		Brands:   husk.NewTable(Brand{}),
		Types:    husk.NewTable(Type{}),
	}

	seedBrands()
	seedTypes()

	return ctx
}

func seedBrands() {
	f, err := os.Open("db/brands.seed.json")

	if err != nil {
		panic(err)
	}

	var items []Brand
	dec := json.NewDecoder(f)
	err = dec.Decode(&items)

	if err != nil {
		panic(err)
	}

	coll := collections.ReadOnlyList(reflect.ValueOf(items))

	err = ctx.Brands.Seed(coll)

	if err != nil {
		panic(err)
	}
}

func seedTypes() {
	f, err := os.Open("db/types.seed.json")

	if err != nil {
		panic(err)
	}

	var items []Type
	dec := json.NewDecoder(f)
	err = dec.Decode(&items)

	if err != nil {
		panic(err)
	}

	coll := collections.ReadOnlyList(reflect.ValueOf(items))

	err = ctx.Types.Seed(coll)

	if err != nil {
		panic(err)
	}
}

func Shutdown() {
	ctx.Clothing.Save()
	ctx.Brands.Save()
	ctx.Types.Save()
}

func Context() WearContext {
	return ctx
}

var ctx context

type context struct {
	Clothing husk.Table
	Brands   husk.Table
	Types    husk.Table
}

func (c context) GetTypeSizes(k hsk.Key) (map[string]string, error) {
	t, err := c.GetType(k)

	if err != nil {
		return nil, err
	}

	result := make(map[string]string)

	for _, s := range t.Sizes {
		result[s.ShortCode] = fmt.Sprintf("%s - %s", s.ShortCode, s.Description)
	}

	return result, nil
}

func (c context) GetBrand(k hsk.Key) (Brand, error) {
	rec, err := c.Brands.FindByKey(k)

	if err != nil {
		return Brand{}, err
	}

	return rec.GetValue().(Brand), nil
}

func (c context) FindBrands(page, size int) (records.Page, error) {
	return c.Brands.Find(page, size, op.Everything())
}

func (c context) CreateBrand(obj Brand) (hsk.Key, error) {
	return c.Brands.Create(obj)
}

func (c context) UpdateBrand(k hsk.Key, obj Brand) error {
	return c.Brands.Update(k, obj)
}

func (c context) GetType(k hsk.Key) (Type, error) {
	rec, err := c.Types.FindByKey(k)

	if err != nil {
		return Type{}, err
	}

	return rec.GetValue().(Type), nil
}

func (c context) FindTypes(page, size int) (records.Page, error) {
	return c.Types.Find(page, size, op.Everything())
}

func (c context) CreateType(obj Type) (hsk.Key, error) {
	return c.Types.Create(obj)
}

func (c context) UpdateType(k hsk.Key, obj Type) error {
	return c.Types.Update(k, obj)
}

func (c context) GetClothing(k hsk.Key) (Clothing, error) {
	rec, err := c.Clothing.FindByKey(k)

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

func (c context) UpdateClothing(k hsk.Key, obj Clothing) error {
	return c.Clothing.Update(k, obj)
}
