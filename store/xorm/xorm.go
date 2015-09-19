package xorm

import (
	"github.com/go-xorm/xorm"
	"github.com/zhuharev/users"
	"os"
)

type Store struct {
	Driver string
	eng    *xorm.Engine
}

func (s *Store) Name() string {
	return "xorm"
}

func (s *Store) Connect(setting string) error {
	if s.Driver == "" {
		panic("not driver setted")
	}
	eng, e := xorm.NewEngine(s.Driver, setting)
	if e != nil {
		return e
	}
	e = eng.Sync(&users.User{})
	if e != nil {
		return e
	}
	f, e := os.Create("sql.log")
	if e != nil {
		return e
	}
	eng.Logger = xorm.NewSimpleLogger(f)
	s.eng = eng
	return nil
}

func (s *Store) Save(u *users.User) error {
	if u.Id == 0 {
		_, e := s.eng.Insert(u)
		if e != nil {
			return e
		}
	} else {
		_, e := s.eng.Id(u.Id).Update(u)
		if e != nil {
			return e
		}
	}
	return nil
}

func (s *Store) Get(id int64) (*users.User, error) {
	u := new(users.User)
	has, e := s.eng.Id(id).Get(u)
	if e != nil {
		return nil, e
	}
	if !has {
		return nil, users.ErrNotFound
	}
	return u, nil
}

func (s *Store) Read(int64, ...int64) ([]*users.User, error) {
	return nil, nil
}

func (s *Store) Delete(int64) error {
	return nil
}

func (s *Store) Count() (int64, error) {
	return s.eng.Count(new(users.User))
}

func init() {
	users.Register("xorm", &Store{})
	users.Register("sqlite3", &Store{Driver: "sqlite3"})
}
