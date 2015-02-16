package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/asteris-llc/gomarathon"
	"github.com/codegangsta/cli"
	"encoding/json"
	"io/ioutil"
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

// Creates an app from a JSON template
func MkApp(c *cli.Context) (*gomarathon.Response, error) {

	file := c.String("file")
	dat, err := ioutil.ReadFile(file)

  if err != nil {
    log.Error(fmt.Sprintf("Unable to read file %s: %s", file, err)) 
    return nil, err
  }

	var app *gomarathon.Application

	err = json.Unmarshal(dat, &app)

  if err != nil {
    log.Error("Unable to parse json config: ", err)
  }

	m, _ := MarathonClient(c.GlobalString("host"))

	resp, err := m.CreateApp(app)

  if err != nil {
    log.Error("Unable to launch app: ", err) 
  }

  if resp.Code == 201 { 
    log.Info("Application deployed app: ", app.ID) 
  }

	return resp, err
}

//Removes an app from Marathon
func RmApp(c *cli.Context) (*gomarathon.Response, error) {
	if len(c.Args()) == 0 {
		log.Error("Please provide the id of an app to delete")
	}

	app := c.Args()[0]
	m, _ := MarathonClient(c.GlobalString("host"))

	r, err := m.DeleteApp(app)

	if err != nil {
		log.Error("Unable to delete app: ", err)
	} else {
		log.Info("Application deleted: ", app)
	}
	return r, err
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
