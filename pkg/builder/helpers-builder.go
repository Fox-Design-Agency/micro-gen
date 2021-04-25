package builder

import (
	"io/ioutil"
	"log"
	"os"

	helperConstructor "micro-gen/pkg/constructors/helper-constructor"
)

// initializeHelpers will create the helpers folder and any desired files
// within the folder as required
func initializeHelpers(microType, projectName string, hasDB bool) (err error) {
	if hasDB {
		err = initializeDBHelpers(microType, projectName)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
	//	create the generic response helpers
	err = initializeResponse(microType, projectName)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return nil
}

// initializeDBHelpers will create the db helper struct to
// pass to the service layer
func initializeDBHelpers(microType, projectName string) (err error) {
	// create dbhelpers.go
	b, _ := helperConstructor.ReturnGoDBHelpers()
	var fileEnding string
	var filePath string
	switch microType {
	case "go":
		fileEnding = "dbhelpers.go"
		filePath = "/pkg/helpers/"
	}
	err = ioutil.WriteFile(projectName+filePath+fileEnding, b, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// initializeResponse will create generic response funcs to
// be utilized in the controller files
func initializeResponse(microType, projectName string) (err error) {
	// create response.go
	b, _ := helperConstructor.ReturnGoResponseHelpers()
	var fileEnding string
	var filePath string
	switch microType {
	case "go":
		fileEnding = "response.go"
		filePath = "/pkg/helpers/"
	}
	err = ioutil.WriteFile(projectName+filePath+fileEnding, b, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
