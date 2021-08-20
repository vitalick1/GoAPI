package repository

import (
	"GoAPI/repository"
)

type AccountRepositoryFactory struct {
}

func (arf AccountRepositoryFactory) Create() (ar repository.AccountRepository) {
	return ar
}
