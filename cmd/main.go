package main

import (
	"flag"
	"github.com/louisevanderlith/wear/core"
	"github.com/louisevanderlith/wear/handles"
	"net/http"
	"time"
)

func main() {
	issuer := flag.String("issuer", "http://127.0.0.1:8080/auth/realms/mango", "OIDC Provider's URL")
	audience := flag.String("audience", "wear", "Token target 'aud'")
	flag.Parse()

	core.CreateContext()
	defer core.Shutdown()

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8086",
		Handler:      handles.SetupRoutes(*issuer, *audience),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
