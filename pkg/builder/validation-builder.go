package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	validationConstructors "micro-gen/pkg/constructors/validation-constructor"
)

// initializeValidationLayer will initialize the validation layer
// for the desired subService
func initializeValidationLayer(microType, name, modelName, projectName string, hasCRUD bool) (err error) {
	// get the byte slice
	b, _ := validationConstructors.ReturnValidationLayer(microType, name, modelName, projectName, hasCRUD)
	var fileEnding string
	var filePath string
	switch microType {
	case "go":
		fileEnding = "-validator.go"
		filePath = "/pkg/validation/"
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
