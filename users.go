package users

import (
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

	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	Deleted time.Time `xorm:"deleted"`
}
