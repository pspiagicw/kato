package manage

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/pspiagicw/goreland"
)

func promptSelection(label string, options []string) int {
	prompt := &survey.Select{
		Message: label,
		Options: options,
	}

	var selected int
	survey.AskOne(prompt, &selected)

	return selected
}
func promptMultiple(label string, options []string) []int {
	choices := make([]int, 0)

	prompt := &survey.MultiSelect{
		Message: label,
		Options: options,
	}
	survey.AskOne(prompt, &choices)

	if len(choices) == 0 {
		goreland.LogFatal("No option selected")
	}

	return choices
}
