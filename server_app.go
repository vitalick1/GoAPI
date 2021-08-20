package main

import (
	"GoAPI/factory/controller"
	"github.com/gorilla/mux"
	"net/http"
)

type ServerApp struct {
	accountControllerFactory controller.AccountControllerFactory
	r                        *mux.Router
}

func (app *ServerApp) AccountRouterFactory() controller.AccountControllerFactory {
	return app.accountControllerFactory
}

func (app *ServerApp) SetAccControllerFactory(accountControllerFactory controller.AccountControllerFactory) {
	app.accountControllerFactory = accountControllerFactory
}

func (app *ServerApp) Router() *mux.Router {
	return app.r
}

func (app *ServerApp) SetRouter(r *mux.Router) {
	app.r = r
}

func (app ServerApp) Init() {
	app.r.HandleFunc("/health", app.healthCheck).Methods(http.MethodGet)
	app.r.HandleFunc("/health", app.notFound)

	api := app.r.PathPrefix("/accounts/v1").Subrouter()
	accountController := app.accountControllerFactory.Create()
	api.HandleFunc("", accountController.HandleAccountCreate).Methods(http.MethodPost)
	api.HandleFunc("", accountController.HandleAccountGet).Methods(http.MethodGet)
	api.HandleFunc("", app.notFound)
}

func (app ServerApp) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{}`))
}

func (app ServerApp) notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{}`))
}
