package validation

import (
	"fmt"
	models "micro-gen/pkg/shared/models"
	"strings"
)

// returnGoTopValidationFile will return the string of the generic top
// section of the validation layer file
func returnGoTopValidationFile(projectName string) (string, error) {
	// @TODO will need to add imports????
	topDeclaration := fmt.Sprintf("package validation\n\n") +
		fmt.Sprintf("import (\n") +
		fmt.Sprintf("	\"%s/pkg/db\"\n", projectName) +
		fmt.Sprintf("	\"%s/pkg/models\"\n)\n\n", projectName)
	// combine the sections to create the comment block
	commentBlock := models.TopCommentBlock +
		fmt.Sprintf("\n/ Only Change this section if you are adding new features onto this.\n") +
		fmt.Sprintf("/ Otherwise, add all new stuff in the section below.\n") +
		models.BottomCommentBlock

		// return the string, coudl add into a new var???
	return topDeclaration + commentBlock, nil
}

// returnGoValidationDefinitions will return the string of the definitions
// for the validation layer of the service
func returnGoValidationDefinitions(name string) (string, error) {
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

	validationRunnerMethods := models.TopCommentBlock +
		fmt.Sprintf("\n/	Add New validation runner methods below \n") +
		models.BottomCommentBlock

	validationsMethods := fmt.Sprintf("\n\n") +
		models.TopCommentBlock +
		fmt.Sprintf("\n/	Add New validation/normalization methods below\n") +
		models.BottomCommentBlock

	fullValidation := validationDefinitions +
		validationRunnerFunc +
		validationRunnerMethods +
		validationsMethods

	return fullValidation, nil
}
