package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
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
