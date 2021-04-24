package models

import (
	"fmt"
	"log"
	"strings"
)

// ReturnValidationLayer will return the byte slice of the generic validation file
func ReturnValidationLayer(name, projectName string) ([]byte, error) {
	// top section
	topString, err := returnTopValidationFile(projectName)
	if err != nil {
		// handle the err
		log.Println(err)
	}
	// definition section
	definitionString, err := returnValidationDefinitions(name)
	if err != nil {
		// handle the err
		log.Println(err)
	}
	// method section

	return []byte(topString + definitionString), nil
}

// returnTopValidationFile will return the string of the generic top
// section of the validation layer file
func returnTopValidationFile(projectName string) (string, error) {
	// @TODO will need to add imports????
	topDeclaration := fmt.Sprintf("package validation\n\n") +
		fmt.Sprintf("import (\n") +
		fmt.Sprintf("	\"%s/pkg/db\"\n", projectName) +
		fmt.Sprintf("	\"%s/pkg/models\"\n)\n\n", projectName)
	// combine the sections to create the comment block
	commentBlock := topCommentBlock +
		fmt.Sprintf("\n/ Only Change this section if you are adding new features onto this.\n") +
		fmt.Sprintf("/ Otherwise, add all new stuff in the section below.\n") +
		bottomCommentBlock

		// return the string, coudl add into a new var???
	return topDeclaration + commentBlock, nil
}

// returnValidationDefinitions will return the string of the definitions
// for the validation layer of the service
func returnValidationDefinitions(name string) (string, error) {
	validationDefinitions := fmt.Sprintf("\n\n // %s is the validation/normalization struct for %s\n", strings.Title(name)+"Validator", name) +
		fmt.Sprintf("type %s struct {\n", strings.Title(name)+"Validator") +
		fmt.Sprintf("	db.I%s\n}\n\n", strings.Title(name)+"DB") +
		fmt.Sprintf("type %s func(*models.%s) error\n\n", strings.ToLower(name)+"ValFunc", strings.Title(name))

	validationRunnerFunc := fmt.Sprintf("func run%sValFuncs(%s *models.%s, fns ...%s) error {\n",
		strings.Title(name), name, strings.Title(name), strings.ToLower(name)+"ValFunc") +
		fmt.Sprintf("	for _, fn := range fns {\n") +
		fmt.Sprintf("		if err := fn(%s); err != nil {\n", strings.ToLower(name)) +
		fmt.Sprintf("			return err\n }\n }\n") +
		fmt.Sprintf("		return nil\n}\n\n") +
		//type safety checker
		fmt.Sprintf("// for safety, comfort, and security\n") +
		fmt.Sprintf("var _ db.I%s = &%s {}\n\n", strings.Title(name)+"DB", strings.Title(name)+"Validator")

	validationRunnerMethods := topCommentBlock +
		fmt.Sprintf("\n/	Add New validation runner methods below \n") +
		bottomCommentBlock

	validationsMethods := fmt.Sprintf("\n\n") +
		topCommentBlock +
		fmt.Sprintf("\n/	Add New validation/normalization methods below\n") +
		bottomCommentBlock

	fullValidation := validationDefinitions +
		validationRunnerFunc +
		validationRunnerMethods +
		validationsMethods

	return fullValidation, nil
}
