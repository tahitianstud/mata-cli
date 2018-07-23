package survey

import (
	"strings"

	survey "gopkg.in/AlecAivazis/survey.v1"
	model "gopkg.in/jeevatkm/go-model.v1"
)

// AskQuestionsForConfig will use an annotated struct to prompt user for questions
func AskQuestionsForConfig(config interface{}) error {

	var questions, err = createQuestionsFrom(config)
	if err != nil {
		return err
	}

	err = survey.Ask(questions, config)
	if err != nil {
		return err
	}

	return nil
}

// AskForPassword will prompt the user for his password
func AskForPassword(message string) string {
	password := ""

	prompt := &survey.Password{
		Message: message,
	}
	survey.AskOne(prompt, &password, nil)

	return password
}

// AskForSelection will draw a selection survey, prompts the user for a choice and return its value
func AskForSelection(message string, options []string) string {
	choice := ""

	prompt := &survey.Select{
		// Message:  "Choose a stream to search in:",
		Message:  message,
		Options:  options,
		PageSize: 15,
	}
	survey.AskOne(prompt, &choice, nil)

	return choice
}

// createQuestionsFrom will create survey questions from an annotated struct
func createQuestionsFrom(config interface{}) (questions []*survey.Question, err error) {
	// use introspection to ask questions and fill out the necessary config
	configFields, err := model.Fields(config)

	if err != nil {
		return nil, err
	}

	// create questions array that will be shown to the user
	questions = make([]*survey.Question, 0)
	for _, configField := range configFields {
		configName := configField.Name
		configDescription := configField.Tag.Get("description")
		configType := configField.Tag.Get("type")
		defaultValue, err := model.Get(config, configName)

		if err != nil {
			return nil, err
		}

		// create right type of question according to model
		switch configType {
		case "password":
			passwordQuestion := survey.Question{
				Name: configName,
				Prompt: &survey.Password{
					Message: configDescription,
				},
				Validate: survey.Required,
			}

			questions = append(questions, &passwordQuestion)
		case "select":
			configChoice := configField.Tag.Get("choice")
			options := strings.Split(configChoice, "|")

			selectQuestion := survey.Question{
				Name: configName,
				Prompt: &survey.Select{
					Message: configDescription,
					Options: options,
					Default: defaultValue.(string),
				},
				Validate: survey.Required,
			}

			questions = append(questions, &selectQuestion)
		default:

			question := survey.Question{
				Name: configName,
				Prompt: &survey.Input{
					Message: configDescription,
					Default: defaultValue.(string),
				},
				Validate: survey.Required,
			}

			questions = append(questions, &question)
		}
	}

	return questions, nil
}
