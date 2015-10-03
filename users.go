package users

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id   int64
	Name string `xorm:"unique index"`

	FirstName, LastName, Patronymic string

	Phone string
	Email string `xorm:"unique index"`

	Status Status

	HashedPassword []byte

	Data UserData

	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	Deleted time.Time `xorm:"deleted"`
}

type UserData map[string]interface{}

func (u *User) ValidatePassword(password string) bool {
	e := bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(password))
	if e != nil {
		return false
	}
	return true
}
