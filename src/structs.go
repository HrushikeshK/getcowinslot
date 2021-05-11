package main

import "time"

type Session struct {
	Sessions []struct {
		Name         string `json:"name"`
		StateName    string `json:"state_name"`
		DistrictName string `json:"district_name"`
		BlockName    string `json:"block_name"`
		Pincode      int    `json:"pincode"`
		Date         string `json:"date"`
		Capacity     int    `json:"available_capacity"`
		Vaccine      string `json:"vaccine"`
		Age          int    `json:"min_age_limit"`
		FeeType      string `json:"fee_type"`
	} `json:"sessions"`
}

type recaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float32   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

type Details struct {
	Pincode string
	Date    string
	Email   string
	Age     string
}

type Config struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	RecaptchaSecret string `json:"recaptcha_secret"`
}

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}
