package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"net/http"
	"os"
)

const (
	VERSION	= "0.1.0"
        AUTHOR = "Steven Borrelli"
        EMAIL = "steve@aster.is"
  
        API = "v2"
)

func LsApps(c *cli.Context) {
	log.Warn("Not implemented")
}

func Ping(c *cli.Context) {
        url := fmt.Sprintf("http://%s/%s/ping", c.GlobalString("host"), API)

	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error contacting marathon url: %s", err)
	}

	log.Info("Ping successful: ", url)
}

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
			Name:   "lsapps",
			Usage:  "List running apps",
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
