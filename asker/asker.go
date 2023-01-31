package asker

import (
	"errors"
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
	"golang.org/x/exp/slices"
)

func ShowContextsSurvey(names []string, selected string) string {
	qs := []*survey.Question{
		{
			Name: "context",
			Prompt: &survey.Select{
				Message: "select a context:",
				Options: names,
				Default: selected,
			},
			Validate: func(val interface{}) error {
				b, ok := val.(core.OptionAnswer)
				if !ok {
					return errors.New("error in cast response")
				}
				if slices.Contains(names, b.Value) {
					return errors.New("cannot be different than the options listed")
				}
				return nil
			},
		},
	}

	answers := struct {
		ContextName string `survey:"context"`
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln("Error in ASK survey")
	}
	return answers.ContextName
}
