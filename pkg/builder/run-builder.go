package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	runModels "micro-gen/pkg/models/run-models"
)

// initializeRun will initialize the files within the run folder
func initializeRun(projectName string, serviceArray []string, hasDB bool) (err error) {

	err = initializeMainFile(projectName, serviceArray, hasDB)
	if err != nil {
		//handle err
	}
	if hasDB {
		err = initializeConfigFile(projectName)
		if err != nil {
			//handle err
		}
	}
	err = initializeSecretFile(projectName)
	if err != nil {
		//handle err
	}
	return nil
}

// initializeMainFile will initialize the main.go file and write
// the required items within the file
func initializeMainFile(projectName string, serviceArray []string, hasDB bool) (err error) {
	// get the byte slice
	b, _ := runModels.ReturnMainFile("go", serviceArray, hasDB, projectName)
	// set the filename
	fileName := fmt.Sprintf("/run/main.go")
	// write the file
	err = ioutil.WriteFile(projectName+fileName, b, 0755)
	if err != nil {
		// handle the error
		log.Fatal(err)
	}
	return nil
}

// initializeConfigFile will initialize the config.go file and write
// the required configurations within the file
func initializeConfigFile(projectName string) (err error) {
	// get the byte slice
	b, _ := runModels.ReturnConfigFile("go")
	// set the filename
	fileName := fmt.Sprintf("/run/config.go")
	// write the file
	err = ioutil.WriteFile(projectName+fileName, b, 0755)
	if err != nil {
		// handle the error
		log.Fatal(err)
	}
	return nil
}

// initializeSecretFile will initialize the secretStuff.go file and write
// the required configurations within the file
func initializeSecretFile(projectName string) (err error) {
	// get the byte slice
	b, _ := runModels.ReturnSecretStuffFile("go")
	// set the filename
	fileName := fmt.Sprintf("/run/secretStuff.go")
	// write the file
	err = ioutil.WriteFile(projectName+fileName, b, 0755)
	if err != nil {
		// handle the error
		log.Fatal(err)
	}
	return nil
}
