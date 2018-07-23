package json

import "github.com/json-iterator/go"

func Marshal(data interface{}) (string, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	bytes, err := json.Marshal(&data)

	if err != nil {
		return "", err
	}

	return string(bytes[:]), nil
}