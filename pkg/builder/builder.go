package builder

import (
	"log"
	"micro-gen/pkg/models"
	"os"
	"strings"
)

// IntializeBuild will start the logic stuctures to
// build the desired micro service pattern
func IntializeBuild(answers *models.Questions) {
	// run logic train on the questions models
	log.Println("Starting Build")

	// build container folder and root
	err := initializeRootProject(answers)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// build pkg folder and run folder
	err = initializeRootFolders(answers.ProjectName, answers.HasDB)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// create required folders for services
	err = createPKGFolders(answers.ProjectName)
	if err != nil {
		// handle err
		log.Println(err)
		os.Exit(1)
	}

	// should run go mod init here and save package name

	// if has helpers is enabled, then generate the generic helpers
	if answers.HasHelpers {
		err = initializeHelpers(answers.ProjectName, answers.HasDB)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	// go ahead and initialize middleware
	err = initializeMiddleware(answers.ProjectName)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// temp serviceArray holder to manage service creation in the main
	tempServiceArray := []string{}

	// iter on services to populate pkg
	for _, v := range answers.SubServices {
		// Create model layer
		err = initializeSubServiceModel(v.SubServiceName, answers.ProjectName)
		if err != nil {
			// TODO handle error better
		}

		// if has DB
		if v.HasDB {
			// Create DB service layer
			err = initializeDBLayer(v.HasCRUD, v.SubServiceName, answers.ProjectName)
			if err != nil {
				// TODO handle error better
			}

			// Create Validation Service Layer
			err = initializeValidationLayer(v.SubServiceName, answers.ProjectName)
			if err != nil {
				// TODO handle error better
			}

			// Create service layer for interface chaining
			err = initializeSubServiceService(v.SubServiceName, answers.ProjectName)
			if err != nil {
				// TODO handle error better
			}

			// slot into services array for service population
			subServiceName := strings.Join(strings.Fields(strings.TrimSpace(v.SubServiceName)), "")
			tempServiceArray = append(tempServiceArray, subServiceName)

		}

		// create subService route handler
		err = intializeSubServiceRouteHandler(v, answers.ProjectName)
		if err != nil {
			// TODO handle error better
		}

	}

	// @TODO REMOVE
	// just to check the service array
	for _, v := range tempServiceArray {
		log.Println(v)
	}

	// iter on services array for service population
	err = initializeServicesFile(tempServiceArray, answers.ProjectName, answers.HasDB)
	if err != nil {
		// TODO handle error better
	}

	// create run
	err = initializeRun(answers.ProjectName, tempServiceArray, answers.HasDB)
	if err != nil {
		// TODO handle error better
	}

	// run go mod tidy here

	log.Println("Ending Build")
}

func createPKGFolders(projectName string) (err error) {
	// create the required folders
	// route-handlers
	err = os.Mkdir(projectName+"/pkg/route-handlers", 0755)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// db
	err = os.Mkdir(projectName+"/pkg/db", 0755)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// helpers
	err = os.Mkdir(projectName+"/pkg/helpers", 0755)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//middleware
	err = os.Mkdir(projectName+"/pkg/middleware", 0755)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//models
	err = os.Mkdir(projectName+"/pkg/models", 0755)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	//sub-services
	err = os.Mkdir(projectName+"/pkg/sub-services", 0755)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//validation
	err = os.Mkdir(projectName+"/pkg/validation", 0755)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return nil
}
