package model

type Account struct {
	id    string
	email string
}

func (a Account) SetId(id string) Account {
	a.id = id
	return a
}

func (a Account) SetEmail(email string) Account {
	a.email = email
	return a
}
