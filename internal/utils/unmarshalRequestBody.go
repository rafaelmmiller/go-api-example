package utils

import (
	"errors"
	"io"

	"github.com/goccy/go-json"
)

func UnmarshalRequestBody(body io.ReadCloser) (map[string]interface{}, error) {
	reqBody, err := io.ReadAll(body)
	if err != nil {
		return nil, errors.New("Invalid request body")
	}

	var jsonBody map[string]interface{}
	if err := json.Unmarshal(reqBody, &jsonBody); err != nil {
		return nil, errors.New("Invalid JSON format")
	}

	return jsonBody, nil
}
