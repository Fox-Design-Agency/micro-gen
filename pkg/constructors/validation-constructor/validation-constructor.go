package validation

import (
	"log"
)

// ReturnValidationLayer will return the byte slice of the generic validation file
func ReturnValidationLayer(microType, name, modelName, projectName string) ([]byte, error) {
	switch microType {
	case "go":
		// top section
		topString, err := returnGoTopValidationFile(projectName)
		if err != nil {
			// handle the err
			log.Println(err)
		}
		// definition section
		definitionString, err := returnGoValidationDefinitions(name, modelName)
		if err != nil {
			// handle the err
			log.Println(err)
		}
		// method section

		return []byte(topString + definitionString), nil
	default:
		return nil, nil
	}

}
