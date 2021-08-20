package controller

import (
	"GoAPI/controller"
	"GoAPI/service"
)

type AccountControllerFactory struct {
	accountService service.AccountService
}

func (acf *AccountControllerFactory) AccountService() service.AccountService {
	return acf.accountService
}

func (acf *AccountControllerFactory) SetAccountService(accountService service.AccountService) {
	acf.accountService = accountService
}

func (acf AccountControllerFactory) Create() controller.AccountController {
	accController := controller.AccountController{}
	accController.SetAccService(acf.accountService)
	return accController
}
