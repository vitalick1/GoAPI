package service

import (
	modelFactory "GoAPI/factory/model"
	"GoAPI/model"
	"GoAPI/repository"
	"github.com/google/uuid"
)

type AccountService struct {
	accRepo    repository.AccountRepository
	accFactory modelFactory.AccountFactory
	uuid       uuid.UUID
}

func (as *AccountService) AccRepo() repository.AccountRepository {
	return as.accRepo
}

func (as *AccountService) SetAccRepo(accRepo repository.AccountRepository) {
	as.accRepo = accRepo
}

func (as *AccountService) AccFactory() modelFactory.AccountFactory {
	return as.accFactory
}

func (as *AccountService) SetAccFactory(accFactory modelFactory.AccountFactory) {
	as.accFactory = accFactory
}

func (as *AccountService) Uuid() uuid.UUID {
	return as.uuid
}

func (as *AccountService) SetUuid(uuid uuid.UUID) {
	as.uuid = uuid
}

func (as AccountService) CreateAccount(email string) model.Account {
	acc := as.accFactory.Create(as.uuid.String(), email)
	as.accRepo.Save(acc)
	return acc
}
