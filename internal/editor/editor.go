package editor

import (
	"fmt"
	"os/exec"
	"strings"
)

var ValidEditors = []string{"code", "emacs", "gedit", "nano", "vi", "vim"}

func OpenScriptInEditor(path string, editor string) {
	if editor == "" {
		editor = "gedit" // Default editor
	} else {
		editor = strings.ToLower(editor)
		found := false
		for _, validEditor := range ValidEditors {
			if validEditor == editor {
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Invalid editor. Using default editor (gedit).")
			editor = "gedit" // Fall back to default editor
		}
	}

	fmt.Printf("Opening %s in %s\n", path, editor)

	cmd := exec.Command(editor, path)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error opening script in editor:", err)
	}
}
