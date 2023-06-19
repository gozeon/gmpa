package utils

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func PromptBool(msg string) (bool, error) {
	prompt := promptui.Select{
		Label:    fmt.Sprintf("%s,Select[Yes/No]", msg),
		Items:    []string{"Yes", "No"},
		HideHelp: true,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return false, err
	}
	return result == "Yes", nil
}
