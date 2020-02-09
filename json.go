package goutils

import (
	"encoding/json"
)

func JsonBytesToType(b []byte, v interface{}) error {
	return json.Unmarshal(b, &v)
}

func JsonStringToType(s string, v interface{}) error {
	return JsonBytesToType([]byte(s), &v)
}

func JsonBytesToInterface(b []byte) (map[string]interface{}, error) {
	var v map[string]interface{}
	err := JsonBytesToType(b, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func JsonStringToInterface(s string) (map[string]interface{}, error) {
	return JsonBytesToInterface([]byte(s))
}

func ToJsonBytes(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func ToJsonString(v interface{}) (string, error) {
	b, err := ToJsonBytes(v)
	return string(b), err
}
func ToJsonBytesIndent(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
func ToJsonStringIndent(v interface{}) (string, error) {
	b, err := ToJsonBytesIndent(v)
	return string(b), err
}
