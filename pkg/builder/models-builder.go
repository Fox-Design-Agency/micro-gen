package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	models "micro-gen/pkg/models/models-models"
)

// initializeSubServiceModel will initialize the model files for
// the desired subService
func initializeSubServiceModel(name, modelName string, projectName string) (err error) {
	// get the byte slice
	b, _ := models.ReturnGoModelLayer("go", name, modelName)
	// set the filename
	fileName := fmt.Sprintf("/pkg/models/%s", name+"-model.go")
	// write the file
	err = ioutil.WriteFile(projectName+fileName, b, 0755)
	if err != nil {
		// handle the error
		log.Fatal(err)
	}
	return nil
}
