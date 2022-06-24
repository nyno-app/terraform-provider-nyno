package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type LoginRequestPayload struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Organization string `json:"organization"`
}

type LoginResponse struct {
	SessionToken string `json:"sessionToken"`
}

func getSessionToken(base_url string, username string, password string, organization string) string {
	log.Print("Getting token")
	client := &http.Client{Timeout: 10 * time.Second}

	payload := &LoginRequestPayload{
		Username:     username,
		Password:     password,
		Organization: organization,
	}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		log.Print(err)
		return "Error during login - before building request"
	}
	body := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%[1]s/auth/credentials", base_url), body)

	if err != nil {
		return "Error during login - before request"
	}

	req.Header.Set("Content-Type", "application/json")
	r, err := client.Do(req)
	if err != nil {
		return "Error during login - after request"
	}

	if r.StatusCode != 200 {
		var response *ResponseError
		err = json.NewDecoder(r.Body).Decode(&response)

		if response == nil {
			return fmt.Sprintf("Unable to update role. Status Code: %v", r.StatusCode)
		}

		return fmt.Sprintf("Unable to update role. Status Code: %v. Message: %s", r.StatusCode, response.Error)
	}

	var response LoginResponse
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return "Unable to decode response"
	}

	return response.SessionToken
}
