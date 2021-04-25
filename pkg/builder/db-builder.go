package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	dbConstructor "micro-gen/pkg/constructors/db-constructor"
)

// initializeDBLayer will initialize the db layer for the subService
func initializeDBLayer(hasCRUD bool, name, microType, modelName, projectName string) (err error) {
	// get the byte slice
	b, _ := dbConstructor.ReturnDBLayer(microType, hasCRUD, name, modelName, projectName)
	var fileEnding string
	var filePath string
	switch microType {
	case "go":
		fileEnding = "-db.go"
		filePath = "/pkg/db/"
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
