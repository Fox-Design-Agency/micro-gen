package builder

import (
	"io/ioutil"
	"log"
	middlewareModels "micro-gen/pkg/models/middleware-models"
)

// initializeMiddleware will initialize the middleware folder and
// a README, current middleware generation is not supported
func initializeMiddleware(microType, projectName string) (err error) {
	switch microType {
	case "go":
		// create a blank middleware README file
		b, _ := middlewareModels.ReturnGoMiddleware()
		err = ioutil.WriteFile(projectName+"/pkg/middleware/README", b, 0755)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	default:
		return nil
	}

}
