package main

import (
        "fmt"
        log "github.com/Sirupsen/logrus"
        "github.com/jbdalido/go-marathon"
)

func MarathonClient(host string) (*gomarathon.Client) {
        url := fmt.Sprintf("http://%s", host)
        m, err := gomarathon.NewClient(url, nil)
        if err != nil {
                log.Fatal("Marathon client error: ", err)
        }
        return m
}
