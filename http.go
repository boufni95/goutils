package goutils

import (
	"fmt"
	"io"
	"net/http"
)

//ExtractReqBody extracts the body from the request and parse it, max body size 10000bytes
func ExtractReqBody(r *http.Request, i interface{}) error {
	b := make([]byte, 10000)
	n, err := r.Body.Read(b)
	b = b[0:n]
	if err != nil && err != io.EOF {
		return err
	}
	err = JsonBytesToType(b, &i)
	if err != nil {
		return err

	}
	return nil
}

func SendResError(w http.ResponseWriter, reason string) error {
	errorMex := make(map[string]interface{})
	errorMex["status"] = "ERROR"
	errorMex["reason"] = reason
	b, err := ToJsonStringIndent(errorMex)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, b)
	return nil
}

func SendResOk(w http.ResponseWriter) error {
	okMex := make(map[string]interface{})
	okMex["status"] = "OK"
	b, err := ToJsonStringIndent(okMex)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, b)
	return nil
}
