package goutils

import (
	"io/ioutil"
	"os"
)

func ReadFileBytes(path string) ([]byte, error) {
	var b []byte
	file, err := os.Open(path)
	if err != nil {
		return b, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)

}
func ReadFileString(path string) (string, error) {
	b, err := ReadFileBytes(path)
	return string(b), err
}

func JsonFileToType(path string, v interface{}) error {
	b, err := ReadFileBytes(path)
	if err != nil {
		return err
	}
	return JsonBytesToType(b, &v)
}
func JsonFileToInterface(path string) (map[string]interface{}, error) {
	b, err := ReadFileBytes(path)
	if err != nil {
		return nil, err
	}
	return JsonBytesToInterface(b)
}
