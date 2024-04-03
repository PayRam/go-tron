package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func MakeRequest(baseURL string, endpoint string, params interface{}) ([]byte, error) {
	fullURL := fmt.Sprintf("%s%s", baseURL, endpoint)
	var paramsStr io.Reader
	var err error
	if params != nil {
		// Marshal the map to a JSON string
		jsonBytes, err := json.Marshal(params)
		if err != nil {
			log.Println("Error marshaling request body: %v", err)
			return nil, err
		}

		// Create a *strings.Reader from the JSON string
		paramsStr = strings.NewReader(string(jsonBytes))
	}
	//payload := strings.NewReader("{\"id_or_num\":\"1000000\",\"detail\":true}")

	req, httpReqError := http.NewRequest("POST", fullURL, paramsStr)

	if httpReqError != nil {
		log.Println("Error creating http request: ", httpReqError)
		return nil, httpReqError
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, clientError := http.DefaultClient.Do(req)

	if clientError != nil {
		log.Println("Error making request: ", clientError)
		return nil, clientError
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return nil, err
	}

	return body, nil
}
