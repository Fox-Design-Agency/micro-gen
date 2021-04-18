package builder

import (
	"io/ioutil"
	"log"
	"micro-gen/pkg/models"
)

// initializeMiddleware will initialize the middleware folder and
// a README, current middleware generation is not supported
func initializeMiddleware(projectName string) (err error) {
	// create a blank middleware README file
	b, _ := models.ReturnMiddleware()
	err = ioutil.WriteFile(projectName+"/pkg/middleware/README", b, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
