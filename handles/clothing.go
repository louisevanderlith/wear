package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/wear/core"
	"log"
	"net/http"
)

func GetClothing(w http.ResponseWriter, r *http.Request) {
	results, err := core.Context().FindClothing(1, 10)

	if err != nil {
		log.Println("Find Clothing Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func ViewClothing(w http.ResponseWriter, r *http.Request) {
	k := drx.FindParam(r, "key")
	key, err := keys.ParseKey(k)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := core.Context().GetClothing(key)

	if err != nil {
		log.Println("Get Service Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(rec))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func SearchClothing(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)

	results, err := core.Context().FindClothing(page, size)

	if err != nil {
		log.Println("Find Parts Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func CreateClothing(w http.ResponseWriter, r *http.Request) {
	var obj core.Clothing
	err := drx.JSONBody(r, &obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := core.Context().CreateClothing(obj)

	if err != nil {
		log.Println("Create Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(rec))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func UpdateClothing(w http.ResponseWriter, r *http.Request) {
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	obj := core.Clothing{}
	err = drx.JSONBody(r, &obj)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = core.Context().UpdateClothing(key, obj)

	if err != nil {
		log.Println("Update Service Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(nil))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
