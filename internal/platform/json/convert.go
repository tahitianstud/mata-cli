package json

import (
	"gopkg.in/jeevatkm/go-model.v1"
	"strings"
	"encoding/json"
)

func SerializeToString(data interface{}) string {
	bytes, e := json.Marshal(data)

	if e != nil {
		return ""
	}

	return string(bytes[:])
}

func DeSerializeString(ticket string, target interface{}) error {

	err := json.Unmarshal([]byte(ticket), target)

	if err != nil {
		return err
	}

	return nil
}

// ToStringWithContext will marshall data structure into Json string
// while taking into account tags for :
// - sensitive information
// - other metadata
func ToStringWithContext(data interface{}) (string, error) {

	fields, err := model.Fields(data)
	if err != nil {
	   return "", err
	}

	marshalledString, err := Marshal(data)
	if err != nil {
		return "", err
	}

	// traverse metadata to blank out sensitive information
	for _, field := range fields {
		fieldContext := field.Tag.Get("context")
		if fieldContext == "sensitive" {
			fieldValue, err := model.Get(data, field.Name)
			if err != nil {
				return "", err
			}

			marshalledString = strings.Replace(marshalledString, fieldValue.(string), "**********", -1)
		}
	}

	return marshalledString, nil
}