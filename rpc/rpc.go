package rpc

import (
	"github.com/zhuharev/users"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Users struct {
	us *users.Service
}

func New(filepath string) (*Users, error) {
	us, e := users.New(filepath)
	if e != nil {
		return nil, e
	}

	s := new(Users)
	s.us = us
	return s, nil
}

func Run(s *Users) {
	rpc.Register(s)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

func Block() {
	for {
	}
}

type GetRequest struct {
	Id int64
}

type Error struct {
	Message string
}

type GetResponse struct {
	Error Error
	User  *users.User
}

func (s *Users) Get(r *GetRequest, resp *GetResponse) error {
	resp.User = new(users.User)

	u, e := s.us.Store.Get(r.Id)
	if e != nil {
		return e
	}
	resp.User = u

	return nil
}
