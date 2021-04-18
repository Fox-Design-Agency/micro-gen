package models

import (
	"fmt"
	"log"
)

// ReturnModelLayer will return the byte slice of the generic model file
func ReturnModelLayer(name string) ([]byte, error) {
	// top section
	topString, err := returnTopModelFile(name)
	if err != nil {
		// handle the err
		log.Println(err)
	}

	return []byte(topString), nil
}

func returnTopModelFile(name string) (string, error) {
	modelString := "package models \n\n"

	// combine the sections to create the comment block
	modelString += topCommentBlock +
		fmt.Sprintf("\n/ define struct types for  %s\n", name) +
		bottomCommentBlock

	return modelString, nil
}
