package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/wear/core"
	"log"
	"net/http"
)

func GetTypes(w http.ResponseWriter, r *http.Request) {
	result, err := core.Context().FindTypes(1, 10)

	if err != nil {
		log.Println("Get Categories Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func SearchTypes(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)
	result, err := core.Context().FindTypes(page, size)

	if err != nil {
		log.Println("Search Categories Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func ViewType(w http.ResponseWriter, r *http.Request) {
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println("Parse Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	obj, err := core.Context().GetType(key)

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

func CreateType(w http.ResponseWriter, r *http.Request) {
	obj := core.Type{}

	err := drx.JSONBody(r, &obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	result, err := core.Context().CreateType(obj)

	if err != nil {
		log.Println("Create Type Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func UpdateType(w http.ResponseWriter, r *http.Request) {
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	obj := core.Type{}
	err = drx.JSONBody(r, &obj)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = core.Context().UpdateType(key, obj)

	if err != nil {
		log.Println("Update Type Error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON("Saved"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
