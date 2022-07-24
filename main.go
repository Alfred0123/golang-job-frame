package main

import (
	config "golang-job-frame/configs"
	sampleController "golang-job-frame/di/sample"

	// sample "golang-job-frame/pkg/sample"
	// sampleController "golang-job-frame/pkg/sample"
	"log"
)

func init() {
	config.SetEnv()
	log.Printf("ENV: %s", config.RuntimeConf.ENV)
}

func main() {	
	// sample.Run()
	sampleHandler, error := sampleController.Wire()
	// sampleService, error2 := sampleController.WireService()
	println(error)
	// println(error2)

	sampleHandler.Run()
	// sampleService.GetSample(10)
}
