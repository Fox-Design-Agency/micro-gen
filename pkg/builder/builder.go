package builder

import (
	"log"
	models "micro-gen/pkg/shared/models"
	"os"
	"os/exec"
	"strings"
)

// IntializeBuild will start the logic stuctures to
// build the desired micro service pattern
func IntializeBuild(answers *models.Questions, microType string) {
	// run logic train on the questions models
	log.Println("Starting Build")

	// build container folder and root
	err := initializeRootProject(answers, microType)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// build pkg folder and run folder
	err = initializeRootFolders(microType, answers.ProjectName, answers.HasDB)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// create required folders for services
	err = createPKGFolders(microType, answers.ProjectName)
	if err != nil {
		// handle err
		log.Println(err)
		os.Exit(1)
	}
	// @TODO move to some code specific initiazation func
	wd, _ := os.Getwd()
	if microType == "go" {
		// should run go mod init here and save package name

		os.Chdir(answers.ProjectName)
		log.Println("running init")
		cmd := exec.Command("go", "mod", "init", answers.ProjectName)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		os.Chdir(wd)
	}
	// if has helpers is enabled, then generate the generic helpers
	if answers.HasHelpers {
		err = initializeHelpers(microType, answers.ProjectName, answers.HasDB)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	// go ahead and initialize middleware
	err = initializeMiddleware(microType, answers.ProjectName)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// temp serviceArray holder to manage service creation in the main
	tempServiceArray := []string{}
	routeHandlerArray := []string{}

	// iter on services to populate pkg
	for _, v := range answers.SubServices {
		// Create model layer
		err = initializeSubServiceModel(microType, v.SubServiceName, v.ModelName, answers.ProjectName)
		if err != nil {
			// TODO handle error better
		}

		// if has DB
		if v.HasDB {
			// Create DB service layer
			err = initializeDBLayer(v.HasCRUD, v.SubServiceName, microType, v.ModelName, answers.ProjectName)
			if err != nil {
				// TODO handle error better
			}

			// Create Validation Service Layer
			err = initializeValidationLayer(microType, v.SubServiceName, v.ModelName, answers.ProjectName)
			if err != nil {
				// TODO handle error better
			}

		}

		// Create service layer for potential interface chaining
		err = initializeSubServiceService(microType, v.SubServiceName, answers.ProjectName, v.HasDB)
		if err != nil {
			// TODO handle error better
		}
		// slot into services array for service population
		subServiceName := strings.Join(strings.Fields(strings.TrimSpace(v.SubServiceName)), "")
		tempServiceArray = append(tempServiceArray, subServiceName)

		if v.HasRouteHandler {

			// slot into services array for service population
			routeHandlerArray = append(routeHandlerArray, subServiceName)

			// create subService route handler
			err = intializeSubServiceRouteHandler(v, microType, answers.ProjectName, answers.HasHelpers)
			if err != nil {
				// TODO handle error better
			}
		}
	}

	// iter on services array for service population
	err = initializeServicesFile(tempServiceArray, microType, answers.ProjectName, answers.HasDB)
	if err != nil {
		// TODO handle error better
	}

	// create run
	err = initializeRun(microType, answers.ProjectName, tempServiceArray, routeHandlerArray, answers.HasDB)
	if err != nil {
		// TODO handle error better
	}
	if microType == "go" {
		// run go mod tidy here
		os.Chdir(answers.ProjectName)
		log.Println("running tidy")
		cmd := exec.Command("go", "mod", "tidy")
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
	// end the thing
	log.Println("Ending Build")
}

func createPKGFolders(microType, projectName string) (err error) {
	switch microType {
	case "go":
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
	default:
		// handle error on bad microType
		return nil
	}
}
