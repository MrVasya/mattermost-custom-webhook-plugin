package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendPostRequest(url string, data map[string]string) map[string]interface{} {
	jsonValue, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result
}