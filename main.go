package main

import (
	"GoAPI/factory/controller"
	"GoAPI/factory/model"
	"GoAPI/factory/repository"
	serviceFactory "GoAPI/factory/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	accFactory := model.AccountFactory{}
	accRepoFactory := repository.AccountRepositoryFactory{}

	accountServiceFactory := serviceFactory.AccountServiceFactory{}
	accountServiceFactory.SetAccRepo(accRepoFactory.Create())
	accountServiceFactory.SetAccFactory(accFactory)
	accountServiceFactory.SetUuid(uuid.UUID{})

	accountControllerFactory := controller.AccountControllerFactory{}
	accountControllerFactory.SetAccountService(accountServiceFactory.Create())

	app := ServerApp{}
	app.SetRouter(r)
	app.SetAccControllerFactory(accountControllerFactory)
	app.Init()

	log.Fatal(http.ListenAndServe(":8081", r))
}
