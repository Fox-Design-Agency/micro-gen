package builder

import (
	"io/ioutil"
	"log"
	rootfileConstructors "micro-gen/pkg/constructors/root-file-constructor"
	models "micro-gen/pkg/shared/models"
	"os"
)

// initializeRootProject will create the base folder and the initial files
// that are found in the root of the project
func initializeRootProject(answers *models.Questions, microType string) error {
	// build root folder
	err := os.Mkdir(answers.ProjectName, 0755)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// build root files
	err = initializeRootFiles(microType, answers)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return nil
}

// initializeRootFolders will create the pkg and run folders required for
// file creation
func initializeRootFolders(microType, projectName string, hasDB bool) (err error) {
	switch microType {
	case "go":
		// build pkg
		err = os.Mkdir(projectName+"/pkg", 0755)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		// build run
		err = os.Mkdir(projectName+"/run", 0755)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		// ensure that the resources folder exists
		err = os.Mkdir(projectName+"/resources", 0755)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		// if has DB, then build db migrations
		if hasDB {
			initializeMigrations(microType, projectName)
		}
		// build postman
		err = os.Mkdir(projectName+"/resources/postman", 0755)
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

// initializeRootFiles will create the generic root files for the project
// these files include:
// .gitignore
// .dockerignore
// README
func initializeRootFiles(microType string, answers *models.Questions) (err error) {
	switch microType {
	case "go":
		// .gitignore creation
		b, _ := rootfileConstructors.ReturnGoGitignore()
		err = ioutil.WriteFile(answers.ProjectName+"/.gitignore", b, 0755)
		if err != nil {
			log.Fatal(err)
		}
		// .dockerignore creation
		b, _ = rootfileConstructors.ReturnGoDockerignore()
		err = ioutil.WriteFile(answers.ProjectName+"/.dockerignore", b, 0755)
		if err != nil {
			log.Fatal(err)
		}
		// README creation
		b, _ = rootfileConstructors.ReturnGoREADME()
		err = ioutil.WriteFile(answers.ProjectName+"/README", b, 0755)
		if err != nil {
			log.Fatal(err)
		}
		// README creation
		b, _ = rootfileConstructors.ReturnGoDockerfile()
		err = ioutil.WriteFile(answers.ProjectName+"/Dockerfile", b, 0755)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	default:
		// handle error on bad microType
		return nil
	}

}

// initializeMigrations will create a migrations folder and a blank README
// placeholder for DB migration files
// creates:
// - 001_seed.sql
// - 002_seed_func.sql
func initializeMigrations(microType, projectName string) (err error) {
	switch microType {
	case "go":
		// ensure that the migrations folder exists
		err = os.Mkdir(projectName+"/resources/migrations", 0755)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		// create a blank 001 seed file
		b1, _ := rootfileConstructors.ReturnGoMigrationSeedFile()
		err = ioutil.WriteFile(projectName+"/resources/migrations/001_seed.sql", b1, 0755)
		if err != nil {
			log.Fatal(err)
		}

		// create a blank 001 seed func file
		err = ioutil.WriteFile(projectName+"/resources/migrations/002_seed_func.sql", b1, 0755)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	default:
		return nil
	}

}
