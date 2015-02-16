package main

import (
	log "github.com/Sirupsen/logrus"
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
		cli.StringFlag{
			Name:  "format",
			Value: "json",
			Usage: "Output format (json)",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ping",
			Usage: "Test Marathon connection",
			Action: func(c *cli.Context) {
				_, _ = Ping(c.GlobalString("host"))
			},
		}, {
			Name:  "info",
			Usage: "Gets information about the Marathon cluster",
			Action: func(c *cli.Context) {
				r, _ := Info(c.GlobalString("host"))
				Output(c.GlobalString("format"), r)
			},
		},
		{
			Name:  "leader",
			Usage: "Gets the current leader",
			Action: func(c *cli.Context) {
				r, _ := Leader(c.GlobalString("host"))
				Output(c.GlobalString("format"), r)
			},
		},
		{
			Name:  "rmleader",
			Usage: "Forces current leader to abdicate",
			Action: func(c *cli.Context) {
				r, _ := DeleteLeader(c.GlobalString("host"))
				Output(c.GlobalString("format"), r)
			},
		},
		{
			Name:  "app",
			Usage: "Launches a new app",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "json App config file",
				},
			},
			Action: func(c *cli.Context) {
				if c.String("file") == "" {
					log.Fatal("Please provide a --file")
				}
				r, err := MkApp(c)
				if err == nil {
					Output(c.GlobalString("format"), r)
				}
			},
		},
		{
			Name:  "lsapp",
			Usage: "List all apps, or list a single <appId> ",
			Action: func(c *cli.Context) {
				r, _ := LsApps(c)
				Output(c.GlobalString("format"), r)
			},
		},
		{
			Name:  "rmapp",
			Usage: "Deletes app <appId>",
			Action: func(c *cli.Context) {
				_, _ = RmApp(c)
			},
		},
		{
			Name:  "lstask",
			Usage: "List all running tasks in the cluster",
			Action: func(c *cli.Context) {
				r, err := LsTask(c)
				if err == nil {
					Output(c.GlobalString("format"), r)
				}
			},
		},
	}

	app.Run(os.Args)
}

func main() {
	run_cli()
}
