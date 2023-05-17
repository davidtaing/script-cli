package promptutil

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func PromptSelectItems(items []string, label string) (string, error) {
	p := promptui.Select{
		Label: label,
		Items: items,
	}

	i, _, err := p.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return items[i], nil
}
