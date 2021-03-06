package questions

import (
	"log"
	"strconv"

	"github.com/manifoldco/promptui"
)

// runStringPrompt prompts for a string value and returns the string result
func runStringPrompt(question string, label string) (string, error) {
	log.Println(question)
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Println(err)
		log.Println("Prompt failed")
		return "", err
	}

	return result, nil
}

// runNumberPrompt prompts for a number value and returns the int64 result
func runNumberPrompt(question string, label string) (int64, error) {
	log.Println(question)
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Println("Prompt failed")
		return 0, err
	}
	res, err := strconv.ParseInt(result, 0, 0)
	if err != nil {
		log.Println("Prompt failed")
		return 0, err
	}

	return res, nil
}

// runBoolPrompt prompts for a yes or no and returns the bool equivilant
// yes -> TRUE
// no -> FALSE
func runBoolPrompt(question string, label string, items []string) (bool, error) {
	log.Println(question)
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Println("Prompt failed")
		return false, err
	}

	if result == "Yes" {
		return true, nil
	}

	return false, nil
}
