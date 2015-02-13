package main

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"os"
)

const (
	VERSION = "0.1.0"
	AUTHOR  = "Steven Borrelli"
	EMAIL   = "steve@aster.is"
)

func run_cli() {
	app := cli.NewApp()
	app.Name = "marathon-cli"
	app.Usage = "Mange marathon apps and groups"
	app.Version = VERSION
	app.Author = AUTHOR
	app.Email = EMAIL

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host",
			Value:  "localhost:8080",
			Usage:  "Marathon host (default localhost:8080)",
			EnvVar: "MARATHON_HOST",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "lsapp",
			Usage:  "List all or one running app",
			Action: LsApps,
		},

		{
			Name:   "ping",
			Usage:  "Test Marathon connection",
			Action: Ping,
		},
	}

	app.Run(os.Args)
}

func main() {
	run_cli()
}
