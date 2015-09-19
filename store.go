package users

import (
	"fmt"
	"sync"
)

type Store interface {
	Name() string
	Connect(string) error

	Save(*User) error
	Get(int64) (*User, error)
	Read(int64, ...int64) ([]*User, error)
	Delete(int64) error

	GetByUserName(string) (*User, error)
	GetByEmail(string) (*User, error)

	Count() (int64, error)
}

var (
	stores = map[string]Store{}
	smu    sync.Mutex
)

var (
	ErrNotFound              = fmt.Errorf("not found")
	ErrUsernameAlreadyExists = fmt.Errorf("user already exists")
)

func Register(name string, driver Store) {
	smu.Lock()
	defer smu.Unlock()
	stores[name] = driver
}

func NewStore(name string, setting string) (Store, error) {
	if driver, ok := stores[name]; !ok {
		panic(fmt.Sprintf("driver %s not found, %s", name, stores))
	} else {
		e := driver.Connect(setting)
		if e != nil {
			return nil, e
		}
		return driver, nil
	}
	panic("unrechable")
}
