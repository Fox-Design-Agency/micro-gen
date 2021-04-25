package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	routeHandlerConstructors "micro-gen/pkg/constructors/route-handlers-constructor"
	models "micro-gen/pkg/shared/models"
)

// intializeSubServiceRouteHandler will intiialize the route-handler file for the
// subService
func intializeSubServiceRouteHandler(srvc *models.SubService, microType, projectName string, hasHelpers bool) (err error) {
	// get the byte slice
	b, _ := routeHandlerConstructors.ReturnRouteHandlerLayer(microType, srvc.SubServiceName, srvc.ModelName, projectName, srvc.HasCRUD, hasHelpers)
	// set the filename
	var fileEnding string
	var filePath string
	switch microType {
	case "go":
		fileEnding = ".go"
		filePath = "/pkg/route-handlers/"
	}
	fileName := fmt.Sprintf("%s", filePath+srvc.SubServiceName+fileEnding)
	// write the file
	err = ioutil.WriteFile(projectName+fileName, b, 0755)
	if err != nil {
		// handle the error
		log.Fatal(err)
	}

	return nil
}
