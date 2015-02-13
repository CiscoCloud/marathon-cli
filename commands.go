package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/jbdalido/go-marathon"
	"net/http"
)

func LsApps(c *cli.Context) {
	m := MarathonClient(c.GlobalString("host"))

	if len(c.Args()) > 0 {
		for _, app := range c.Args() {
			r, err := m.GetApp(app)
			if err != nil {
				log.Error("App not found: ", app)
			} else {
				PrettyJson(r)
			}
		}
	} else {
		r, err := m.ListApps()

		if err != nil {
			log.Fatal("Error listing apps ", err)
		}
		PrettyJson(r)
	}
}

func Ping(c *cli.Context) {
	url := fmt.Sprintf("http://%s%s/ping", c.GlobalString("host"), gomarathon.APIVersion)

	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error contacting marathon url: %s", err)
	}

	log.Info("Ping successful: ", url)
}
