package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/asteris-llc/gomarathon"
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
