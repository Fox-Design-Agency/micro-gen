package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	"micro-gen/pkg/models"
)

// initializeSubServiceModel will initialize the model files for
// the desired subService
func initializeSubServiceModel(name string, projectName string) (err error) {
	// get the byte slice
	b, _ := models.ReturnModelLayer(name)
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
