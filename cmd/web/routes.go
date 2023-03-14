// Filename: cmd/web/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func(app *application) routes() *httprouter.Router {
	router := httprouter.New()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// ROUTES: 10
  router.Handler(http.MethodGet,"/static/*filepath", http.StripPrefix("/static", fileServer))
	router.HandlerFunc(http.MethodGet,"/quote/create", app.quoteCreateShow)
	router.HandlerFunc(http.MethodGet,"/quote/create", app.quoteCreateSubmit)



	
	return router
}