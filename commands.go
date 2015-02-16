package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/asteris-llc/gomarathon"
	"github.com/codegangsta/cli"
)

//Gets information about the Marathon cluster
func Info(host string) (*gomarathon.Response, error) {
	m, _ := MarathonClient(host)

	resp, err := m.Info()

	if err != nil {
		log.Error("Unable to get cluster information: ", err)
		return resp, err
	}
	return resp, err
}

//Gets the current Leader
func Leader(host string) (*gomarathon.Response, error) {

	m, _ := MarathonClient(host)

	resp, err := m.Leader()

	if err != nil {
		log.Fatal("Leader data error: Make sure your hosts can resolve peer hostnames: ", err)
		return nil, err
	}

	if resp.Code == 404 {
		log.Info("No leader elected")
	}

	return resp, nil
}

//Forces current leader to step down
func DeleteLeader(host string) (*gomarathon.Response, error) {
	m, _ := MarathonClient(host)

	resp, err := m.DeleteLeader()

	if err != nil {
		log.Error("Error on Delete Leader request: ", err)
		return nil, err
	}

	if resp.Code == 404 {
		log.Error("Request for leader to step down failed")
	}

	return resp, err
}

//Removes an app from Marathon
func RmApp(c *cli.Context) {
	if len(c.Args()) == 0 {
		log.Error("Please provide and app to delete")
	}

	app := c.Args()[0]
	m, err := MarathonClient(c.GlobalString("host"))

	_, err = m.DeleteApp(app)

	if err != nil {
		log.Error("Unable to delete app: ", err)
	} else {
		log.Info("Application deleted: ", app)
	}
}

//Lists all the running apps being managed from a Marathon instance
//if argument is supplied, only list that apps
func LsApps(c *cli.Context) (*gomarathon.Response, error) {
	m, _ := MarathonClient(c.GlobalString("host"))

	filter := ""

	var r *gomarathon.Response
	var err error

	if len(c.Args()) > 0 {
		filter = c.Args()[0]
		r, err = m.GetApp(filter)
	} else {
		r, err = m.ListAppsByCmd(filter)
	}

	if err != nil {
		log.Error("Error listing apps: ", err)
	}

	return r, err
}

func Ping(host string) (string, error) {
	m, err := MarathonClient(host)

	resp, err := m.Ping(host)

	if err != nil {
		log.Fatal("Ping Unsuccessful: ", err)
		return resp, err
	} else {
		log.Info(fmt.Sprintf("Ping successful, recieved '%s' from %s", resp, host))
		return resp, nil
	}

}
