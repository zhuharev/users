package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/zhuharev/users"
	"github.com/zhuharev/users/config"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Action = run

	app.Run(os.Args)
}

func run(c *cli.Context) {
	cnf := &config.Config{}
	cnf.Database.Driver = "sqlite3"
	cnf.Database.Setting = ":memory:"
	srv, e := users.NewFromConfig(cnf)
	if e != nil {
		panic(e)
	}
	fmt.Println(srv)
}
