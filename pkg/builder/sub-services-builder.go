package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	subserviceConstructors "micro-gen/pkg/constructors/sub-service-constructor"
)

// initializeServicesFile will initialize the services.go file to hold all
// subService information
func initializeServicesFile(serviceArray []string, microType, projectName string, hasDB bool) (err error) {
	// get the byte slice
	b, _ := subserviceConstructors.ReturnSubServiceServiceFile(serviceArray, hasDB)
	var fileEnding string
	var filePath string
	switch microType {
	case "go":
		fileEnding = ".go"
		filePath = "/pkg/sub-services/"
	}
	// set the filename
	fileName := fmt.Sprintf("%s", filePath+"services"+fileEnding)
	// write the file
	err = ioutil.WriteFile(projectName+fileName, b, 0755)
	if err != nil {
		// handle the error
		log.Fatal(err)
	}
	return nil
}

// initializeSubServiceService will create the service file for
// the desired subService in the services folder
func initializeSubServiceService(microType, name, projectName string, hasDB bool) (err error) {
	// get the byte slice
	b, _ := subserviceConstructors.ReturnSubServiceLayer(microType, name, projectName, hasDB)
	var fileEnding string
	var filePath string
	switch microType {
	case "go":
		fileEnding = "-service.go"
		filePath = "/pkg/sub-services/"
	}
	// set the filename
	fileName := fmt.Sprintf("%s", filePath+name+fileEnding)
	// write the file
	err = ioutil.WriteFile(projectName+fileName, b, 0755)
	if err != nil {
		// handle the error
		log.Fatal(err)
	}
	return nil
}
