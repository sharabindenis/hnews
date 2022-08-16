package models

import "github.com/upper/db/v4"

type UsersModel struct {
	db db.Session
}

func (m UsersModel) Get() {

	//m.db.Collection("users").Find(db.Cond{"id": 1}).One()
}
