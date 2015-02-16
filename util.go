package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/asteris-llc/gomarathon"
)

func Output(format string, r *gomarathon.Response, fields ...string) {
	switch format {
	case "json", "JSON":
		OutputJson(r)
	default:
		log.Fatal(fmt.Sprintf("Output format '%s' is not supported", format))
	}
}

func OutputJson(r *gomarathon.Response) {
	v, _ := json.MarshalIndent(r, "", "   ")
	fmt.Printf("%s\n", v)
}

func PrettyJson(r interface{}) {
	v, _ := json.MarshalIndent(r, "", "    ")
	log.Info(string(v))
}
