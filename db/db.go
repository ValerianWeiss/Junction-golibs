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
	return exec(url, data)
}

// Find sends a query request to the find in db function and returns the data.
// The environment variable "findindb_uri" must point to the findindb function.
func Find(query interface{}) ([]byte, error) {
	url := os.Getenv("findindb_uri")
	return exec(url, query)
}

func exec(url string, data interface{}) ([]byte, error) {
	if url == "" {
		return nil, errors.New("Mongo Url env is empty")
	}

	body, err := createBody(data)

	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", body)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusAccepted {
		return nil, errors.New("Invalid response status code")
	}

	return ioutil.ReadAll(response.Body)
}

func createBody(data interface{}) (io.Reader, error) {
	json, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(json), nil
}
