package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zhuharev/users"
	_ "github.com/zhuharev/users/store/xorm"
)

func main() {
	srv, e := users.New("cnf")
	if e != nil {
		panic(e)
	}
	fmt.Println(srv)
	u, e := srv.CreateUser("kirill", "123")
	if e != nil {
		panic(e)
	}

	l, e := srv.Store.Count()
	if e != nil {
		panic(e)
	}
	if l != 1 {
		panic("len should be 1")
	}

	u.FirstName = "kirill"
	e = srv.Store.Save(u)
	if e != nil {
		panic(e)
	}

	u, e = srv.CreateUser("kirill", "123")
	if e != users.ErrUsernameAlreadyExists {
		panic("should be error")
	}

	u, e = srv.Store.GetByUserName("kirill")
	if e != nil {
		panic(e)
	}

	_, e = srv.SendConfirmEmail(u)
	if e != users.ErrInvalidEmail {
		panic("should be error")
	}

	if u == nil {
		panic("useris nil")
	}
	u.Email = "kirill@zhuharev.ru"
	cf, e := srv.SendConfirmEmail(u)
	if e != nil {
		panic(e)
	}
	e = srv.ConfirmEmail(u.Email, cf.Code, cf.Hash())
	if e != nil {
		panic(e)
	}
}
