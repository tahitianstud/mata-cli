package file

import (
	"os"
	"fmt"
	"strings"
)

// DeleteIfExists will delete a file if it exists and return true
// or return false
func DeleteIfExists(filepath string) (ok bool) {

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	} else {

		err := os.RemoveAll(filepath)

		if err != nil {
			return false
		}

		return true
	}

}

// DeleteIfExistsIn will delete a file if it exists and return true
// or return false
func DeleteIfExistsIn(folder string, filename string) (ok bool) {

	file := fmt.Sprintf("%s%s%s", strings.TrimRight(folder, "/"), string(os.PathSeparator), filename)

	return DeleteIfExists(file)

}
