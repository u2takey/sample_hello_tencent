package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ianschenck/envflag"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// 创建一个命令行工具
func main() {
	envflag.Parse()

	app := cli.NewApp()
	app.Name = "hello"
	app.Version = "0.01"
	app.Usage = "hello command line utility"
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		ServerCommand,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// ServerCommand ...
var ServerCommand = cli.Command{
	Name:   "server",
	Usage:  "starts the kubegate server daemon",
	Action: server,
	Flags: []cli.Flag{
		cli.BoolFlag{
			EnvVar: "DEBUG",
			Name:   "debug",
			Usage:  "start the server in debug mode",
		},
		cli.StringFlag{
			EnvVar: "SERVER_ADDR",
			Name:   "server-addr",
			Usage:  "server address",
			Value:  ":8808",
		},
		cli.StringFlag{
			EnvVar: "DB_CONN_STR",
			Name:   "dbconnstr",
			Usage:  "dbconnstr",
			Value:  "",
		},
	},
}

// db: test
// table :
// CREATE TABLE `visitor` (
//     `count` int(11) unsigned NOT NULL
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
// mysql> SELECT count FROM visitor LIMIT 1;
var db *sqlx.DB

// 创建一个server, 这里用了gin 作为rest框架
func server(c *cli.Context) error {
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	if c.String("dbconnstr") != "" {
		var err error
		db, err = sqlx.Connect("mysql", c.String("dbconnstr"))
		if err != nil {
			logrus.Warnln(err)
		}
	}

	e := gin.Default()
	e.GET("/hello", hello)

	// start the server
	return http.ListenAndServe(
		c.String("server-addr"),
		e,
	)
}

func hello(c *gin.Context) {
	if db == nil {
		c.String(200, "Hello!")
		return
	}

	var count int
	err := db.Get(&count, `SELECT count FROM visitor LIMIT 1`)
	if err != nil {
		c.AbortWithError(500, err)
	}

	c.String(200, fmt.Sprintf("Hello! You are the %vth Visitor", count))
	go func() {
		_, err = db.Exec(`UPDATE visitor SET count = count+1`)
		if err != nil {
			logrus.Warnln(err)
		}
	}()

	return

}
