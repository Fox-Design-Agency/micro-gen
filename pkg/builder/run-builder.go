package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	runConstructors "micro-gen/pkg/constructors/run-constructor"
)

// initializeRun will initialize the files within the run folder
func initializeRun(microType, projectName string, serviceArray, routeHandlerArray []string, hasDB bool) (err error) {

	err = initializeMainFile(microType, projectName, serviceArray, routeHandlerArray, hasDB)
	if err != nil {
		//handle err
	}
	if hasDB {
		err = initializeConfigFile(microType, projectName)
		if err != nil {
			//handle err
		}
	}
	err = initializeSecretFile(microType, projectName)
	if err != nil {
		//handle err
	}
	return nil
}

// initializeMainFile will initialize the main.go file and write
// the required items within the file
func initializeMainFile(microType, projectName string, serviceArray, routeHandlerArray []string, hasDB bool) (err error) {
	// get the byte slice
	b, _ := runConstructors.ReturnMainFile(microType, serviceArray, routeHandlerArray, hasDB, projectName)
	var fileEnding string
	var filePath string
	switch microType {
	case "go":
		fileEnding = "main.go"
		filePath = "/run/"
	}
	// set the filename
	fileName := fmt.Sprintf("%s", filePath+fileEnding)
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
func initializeConfigFile(microType, projectName string) (err error) {
	switch microType {
	case "go":
		// get the byte slice
		b, _ := runConstructors.ReturnConfigFile(microType)
		// set the filename
		fileName := fmt.Sprintf("/run/config.go")
		// write the file
		err = ioutil.WriteFile(projectName+fileName, b, 0755)
		if err != nil {
			// handle the error
			log.Fatal(err)
		}
		return nil
	default:
		return nil
	}
}

// initializeSecretFile will initialize the secretStuff.go file and write
// the required configurations within the file
func initializeSecretFile(microType, projectName string) (err error) {

	switch microType {
	case "go":
		// get the byte slice
		b, _ := runConstructors.ReturnSecretStuffFile(microType)
		// set the filename
		fileName := fmt.Sprintf("/run/secretStuff.go")
		// write the file
		err = ioutil.WriteFile(projectName+fileName, b, 0755)
		if err != nil {
			// handle the error
			log.Fatal(err)
		}
		return nil
	default:
		return nil
	}
}
