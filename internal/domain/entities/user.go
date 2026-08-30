package entities

import "uuid"

type User struct {
	UUID     *uuid.UUID
	ID       *uint
	Name     string
	Email    string
	Password string
}

func (u User) GetUUID() *uuid.UUID {
	return u.UUID
}

func (u User) GetID() *uint {
	return u.ID
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetPassword() string {
	return u.Password
}
