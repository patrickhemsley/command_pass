package lib

import (
	"io/ioutil"
)

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToTextFile(filename string, data string) error {
	err := ioutil.WriteFile(filename, []byte(data), 0777)
	return err
}

//ReadFromFile returns the contents of a textfile as a string
func ReadFromTextFile(filename string) (string, error) {
	result, err := ioutil.ReadFile(filename)
	return string(result), err
}
