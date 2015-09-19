package xorm

import (
	_ "github.com/mattn/go-sqlite3"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zhuharev/users"
	"testing"
)

func TestMem(t *testing.T) {
	Convey("testing mem", t, func() {
		s := new(Store)
		s.Driver = "sqlite3"
		e := s.Connect(":memory:")
		if e != nil {
			t.Error(e)
		}

		u := new(users.User)
		u.Name = "kirill"

		e = s.Save(u)
		if e != nil {
			t.Error(e)
		}

		l, e := s.eng.Count(&users.User{})
		if e != nil {
			t.Error(e)
		}
		So(l, ShouldEqual, 1)
	})
}
