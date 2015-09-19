package config

import (
	"gopkg.in/gcfg.v1"
)

type Config struct {
	App struct {
		Secret string
	}
	Web struct {
		Host string
	}
	Database struct {
		Driver, Setting, KVPath string
	}
	Mail struct {
		Sender, Password string
		Server           string
		Port             int
	}
}

func NewFromFile(filepath string) (*Config, error) {
	cnf := new(Config)
	e := gcfg.ReadFileInto(cnf, filepath)
	return cnf, e
}
