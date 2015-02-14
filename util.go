package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
)

func PrettyJson(r interface{}) {
	v, _ := json.MarshalIndent(r, "", "    ")
	log.Info("%s", v)
}
