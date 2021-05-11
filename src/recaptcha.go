package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func validateRecaptcha(token string) bool {

	log.Println("Recaptcha Validation Start")

	//Get Recaptcha Secret
	var config Config
	fd, err := os.Open("../config.json")
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	creds, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println(err)
	}

	credsString := fmt.Sprintf("%s", creds)

	err = json.Unmarshal([]byte(credsString), &config)

	// check recaptcha Response score
	recaptchaURL := "https://www.google.com/recaptcha/api/siteverify"
	recaptchaSecret := config.RecaptchaSecret

	httpClient := http.Client{Timeout: time.Second * 10}
	resp, err := httpClient.PostForm(recaptchaURL, url.Values{
		"secret":   {recaptchaSecret},
		"response": {token},
	})

	if err != nil {
		log.Printf("Error: %v", err)
		return false
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Recaptcha Failed")
	}

	var response recaptchaResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln("Recaptcha Unmarshalling failed")
	}

	log.Println("Recaptcha Finish")

	return true

}
