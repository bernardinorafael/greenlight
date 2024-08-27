package main

import (
  "net/http"
  
  "github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
  router := httprouter.New()
  
  router.NotFound = http.HandlerFunc(app.notFoundResponse)
  router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
  
  router.HandlerFunc(http.MethodGet, "/v1/health-check", app.healthCheck)
  router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovie)
  router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovie)
  
  return app.recoverPanic(router)
}
