package models 

import (
	"time"
	"database/sql"
)

type user struct{
	ID int
	Name string
	Email string
	Hash_password []byte
	Create time.Time
}

type UserModel struct{
	DB *sql.DB
}

func (m *UserModel) Insert(name,email,password string) error{
	return nil
}

func (m *UserModel) Authenticate(email,password string)(int, error){
	return 0,nil
}

func (m *UserModel) Exists(id int)(bool,error){
	return false,nil
}