package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// curl -H "Content-Type:application/json" https://gitlab.com/api/v4/projects?private_token=TOKEN -d "{ \"name\": \"newRepo\", \"namespace_id\": \"ID\" }"

type CreateRepoPayload struct {
	Name        string `json:"name"`
	NamespaceID string `json:"namespace_id"`
}

// CreateRepo takes a Payload consisting of Name (name of new Repo) and NamespaceID (ID of target group or user)
func CreateRepo(payload CreateRepoPayload, token string)  error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)
	url := fmt.Sprintf("https://gitlab.com/api/v4/projects?private_token=%s", token)
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