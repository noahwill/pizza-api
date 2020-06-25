package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request : facilitates an http request for the client
func Request(method string, path string, payload interface{}) ([]byte, error) {
	payloadData, _ := json.Marshal(&payload)
	url := "http://zarnnr.herokuapp.com/" + path
	req := &http.Request{}
	if payload != nil {
		req, _ = http.NewRequest(method, url, bytes.NewReader(payloadData))
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}
	req.Header.Add("Content-Type", "application/json")
	res, resError := http.DefaultClient.Do(req)
	if resError != nil {
		return []byte{}, resError
	}
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode >= 300 {
		if body != nil {
			return body, fmt.Errorf("error performing request, response code was %v and body was %v", res.Status, string(body))
		}
		return []byte{}, fmt.Errorf("error performing request, respose code was %v", res.Status)
	}
	return body, nil
}
