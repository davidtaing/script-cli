package promptutil

import (
	"fmt"

	"github.com/davidtaing/scriptcli/internal/editor"
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

func PromptOpenEditor() bool {
	p := promptui.Prompt{
		Label:     "Open in editor",
		IsConfirm: true,
	}

	result, _ := p.Run()

	return result == "y" || result == "Y"
}

func PromptSelectEditor() string {
	index := -1
	var result string
	var err error

	for index < 0 {
		p := promptui.Select{
			Label: "Which text editor would you like to use?",
			Items: editor.ValidEditors,
		}

		index, result, err = p.Run()
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}
