package db

import (
	"fmt"
	sharedModels "micro-gen/pkg/shared/models"
	"strings"
)

// returnGoTopDBFile will return the top section of the db layer files
// including the package and import
func returnGoTopDBFile(pkgName string) (string, error) {
	pkgStr := fmt.Sprintf("package db\n\n")

	helpersStr := fmt.Sprintf("import (\n") +
		fmt.Sprintf("	\"%s/pkg/helpers\"\n", pkgName) +
		fmt.Sprintf("	\"%s/pkg/models\"\n", pkgName) +
		fmt.Sprintf(")\n\n")

	return pkgStr + helpersStr, nil
}

// returnGoDBInterfaceDeclartion will return the comment block to declare
// the db layer interface, the db layer interface, and the public struct
func returnGoDBInterfaceDeclartion(hasCrud bool, name, modelName string) (string, error) {
	// combine the sections to create the comment block
	commentBlock := sharedModels.TopCommentBlock +
		fmt.Sprintf("\n/ Create the interaction models for the %s\n", strings.Title(name)+"DB") +
		sharedModels.BottomCommentBlock

	// create the interface
	interfaceBlock := fmt.Sprintf("\n\n// %s is used to interact and manage %s\n", "I"+strings.Title(name)+"DB", name) +
		fmt.Sprintf("type %s interface {\n", "I"+strings.Title(name)+"DB")
	if hasCrud {
		interfaceBlock += fmt.Sprintf("	Create%s(*models.%s) error\n", strings.Title(name), strings.Title(modelName)) +
			fmt.Sprintf("	Delete%s(id string) error\n", strings.Title(name)) +
			fmt.Sprintf("	GetAll%s() ([]models.%s,error)\n", strings.Title(name), strings.Title(modelName)) +
			fmt.Sprintf("	Get%s(id string) (*models.%s, error)\n", strings.Title(name), strings.Title(modelName)) +
			fmt.Sprintf("	Update%s() error\n", strings.Title(name))
	}
	interfaceBlock += "}\n\n"

	// create the struct
	structBlock := fmt.Sprintf("\n // %s is the struct to define the %s db \n", strings.Title(name)+"StructDB", name) +
		fmt.Sprintf("type %s struct {\n", strings.Title(name)+"StructDB") +
		fmt.Sprintf("	helpers.DBContext\n}\n\n")

	// put entire section together to return
	fullInterfaceSection := commentBlock + interfaceBlock + structBlock

	return fullInterfaceSection, nil
}

