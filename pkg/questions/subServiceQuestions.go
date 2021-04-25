package questions

import (
	"log"
	models "micro-gen/pkg/shared/models"
	"os"
)

// IntializeSubServiceQuestions will initialize the question set for the desired number
// of sub srvices. These subservices will be saved in their own structs and return to
// be set into a subsrvice array that will be passed to the builder.
func IntializeSubServiceQuestions(subStruct *models.SubService, hasDB bool) error {
	// Set Subservice name
	stringResult, err := runStringPrompt("What is the name of the sub service?", "SubService Name")
	if err != nil {
		log.Println("Prompt failed")
		os.Exit(1)
	}
	subStruct.SubServiceName = stringResult

	// Set the subservice model name
	stringResult, err = runStringPrompt("What should the model be called?", "Model Name")
	if err != nil {
		log.Println("Prompt failed")
		os.Exit(1)
	}
	subStruct.ModelName = stringResult

	// if the microsrvice has a DB, check if the subservice requires a DB layer
	if hasDB {
		// Has DB
		boolResult, err := runBoolPrompt("Does this sub service need a DB layer?", "Has DB", []string{"Yes", "No"})
		if err != nil {
			log.Println("Prompt failed")
			os.Exit(1)
		}
		subStruct.HasDB = boolResult

		// if subservice has db, then should it intialize crud?
		if subStruct.HasDB {
			boolResult, err := runBoolPrompt("Does DB layer need CRUD?", "Needs CRUD", []string{"Yes", "No"})
			if err != nil {
				log.Println("Prompt failed")
				os.Exit(1)
			}
			subStruct.HasCRUD = boolResult
		}
	}

	// check if this service needs a routeHandler
	boolResult, err := runBoolPrompt("Does this need a route handler?", "Needs route handler", []string{"Yes", "No"})
	if err != nil {
		log.Println("Prompt failed")
		os.Exit(1)
	}
	subStruct.HasRouteHandler = boolResult

	return nil
}
