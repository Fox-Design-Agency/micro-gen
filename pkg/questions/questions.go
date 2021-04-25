package questions

import (
	"log"
	"os"

	models "micro-gen/pkg/shared/models"
)

// IntializeQuestions initializes the questions to
// determine the appropiate service pattern to generate
func IntializeQuestions(answers *models.Questions) {
	log.Println("Welcome to the Micro Service Generator")
	log.Println("This generator will ask a series of questions and generate the template structure files")

	// set project name
	stringResult, err := runStringPrompt("What is the name of your project?", "Project Name")
	if err != nil {
		log.Println("Prompt failed")
		os.Exit(1)
	}
	answers.ProjectName = stringResult

	// check if should initialize with generic helpers
	boolResult, err := runBoolPrompt("Initialize generic helpers?", "Create Helpers", []string{"Yes", "No"})
	if err != nil {
		log.Println("Prompt failed")
		os.Exit(1)
	}
	answers.HasHelpers = boolResult

	// check if is CLI, commented out until further support for CLI is there
	// boolResult, err = runBoolPrompt("Is this a CLI project?", "Is CLI", []string{"Yes", "No"})
	// if err != nil {
	// 	log.Println("Prompt failed")
	// 	os.Exit(1)
	// }
	// answers.IsCLI = boolResult

	// check if has DB
	boolResult, err = runBoolPrompt("Does this have a DB?", "Has DB", []string{"Yes", "No"})
	if err != nil {
		log.Println("Prompt failed")
		os.Exit(1)
	}
	answers.HasDB = boolResult

	// query number of services
	int64Result, err := runNumberPrompt("How many SubServices are in the project?", "SubService Count")
	if err != nil {
		log.Println("Prompt failed")
		os.Exit(1)
	}
	// initialize subservice
	var subServices []*models.SubService

	// iterate on services to gen pattern
	for i := int64Result; i > 0; i-- {
		subService := models.SubService{}

		// run sub service questions
		IntializeSubServiceQuestions(&subService, answers.HasDB)

		// set subservice into subservice slice
		subServices = append(subServices, &subService)
	}
	answers.SubServices = subServices
}
