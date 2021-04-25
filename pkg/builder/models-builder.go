package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	modelConstructor "micro-gen/pkg/constructors/models-constructor"
)

// initializeSubServiceModel will initialize the model files for
// the desired subService
func initializeSubServiceModel(microType, name, modelName string, projectName string) (err error) {
	// get the byte slice
	b, _ := modelConstructor.ReturnGoModelLayer(microType, name, modelName)
	var fileEnding string
	var filePath string
	switch microType {
	case "go":
		fileEnding = "-model.go"
		filePath = "/pkg/models/"
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
