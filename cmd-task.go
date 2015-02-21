package main

// Commands dealing with tasks
import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/asteris-llc/gomarathon"
	"github.com/codegangsta/cli"
)

// List all Tasks running on the cluster
func LsTask(c *cli.Context) (*gomarathon.Response, error) {
	m, _ := MarathonClient(c.GlobalString("host"))

	var resp *gomarathon.Response
	var err error

	//if an arg is supplied, just look for tasks for that app
	if len(c.Args()) > 0 {
		resp, err = m.GetAppTasks(c.Args().First())
	} else {
		resp, err = m.ListTasks()
	}

	if err != nil {
		log.Error("Error fetching tasks: ", err)
		return nil, err
	}
	return resp, err
}

// Kill all tasks belonging to an app
// The app will not be deleted
func KillAppTasks(c *cli.Context) (*gomarathon.Response, error) {

	m, _ := MarathonClient(c.GlobalString("host"))

	var resp *gomarathon.Response
	var err error

	if len(c.Args()) > 0 {
		resp, err = m.KillTasks(c.Args().First(), "", false)
	} else {
		log.Error("Please provide an app ID")
		return nil, errors.New("No app ID provided")
	}

	if err != nil {
		log.Error("Error killing tasks: ", err)
	}

	return resp, err
}
