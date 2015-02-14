package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/jbdalido/go-marathon"
	"io/ioutil"
	"net/http"
)

//Gets information about a Marathon installation
func Info(c *cli.Context) {
	url := fmt.Sprintf("http://%s%s/info", c.GlobalString("host"), gomarathon.APIVersion)

	r, err := http.Get(url)
	if err != nil {
		fmt.Println("Error contacting marathon url: %s", err)
	}

	data := map[string]interface{}{}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error("Unable to read Info response: ", err)
	}

	json.Unmarshal(body, &data)
	j, _ := json.MarshalIndent(data, "", "   ")
	fmt.Println(string(j))
}

//Removes an app from Marathon
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

//Lists all the running apps being managed from a Marathon instance
//if arguments are supplied, only list those apps
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
