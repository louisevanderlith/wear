package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(issuer, audience string) http.Handler {
	r := mux.NewRouter()
	mw := open.BearerMiddleware(audience, issuer)
	r.Handle("/info/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewClothing))).Methods(http.MethodGet)

	r.Handle("/info/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchClothing))).Methods(http.MethodGet)

	r.Handle("/info", mw.Handler(http.HandlerFunc(CreateClothing))).Methods(http.MethodPost)
	r.Handle("/info/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdateClothing))).Methods(http.MethodPut)

	/*

		//clothing
		r.Handle("/clothes", mw.Handler(http.HandlerFunc(GetClothing))).Methods(http.MethodGet)
		r.Handle("/clothes/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewClothing))).Methods(http.MethodGet)
		r.Handle("/clothes/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchClothing))).Methods(http.MethodGet)
		r.Handle("/clothes/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchClothing))).Methods(http.MethodGet)
		r.Handle("/clothes", mw.Handler(http.HandlerFunc(CreateClothing))).Methods(http.MethodPost)
		r.Handle("/clothes/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdateClothing))).Methods(http.MethodPut)
	*/

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
