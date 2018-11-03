package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// Store sends a request to the write to db function with the given data.
// The writetodb function will store this data.
// The environment variable "writetodb_uri" must point to the writetodb function.
func Store(data interface{}) ([]byte, error) {
	url := os.Getenv("writetodb_uri")

	if url == "" {
		return nil, errors.New("Mongo Url env is empty")
	}

	body, err := createBody(data)

	if err != nil {
		return nil, err
	}

	return postData(url, body)
}

// Find sends a query request to the find in db function and returns the data.
// The environment variable "findindb_uri" must point to the findindb function.
func Find(query interface{}) ([]byte, error) {
	url := os.Getenv("findindb_uri")

	if url == "" {
		return nil, errors.New("Mongo Url env is empty")
	}

	body, err := createBody(query)

	if err != nil {
		return nil, err
	}

	return postData(url, body)
}

func createBody(data interface{}) (io.Reader, error) {
	json, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(json), nil
}

func postData(url string, data io.Reader) ([]byte, error) {
	response, err := http.Post(url, "application/json", data)

	if err != nil {
		return nil, errors.New("Error by sending post request")
	}

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusAccepted {
		return nil, errors.New("Invalid response status code")
	}

	return ioutil.ReadAll(response.Body)
}