// returnGoDBMethodDeclarations will return the comment block describing how to
// add new methods off of the struct declared in the returnInterfaceDeclartion
// func, as well as return crud func outlines if desired.
func returnGoDBMethodDeclarations(hasCrud bool, name, modelName string) (string, error) {
	// combine the sections to create the comment block
	commentBlock := sharedModels.TopCommentBlock +
		fmt.Sprintf("\n/	Add new methods onto %s that must be defined on the %s\n", strings.Title(name)+"StructDB", strings.Title(name)+"DB") +
		fmt.Sprintf("/	interface. These methods can be utilized anywhere the %s\n", name) +
		fmt.Sprintf("/ 	subservice is utilized.\n") +
		sharedModels.BottomCommentBlock

		// register the method block with an empty line to create intial seperation
	methodBlock := "\n\n"

	if hasCrud {
		// Create func
		methodBlock += fmt.Sprintf("// Create%s will create a %s in the DB\n", strings.Title(name), name) +
			fmt.Sprintf("func (ctx *%s) Create%s(payload *models.%s) error {\n", strings.Title(name)+"StructDB", strings.Title(name), strings.Title(modelName)) +
			fmt.Sprintf("	// start transaction if desired\n\n") +
			fmt.Sprintf("	// declare SQL\n	SQL := `ADD SQL Statement here`\n\n") +
			fmt.Sprintf("	// declare args to pass to create\n	args := []interface{}{\n	//add args\n	}\n\n	// Make the appropiate SQL Call\n") +
			fmt.Sprintf("	if _, err := ctx.DB.Exec(SQL, args...); err != nil {\n") +
			fmt.Sprintf("		// handle err\n		return err\n	}\n\n") +
			fmt.Sprintf("	return nil\n}\n\n")

		// Delete func
		methodBlock += fmt.Sprintf("// Delete%s will delete a %s in the DB\n", strings.Title(name), name) +
			fmt.Sprintf("func (ctx *%s) Delete%s(id string) error {\n", strings.Title(name)+"StructDB", strings.Title(name)) +
			fmt.Sprintf("	// start transaction if desired\n\n") +
			fmt.Sprintf("	// declare SQL\n	SQL := `ADD SQL Statement here`\n\n	// Make the appropiate SQL Call\n") +
			fmt.Sprintf("	if _, err := ctx.DB.Exec(SQL, id); err != nil {\n") +
			fmt.Sprintf("		// handle err\n		return err\n	}\n\n") +
			fmt.Sprintf("	return nil\n}\n\n")

		// GetAll func
		methodBlock += fmt.Sprintf("// GetAll%s will get all %s in the DB\n", strings.Title(name), name) +
			fmt.Sprintf("func (ctx *%s) GetAll%s() ([]models.%s, error) {\n", strings.Title(name)+"StructDB", strings.Title(name), strings.Title(modelName)) +
			fmt.Sprintf("	// initialize a results var to return\n") +
			fmt.Sprintf("	var results []models.%s\n\n", strings.Title(modelName)) +
			fmt.Sprintf("	// start transaction if desired\n\n") +
			fmt.Sprintf("	// declare SQL\n	SQL := `ADD SQL Statement here`\n\n	// Make the appropiate SQL Call\n") +
			fmt.Sprintf("	if err := ctx.DB.Select(&results, SQL); err != nil {\n") +
			fmt.Sprintf("		// handle err\n		return nil, err\n	}\n\n") +
			fmt.Sprintf("	return results, nil\n}\n\n")

		// Get func
		methodBlock += fmt.Sprintf("// Get%s will get all %s in the DB\n", strings.Title(name), name) +
			fmt.Sprintf("func (ctx *%s) Get%s(id string) (*models.%s, error) {\n", strings.Title(name)+"StructDB", strings.Title(name), strings.Title(modelName)) +
			fmt.Sprintf("	// initialize a results var to return\n") +
			fmt.Sprintf("	var results *models.%s\n\n", strings.Title(modelName)) +
			fmt.Sprintf("	// start transaction if desired\n\n") +
			fmt.Sprintf("	// declare SQL\n	SQL := `ADD SQL Statement here`\n\n	// Make the appropiate SQL Call\n") +
			fmt.Sprintf("	if err := ctx.DB.Get(&results, SQL, id); err != nil {\n") +
			fmt.Sprintf("		// handle err\n		return nil, err\n	}\n\n") +
			fmt.Sprintf("	return results, nil\n}\n\n")

		// Update func
		methodBlock += fmt.Sprintf("// Update%s will update %s in the DB\n", strings.Title(name), name) +
			fmt.Sprintf("func (ctx *%s) Update%s() error {\n", strings.Title(name)+"StructDB", strings.Title(name)) +
			fmt.Sprintf("	// start transaction if desired\n\n") +
			fmt.Sprintf("	// declare SQL\n	SQL := `ADD SQL Statement here`\n\n	// Make the appropiate SQL Call\n") +
			fmt.Sprintf("	if _, err := ctx.DB.Exec(SQL); err != nil {\n") +
			fmt.Sprintf("		// handle err\n		return err\n	}\n\n") +
			fmt.Sprintf("	return nil\n}\n")
	}

	fullMethodSection := commentBlock + methodBlock

	return fullMethodSection, nil
}
