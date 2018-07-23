package file

import (
	"os"
	"path/filepath"
	"github.com/tahitianstud/mata-cli/internal/platform/errors"
	"fmt"
	"strings"
	"io/ioutil"
)

// Write will write out a file
func Write(file string, payload []byte, overwrite ...bool) error {

	// if file does not exist, create it
	if _, err := os.Stat(file); os.IsNotExist(err) {

		parentDir := filepath.Dir(file)

		err = os.MkdirAll(parentDir, os.ModePerm)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(file, payload, 0700)

		if err != nil {
			return err
		}


		return nil

	} else if ! os.IsPermission(err) {  // file is accessible

		if len(overwrite) > 0 && overwrite[0] == true {

			err = os.Remove(file)
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(file, payload, 0700)
			if err != nil {
				return err
			}

			return nil
		} else {

			return errors.New("cannot write: file already exist and overwrite is false")

		}

	} else {

		return errors.New("cannot write: file location is not writeable")

	}

}


// Write will write out a file inside the specified location
func WriteInside(location string, filename string, payload []byte, overwrite ...bool) error {

	file := fmt.Sprintf("%s%s%s", strings.TrimRight(location, "/"), string(os.PathSeparator), filename)

	return Write(file, payload, overwrite...)
}
