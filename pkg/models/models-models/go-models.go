package modelsMod

import (
	"fmt"
	"micro-gen/pkg/shared/models"
	"strings"
)

func returnTopModelFile(name, modelName string) (string, error) {
	modelString := "package models \n\n"

	// combine the sections to create the comment block
	modelString += models.TopCommentBlock +
		fmt.Sprintf("\n/ define struct types for  %s\n", name) +
		models.BottomCommentBlock +
		fmt.Sprintf("\n\n// %s is the struct type for the %s DB\n", strings.Title(modelName), name) +
		fmt.Sprintf("type %s struct{ }\n", strings.Title(modelName))

	return modelString, nil
}
