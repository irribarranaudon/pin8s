package asker

import (
	"errors"
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
	"golang.org/x/exp/slices"
)

type Answers struct {
	Context string
	Pod     string
	Action  string
}

type Optional func(f *survey.Select)

func WithDefaultValue(selected string) Optional {
	return func(f *survey.Select) {
		f.Default = selected
	}
}

func ShowSurvey(names []string, name string, selected ...Optional) Answers {

	config := &survey.Select{
		Message: "select " + name + ":",
		Options: names,
	}

	for _, optional := range selected {
		optional(config)
	}

	qs := []*survey.Question{
		{
			Name:   name,
			Prompt: config,
			Validate: func(val interface{}) error {
				b, ok := val.(core.OptionAnswer)
				if !ok {
					return errors.New("error in cast response")
				}
				if !slices.Contains(names, b.Value) {
					return errors.New("cannot be different than the options listed")
				}
				return nil
			},
		},
	}

	answers := Answers{}
	Ask(qs, &answers)
	return answers
}

func Ask(qs []*survey.Question, answers *Answers) {
	err := survey.Ask(qs, answers)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln("Error in ASK survey")
	}
}
