package service

import (
	modelFactory "GoAPI/factory/model"
	"GoAPI/repository"
	"GoAPI/service"
	"github.com/google/uuid"
)

type AccountServiceFactory struct {
	accRepo    repository.AccountRepository
	accFactory modelFactory.AccountFactory
	uuid       uuid.UUID
}

func (accServiceFactory *AccountServiceFactory) AccRepo() repository.AccountRepository {
	return accServiceFactory.accRepo
}

func (accServiceFactory *AccountServiceFactory) SetAccRepo(accRepo repository.AccountRepository) {
	accServiceFactory.accRepo = accRepo
}

func (accServiceFactory *AccountServiceFactory) AccFactory() modelFactory.AccountFactory {
	return accServiceFactory.accFactory
}

func (accServiceFactory *AccountServiceFactory) SetAccFactory(accFactory modelFactory.AccountFactory) {
	accServiceFactory.accFactory = accFactory
}

func (accServiceFactory *AccountServiceFactory) Uuid() uuid.UUID {
	return accServiceFactory.uuid
}

func (accServiceFactory *AccountServiceFactory) SetUuid(uuid uuid.UUID) {
	accServiceFactory.uuid = uuid
}

func (accServiceFactory AccountServiceFactory) Create() service.AccountService {
	as := service.AccountService{}
	as.SetAccRepo(accServiceFactory.accRepo)
	as.SetUuid(accServiceFactory.uuid)
	as.SetAccFactory(accServiceFactory.accFactory)
	return as
}
