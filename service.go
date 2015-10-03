package users

import (
	"github.com/zhuharev/users/config"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Store  Store
	Config *config.Config
	KV     *KV
}

func New(filepath string) (*Service, error) {
	cnf, e := config.NewFromFile(filepath)
	if e != nil {
		return nil, e
	}
	return NewFromConfig(cnf)
}

func NewFromConfig(cnf *config.Config) (*Service, error) {
	s := new(Service)
	s.Config = cnf

	store, e := NewStore(cnf.Database.Driver, cnf.Database.Setting)
	if e != nil {
		return nil, e
	}
	s.Store = store

	kv, e := NewKV(cnf.Database.KVPath)
	if e != nil {
		return nil, e
	}
	s.KV = kv

	return s, nil
}

func (s *Service) IsExistUserName(username string) (bool, error) {
	u, e := s.Store.GetByUserName(username)
	if e == ErrNotFound || u == nil {
		return false, nil
	}
	if e != nil {
		return false, e
	}
	return true, nil
}

func (s *Service) CreateUser(username, password string) (*User, error) {
	if ok, e := s.IsExistUserName(username); ok {
		return nil, ErrUsernameAlreadyExists
	} else {
		if e != nil {
			return nil, e
		}
	}
	u := new(User)
	u.Name = username
	u.Data = map[string]interface{}{}
	hashedPassword, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if e != nil {
		return nil, e
	}
	u.HashedPassword = hashedPassword
	e = s.Store.Save(u)
	return u, e
}

func (s *Service) MakeAdminPermission(uid int64) error {
	u, e := s.Store.Get(uid)
	if e != nil {
		return e
	}
	u.Status = u.Status.Add(Admin)
	e = s.Store.Save(u)
	if e != nil {
		return e
	}
	return nil
}

func (s *Service) DeleteAdminPermission(uid int64) error {
	u, e := s.Store.Get(uid)
	if e != nil {
		return e
	}
	u.Status = u.Status.Remove(Admin)
	e = s.Store.Save(u)
	if e != nil {
		return e
	}
	return nil
}
