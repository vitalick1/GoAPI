package controller

import (
	"GoAPI/service"
	"net/http"
)

type AccountController struct {
	accService service.AccountService
}

func (ac *AccountController) AccService() service.AccountService {
	return ac.accService
}

func (ac *AccountController) SetAccService(accService service.AccountService) {
	ac.accService = accService
}

func (ac AccountController) HandleAccountCreate(w http.ResponseWriter, r *http.Request) {
	ac.accService.CreateAccount("aaaa@aaa.com")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{}`))
}

func (ac AccountController) HandleAccountGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{}`))
}
