package main

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/asteris-llc/gomarathon"
	"github.com/codegangsta/cli"
	"io/ioutil"
)

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
		log.Error("No application provided to delete")
		return nil, errors.New("Please provide the id of an app to delete")
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
