package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/wear/core"
	"log"
	"net/http"
)

func GetBrands(w http.ResponseWriter, r *http.Request) {
	result, err := core.Context().FindBrands(1, 10)

	if err != nil {
		log.Println("Get Brands Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func SearchBrands(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)
	result, err := core.Context().FindBrands(page, size)

	if err != nil {
		log.Println("Search Brands Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func ViewBrand(w http.ResponseWriter, r *http.Request) {
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println("Parse Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	obj, err := core.Context().GetBrand(key)

	if err != nil {
		log.Println("Get Category Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(obj))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func CreateBrand(w http.ResponseWriter, r *http.Request) {
	obj := core.Brand{}

	err := drx.JSONBody(r, &obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	result, err := core.Context().CreateBrand(obj)

	if err != nil {
		log.Println("Create Brand Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func UpdateBrand(w http.ResponseWriter, r *http.Request) {
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	obj := core.Brand{}
	err = drx.JSONBody(r, &obj)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = core.Context().UpdateBrand(key, obj)

	if err != nil {
		log.Println("Update Vehicle Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON("Saved"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
