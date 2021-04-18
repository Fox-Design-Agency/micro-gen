package models

import (
	"fmt"
	"log"
	"strings"
)

// ReturnDBLayer will return the byte slice of the generic db file
func ReturnDBLayer(hasCrud bool, name string) ([]byte, error) {
	// top section
	topString, err := returnTopDBFile()
	if err != nil {
		// handle the err
		log.Println(err)
	}
	// interface section
	interfaceString, err := returnDBInterfaceDeclartion(hasCrud, name)
	if err != nil {
		// handle the err
		log.Println(err)
	}
	// method section
	methodString, err := returnDBMethodDeclarations(hasCrud, name)
	if err != nil {
		// handle the err
		log.Println(err)
	}

	return []byte(topString + interfaceString + methodString), nil
}

// returnTopDBFile will return the top section of the db layer files
// including the package and import
func returnTopDBFile() (string, error) {
	return `package db

	// import helpers
	
`, nil
}

// returnDBInterfaceDeclartion will return the comment block to declare
// the db layer interface, the db layer interface, and the public struct
func returnDBInterfaceDeclartion(hasCrud bool, name string) (string, error) {
	// combine the sections to create the comment block
	commentBlock := topCommentBlock +
		fmt.Sprintf("\n/ Create the interaction models for the %s\n", strings.Title(name)+"DB") +
		bottomCommentBlock

	// create the interface
	interfaceBlock := fmt.Sprintf("\n // %s is used to interact and manage %s\n", "I"+strings.Title(name)+"DB", name) +
		fmt.Sprintf(" type %s interface {\n", "I"+strings.Title(name)+"DB")
	if hasCrud {
		interfaceBlock += fmt.Sprintf("Create%s() error\n", name) +
			fmt.Sprintf("Delete%s() error\n", name) +
			fmt.Sprintf("GetAll%s() error\n", name) +
			fmt.Sprintf("Get%s() error\n", name) +
			fmt.Sprintf("Update%s() error\n", name)
	}
	interfaceBlock += "}"

	// create the struct
	structBlock := fmt.Sprintf("\n // %s is the struct to define the %s db \n", strings.Title(name)+"StructDB", name) +
		fmt.Sprintf(" type %s struct {\n", strings.Title(name)+"StructDB") +
		fmt.Sprintf("helpers.DBContext\n}\n\n")

	// put entire section together to return
	fullInterfaceSection := commentBlock + interfaceBlock + structBlock

	return fullInterfaceSection, nil
}

// returnDBMethodDeclarations will return the comment block describing how to
// add new methods off of the struct declared in the returnInterfaceDeclartion
// func, as well as return crud func outlines if desired.
func returnDBMethodDeclarations(hasCrud bool, name string) (string, error) {
	// combine the sections to create the comment block
	commentBlock := topCommentBlock +
		fmt.Sprintf("\n/	Add new methods onto %s that must be defined on the %s\n", strings.Title(name)+"StructDB", strings.Title(name)+"DB") +
		fmt.Sprintf("/	interface. These methods can be utilized anywhere the %s\n", name) +
		fmt.Sprintf("/ 	subservice is utilized.\n") +
		bottomCommentBlock

		// register the method block with an empty line to create intial seperation
	methodBlock := "\n\n"

	if hasCrud {
		// Create func
		methodBlock += fmt.Sprintf("// Create%s will create a %s in the DB\n", strings.Title(name), name) +
			fmt.Sprintf("func (ctx *%s) Create%s() error {\n", strings.Title(name)+"StructDB", strings.Title(name)) +
			fmt.Sprintf("// initialize a results var to return\n\n // start transaction if desired\n\n") +
			fmt.Sprintf("// declare SQL\n SQL := `ADD SQL Statement here`\n\n") +
			fmt.Sprintf("// declare args to pass to create\nargs := []interface{}{\n//add args\n}\n\n // Make the appropiate SQL Call\n") +
			fmt.Sprintf("if err := ctx.DB.QueryRow(SQL, args...); err != nil {\n") +
			fmt.Sprintf("// handle err\n return err\n }\n\n") +
			fmt.Sprintf("return nil\n }\n\n")

		// Delete func
		methodBlock += fmt.Sprintf("// Delete%s will delete a %s in the DB\n", strings.Title(name), name) +
			fmt.Sprintf("func (ctx *%s) Delete%s() error {\n", strings.Title(name)+"StructDB", strings.Title(name)) +
			fmt.Sprintf("// start transaction if desired\n\n") +
			fmt.Sprintf("// declare SQL\n SQL := `ADD SQL Statement here`\n\n // Make the appropiate SQL Call\n") +
			fmt.Sprintf("if _, err := ctx.DB.Exec(SQL); err != nil {\n") +
			fmt.Sprintf("// handle err\n return err\n }\n\n") +
			fmt.Sprintf("return nil\n }\n\n")

		// GetAll func
		methodBlock += fmt.Sprintf("// GetAll%s will get all %s in the DB\n", strings.Title(name), name) +
			fmt.Sprintf("func (ctx *%s) GetAll%s() error {\n", strings.Title(name)+"StructDB", strings.Title(name)) +
			fmt.Sprintf("// initialize a results var to return\n\n // start transaction if desired\n\n") +
			fmt.Sprintf("// declare SQL\n SQL := `ADD SQL Statement here`\n\n // Make the appropiate SQL Call\n") +
			fmt.Sprintf("if err := ctx.DB.Select(&results, SQL); err != nil {\n") +
			fmt.Sprintf("// handle err\n return err\n }\n\n") +
			fmt.Sprintf("return nil\n }\n\n")

		// Get func
		methodBlock += fmt.Sprintf("// Get%s will get all %s in the DB\n", strings.Title(name), name) +
			fmt.Sprintf("func (ctx *%s) Get%s() error {\n", strings.Title(name)+"StructDB", strings.Title(name)) +
			fmt.Sprintf("// initialize a results var to return\n\n // start transaction if desired\n\n") +
			fmt.Sprintf("// declare SQL\n SQL := `ADD SQL Statement here`\n\n // Make the appropiate SQL Call\n") +
			fmt.Sprintf("if err := ctx.DB.Get(&results, SQL); err != nil {\n") +
			fmt.Sprintf("// handle err\n return err\n }\n\n") +
			fmt.Sprintf("return nil\n }\n\n")

		// Update func
		methodBlock += fmt.Sprintf("// Update%s will update %s in the DB\n", strings.Title(name), name) +
			fmt.Sprintf("func (ctx *%s) Update%s() error {\n", strings.Title(name)+"StructDB", strings.Title(name)) +
			fmt.Sprintf("// start transaction if desired\n\n") +
			fmt.Sprintf("// declare SQL\n SQL := `ADD SQL Statement here`\n\n // Make the appropiate SQL Call\n") +
			fmt.Sprintf("if _, err := ctx.DB.Exec(SQL); err != nil {\n") +
			fmt.Sprintf("// handle err\n return err\n }\n\n") +
			fmt.Sprintf("return nil\n }\n")
	}

	fullMethodSection := commentBlock + methodBlock

	return fullMethodSection, nil
}
