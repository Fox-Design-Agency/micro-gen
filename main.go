package main

import (
	"micro-gen/pkg/builder"
	"micro-gen/pkg/models"
	"micro-gen/pkg/questions"
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
	builder.IntializeBuild(&answers)

	os.Exit(1)
}
