package modelsMod

import (
	"log"
)

// ReturnModelLayer will return the byte slice of the generic model file
func ReturnGoModelLayer(microType string, name, modelName string) ([]byte, error) {
	switch microType {

	case "go":
		// top section
		topString, err := returnTopModelFile(name, modelName)
		if err != nil {
			// handle the err
			log.Println(err)
		}

		return []byte(topString), nil
	default:
		// handle error
		return nil, nil
	}
}
