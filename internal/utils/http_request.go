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
	payload := strings.NewReader("{\"id_or_num\":\"42835739\",\"detail\":true}")
	fmt.Printf("payload: %s\n", payload)
	var paramsStr *strings.Reader
	var err error
	if params != nil {
		// Marshal the map to a JSON string
		jsonBytes, err := json.Marshal(params)
		if err != nil {
			log.Fatalf("Error marshaling request body: %v", err)
		}

		// Create a *strings.Reader from the JSON string
		paramsStr = strings.NewReader(string(jsonBytes))
		fmt.Printf("paramsStr: %s\n", paramsStr)
	}
	//payload := strings.NewReader("{\"id_or_num\":\"1000000\",\"detail\":true}")

	req, _ := http.NewRequest("POST", fullURL, paramsStr)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
