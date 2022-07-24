package main

import (
	config "golang-job-frame/configs"
	sample "golang-job-frame/pkg/sample"
	"log"
)

func init() {
	config.SetEnv()
	log.Printf("ENV: %s", config.RuntimeConf.ENV)
}

func main() {	
	sample.Test()
}
