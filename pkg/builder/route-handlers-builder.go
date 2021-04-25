package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	routeHandlerModels "micro-gen/pkg/models/route-handlers-models"
	models "micro-gen/pkg/shared/models"
)

// intializeSubServiceRouteHandler will intiialize the route-handler file for the
// subService
func intializeSubServiceRouteHandler(srvc *models.SubService, projectName string, hasHelpers bool) (err error) {
	// get the byte slice
	b, _ := routeHandlerModels.ReturnRouteHandlerLayer("go", srvc.SubServiceName, srvc.ModelName, projectName, srvc.HasCRUD, hasHelpers)
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
