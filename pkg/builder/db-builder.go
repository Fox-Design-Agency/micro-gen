package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	dbModels "micro-gen/pkg/models/db-models"
)

// initializeDBLayer will initialize the db layer for the subService
func initializeDBLayer(hasCRUD bool, name, modelName, projectName string) (err error) {
	// get the byte slice
	b, _ := dbModels.ReturnDBLayer("go", hasCRUD, name, modelName, projectName)
	// set the filename
	fileName := fmt.Sprintf("/pkg/db/%s", name+"-db.go")
	// write the file
	err = ioutil.WriteFile(projectName+fileName, b, 0755)
	if err != nil {
		// handle the error
		log.Fatal(err)
	}
	return nil
}
