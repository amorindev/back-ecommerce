package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// puede ser data de tipo map string string, o crear dos funciiones?
// de momento simple que acepte soo un valor
func (a *Adapter) CreateBranchLink(linkData map[string]string) (string, error) {
	apiUrl := "https://api2.branch.io/v1/url"
	branchKey := "key_test_fte3v7bW0GCvUYhZ0WmOSecbDAbmLBgy"

	payload := map[string]interface{}{
		"branch_key": branchKey,
		"data": linkData,
		/* "datad": map[string]string{
			key: value,
			
			//"user_id": "12",
		}, */
		
	}

	jsonData, _ := json.Marshal(payload)

	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var responseBody map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return "", fmt.Errorf("responce body error: %w", err)
	}

	link, ok := responseBody["url"].(string)
	if !ok {
		return "", fmt.Errorf("error creando el enlace de branch")
	}
	return link, nil
}
