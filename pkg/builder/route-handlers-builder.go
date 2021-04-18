package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	"micro-gen/pkg/models"
)

// intializeSubServiceRouteHandler will intiialize the route-handler file for the
// subService
func intializeSubServiceRouteHandler(srvc *models.SubService, projectName string) (err error) {
	// get the byte slice
	b, _ := models.ReturnRouteHandlerLayer(srvc.SubServiceName, srvc.HasCRUD)
	// set the filename
	fileName := fmt.Sprintf("/pkg/route-handlers/%s", srvc.SubServiceName+".go")
	// write the file
	err = ioutil.WriteFile(projectName+fileName, b, 0755)
	if err != nil {
		// handle the error
		log.Fatal(err)
	}

	return nil
}
