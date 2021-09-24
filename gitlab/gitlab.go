package gitlab

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// curl -H "Content-Type:application/json" https://gitlab.com/api/v4/projects?private_token=TOKEN -d "{ \"name\": \"newRepo\", \"namespace_id\": \"ID\" }"

type Payload struct {
	Name        string `json:"name"`
	NamespaceID string `json:"namespace_id"`
}

func Send(payload Payload, token string)  error {
	data := payload
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)
	url := "https://gitlab.com/api/v4/projects?private_token=" + token
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
}