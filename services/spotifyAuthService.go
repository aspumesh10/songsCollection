package services

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

/**
 * This function is responsible for getting access token using client credentials flow in spotify
 */
func GetAccessToken() string {
	var tempToken token
	url := "https://accounts.spotify.com/api/token"
	method := "POST"

	payload := strings.NewReader("grant_type=client_credentials")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
		return ""
	}

	encodedText := base64.StdEncoding.EncodeToString([]byte(os.Getenv("client_id") + ":" + os.Getenv("secret_key")))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+encodedText)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	errMarshal := json.Unmarshal(body, &tempToken)
	if errMarshal != nil {
		return ""
	}
	return tempToken.AccessToken
}
