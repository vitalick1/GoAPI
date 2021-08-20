package repository

import "GoAPI/model"

type AccountRepository struct {
}

func (ar AccountRepository) Save(acc model.Account) (a model.Account) {
	return acc
}
