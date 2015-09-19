package leveldb

/*
import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/zhuharev/users"
	"github.com/zhuharev/users/store"
)

type Store struct {
	ldb *leveldb.DB
}

func (s *Store) Name() string {
	return "leveldb"
}

func (s *Store) Connect(setting string) error {
	ldb, e := leveldb.OpenFile(setting, nil)
	if e != nil {
		return e
	}
	s.ldb = ldb
	return nil
}

func (s *Store) Save(*users.User) error {
	return nil
}

func (s *Store) Get(int64) (*users.User, error) {
	return nil, nil
}

func (s *Store) Read(int64, ...int64) ([]*users.User, error) {
	return nil, nil
}

func (s *Store) Delete(int64) error {
	return nil
}

func init() {
	store.Register("leveldb", &Store{})
}
*/
