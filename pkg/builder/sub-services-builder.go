package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	subserviceModels "micro-gen/pkg/models/sub-service-models"
)

// initializeServicesFile will initialize the services.go file to hold all
// subService information
func initializeServicesFile(serviceArray []string, projectName string, hasDB bool) (err error) {
	// get the byte slice
	b, _ := subserviceModels.ReturnSubServiceServiceFile(serviceArray, hasDB)
	// set the filename
	fileName := fmt.Sprintf("/pkg/sub-services/services.go")
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
func initializeSubServiceService(name, projectName string, hasDB bool) (err error) {
	// get the byte slice
	b, _ := subserviceModels.ReturnSubServiceLayer("go", name, projectName, hasDB)
	// set the filename
	fileName := fmt.Sprintf("/pkg/sub-services/%s", name+"-service.go")
	// write the file
	err = ioutil.WriteFile(projectName+fileName, b, 0755)
	if err != nil {
		// handle the error
		log.Fatal(err)
	}
	return nil
}
