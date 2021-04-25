package modelsMod

import (
	"fmt"
	"micro-gen/pkg/shared/models"
)

func returnTopModelFile(name string) (string, error) {
	modelString := "package models \n\n"

	// combine the sections to create the comment block
	modelString += models.TopCommentBlock +
		fmt.Sprintf("\n/ define struct types for  %s\n", name) +
		models.BottomCommentBlock

	return modelString, nil
}
