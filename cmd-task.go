package main

// Commands dealing with tasks
import (
	log "github.com/Sirupsen/logrus"
	"github.com/asteris-llc/gomarathon"
	"github.com/codegangsta/cli"
)

// List all Tasks running on the cluster
func LsTask(c *cli.Context) (*gomarathon.Response, error) {
	m, _ := MarathonClient(c.GlobalString("host"))

	resp, err := m.ListTasks()

	if err != nil {
		log.Error("Error fetching tasks: ", err)
		return nil, err
	}
	return resp, err
}
