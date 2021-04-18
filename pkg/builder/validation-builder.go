package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	"micro-gen/pkg/models"
)

// initializeValidationLayer will initialize the validation layer
// for the desired subService
func initializeValidationLayer(name, projectName string) (err error) {
	// get the byte slice
	b, _ := models.ReturnValidationLayer(name)
	// set the filename
	fileName := fmt.Sprintf("/pkg/validation/%s", name+"-validator.go")
	// write the file
	err = ioutil.WriteFile(projectName+fileName, b, 0755)
	if err != nil {
		// handle the error
		log.Fatal(err)
	}
	return nil
}
