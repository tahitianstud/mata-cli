package file

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
)

// Read will read content from a file
func Read(file string, target *string) error {

	// if file does not exist, return an error it
	if _, err := os.Stat(file); os.IsNotExist(err) {

		return err

	}

	bytes, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	*target = string(bytes)

	return nil

}


// ReadFrom will read a file from a specified location
func ReadFrom(location string, filename string, target *string) error {

	file := fmt.Sprintf("%s%s%s", strings.TrimRight(location, "/"), string(os.PathSeparator), filename)

	return Read(file, target)
}
