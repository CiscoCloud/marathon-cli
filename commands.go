package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/jbdalido/go-marathon"
	"net/http"
)

func RmApp(c *cli.Context) {
	if len(c.Args()) == 0 {
		log.Error("Please provide and app to delete")
	}

	app := c.Args()[0]
	m := MarathonClient(c.GlobalString("host"))

	_, err := m.DeleteApp(app)

	if err != nil {
		log.Error("Unable to delete app: ", err)
	} else {
		log.Info("Application deleted: ", app)
	}
}

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
