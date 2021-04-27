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
func returnGoValidationDefinitions(name, modelName string, hasCRUD bool) (string, error) {
	validationDefinitions := fmt.Sprintf("\n\n // %s is the validation/normalization struct for %s\n", strings.Title(name)+"Validator", name) +
		fmt.Sprintf("type %s struct {\n", strings.Title(name)+"Validator") +
		fmt.Sprintf("	db.I%s\n}\n\n", strings.Title(name)+"DB") +
		fmt.Sprintf("type %s func(*models.%s) error\n\n", strings.ToLower(name)+"ValFunc", strings.Title(modelName))

	validationRunnerFunc := fmt.Sprintf("func run%sValFuncs(%s *models.%s, fns ...%s) error {\n",
		strings.Title(name), strings.ToLower(name), strings.Title(modelName), strings.ToLower(name)+"ValFunc") +
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

	// if hasCRUD, add a validation placeholder for create
	if hasCRUD {
		validationRunnerMethods += fmt.Sprintf("\n\n// Create%s will validate and normalize the %s fields\n", strings.Title(name), strings.Title(modelName)) +
			fmt.Sprintf("func (v *%sValidator) Create%s(payload *models.%s) error {\n", strings.Title(name), strings.Title(name), strings.Title(modelName)) +
			fmt.Sprintf("	err := run%sValFuncs(payload);// add in validation/normalization methods\n", strings.Title(name)) +
			fmt.Sprintf("	if err != nil {\n") +
			fmt.Sprintf("		return err\n") +
			fmt.Sprintf("	}\n") +
			fmt.Sprintf("	return v.I%sDB.Create%s(payload)\n}\n\n", strings.Title(name), strings.Title(name))
	} else {
		validationRunnerMethods += fmt.Sprintf("\n\n/* this is an example method runner\n") +
			fmt.Sprintf("func (v *ExampleValidator) Example(payload *models.Example) error {\n") +
			fmt.Sprintf("	if err := runExampleValFuncs(payload); // add in validation/normalization methods") +
			fmt.Sprintf("		); err != nil {\n") +
			fmt.Sprintf("		return err\n") +
			fmt.Sprintf("	}\n") +
			fmt.Sprintf("	return v.ExampleDB.Example(payload)\n}*/\n\n")
	}

	validationsMethods := fmt.Sprintf("\n\n") +
		models.TopCommentBlock +
		fmt.Sprintf("\n/	Add New validation/normalization methods below\n") +
		models.BottomCommentBlock +
		fmt.Sprintf("\n\n/* this is an example method\n") +
		fmt.Sprintf("func (v *ExampleValidator) exampleIDMinLength(example *models.Example) error {\n") +
		fmt.Sprintf("	if len(example.ID) > 256 {\n") +
		fmt.Sprintf("		return models.ErrIDTooLong\n	}\n") +
		fmt.Sprintf("	return nil\n}*/\n\n")

	// add a commented method example

	fullValidation := validationDefinitions +
		validationRunnerFunc +
		validationRunnerMethods +
		validationsMethods

	return fullValidation, nil
}
