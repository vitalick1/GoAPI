package model

import "GoAPI/model"

type AccountFactory struct {
}

func (af AccountFactory) Create(id string, email string) (a model.Account) {
	a.SetId(id).SetEmail(email)
	return a
}
