package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/jbdalido/go-marathon"
)

func PrettyJson(r *gomarathon.Response) {
	v, _ := json.MarshalIndent(r, "", "    ")
	log.Printf("%s", v)
}
