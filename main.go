package main

import (
	"micro-gen/pkg/builder"
	"micro-gen/pkg/questions"
	models "micro-gen/pkg/shared/models"
	"os"
)

func init() {
	// init flags and set to vars
}

func main() {
	// check for help flag to print instructions
	// TODO

	// intialize struct pointer to pass to the questions
	answers := models.Questions{}

	// run question path passing in pointer
	questions.IntializeQuestions(&answers)

	// utlize answers to pass to the builder
	// @TODO capture which code to build microservice in
	// for now just default to go
	builder.IntializeBuild(&answers, "go")

	os.Exit(1)
}
