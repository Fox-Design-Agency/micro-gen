package builder

import (
	"io/ioutil"
	"log"
	"micro-gen/pkg/models"
	"os"
)

// initializeRootProject will create the base folder and the initial files
// that are found in the root of the project
func initializeRootProject(answers *models.Questions) error {
	// build root folder
	err := os.Mkdir(answers.ProjectName, 0755)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// build root files
	err = initializeRootFiles(answers)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return nil
}

// initializeRootFolders will create the pkg and run folders required for
// file creation
func initializeRootFolders(projectName string, hasDB bool) (err error) {
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
		initializeMigrations(projectName)
	}
	// build postman
	err = os.Mkdir(projectName+"/resources/postman", 0755)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return nil
}

// initializeRootFiles will create the generic root files for the project
// these files include:
// .gitignore
// .dockerignore
// README
func initializeRootFiles(answers *models.Questions) (err error) {
	// .gitignore creation
	b, _ := models.ReturnGitignore()
	err = ioutil.WriteFile(answers.ProjectName+"/.gitignore", b, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// .dockerignore creation
	b, _ = models.ReturnDockerignore()
	err = ioutil.WriteFile(answers.ProjectName+"/.dockerignore", b, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// README creation
	b, _ = models.ReturnREADME()
	err = ioutil.WriteFile(answers.ProjectName+"/README", b, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// README creation
	b, _ = models.ReturnDockerfile()
	err = ioutil.WriteFile(answers.ProjectName+"/Dockerfile", b, 0755)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// initializeMigrations will create a migrations folder and a blank README
// placeholder for DB migration files
// creates:
// - 001_seed.sql
// - 002_seed_func.sql
func initializeMigrations(projectName string) (err error) {
	// ensure that the migrations folder exists
	err = os.Mkdir(projectName+"/resources/migrations", 0755)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// create a blank 001 seed file
	b1, _ := models.ReturnMigrationSeedFile()
	err = ioutil.WriteFile(projectName+"/resources/migrations/001_seed.sql", b1, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// create a blank 001 seed func file
	b2, _ := models.ReturnMigrationSeedFile()
	err = ioutil.WriteFile(projectName+"/resources/migrations/002_seed_func.sql", b2, 0755)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
