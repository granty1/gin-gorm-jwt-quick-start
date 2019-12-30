package tools

import "os"

//PathIsExist return this file whether exist
func PathIsExist(directory string) bool {
	_, err := os.Stat(directory)
	if err != nil {
		return false
	}
	return true
}
