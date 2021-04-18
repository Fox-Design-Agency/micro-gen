package builder

import (
	"io/ioutil"
	"log"
	"os"

	"micro-gen/pkg/models"
)

// initializeHelpers will create the helpers folder and any desired files
// within the folder as required
func initializeHelpers(projectName string, hasDB bool) (err error) {
	if hasDB {
		err = initializeDBHelpers(projectName)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
	//	create the generic response helpers
	err = initializeResponse(projectName)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return nil
}

// initializeDBHelpers will create the db helper struct to
// pass to the service layer
func initializeDBHelpers(projectName string) (err error) {
	// create dbhelpers.go
	b, _ := models.ReturnDBHelpers()
	err = ioutil.WriteFile(projectName+"/pkg/helpers/dbhelpers.go", b, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// initializeResponse will create generic response funcs to
// be utilized in the controller files
func initializeResponse(projectName string) (err error) {
	// create response.go
	b, _ := models.ReturnResponseHelpers()
	err = ioutil.WriteFile(projectName+"/pkg/helpers/response.go", b, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}