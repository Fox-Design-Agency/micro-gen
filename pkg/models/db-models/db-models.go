package db

import (
	"log"
)

// ReturnDBLayer will return the byte slice of the generic db file
func ReturnDBLayer(microType string, hasCrud bool, name, modelName, projectName string) ([]byte, error) {
	switch microType {
	case "go":
		// top section
		topString, err := returnGoTopDBFile(projectName)
		if err != nil {
			// handle the err
			log.Println(err)
		}
		// interface section
		interfaceString, err := returnGoDBInterfaceDeclartion(hasCrud, name, modelName)
		if err != nil {
			// handle the err
			log.Println(err)
		}
		// method section
		methodString, err := returnGoDBMethodDeclarations(hasCrud, name, modelName)
		if err != nil {
			// handle the err
			log.Println(err)
		}

		return []byte(topString + interfaceString + methodString), nil
	default:
		// handle error better
		return nil, nil
	}
}
