package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
)

// HTML Form
func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	pin := r.FormValue("pincode")
	dateForm := r.FormValue("date")
	email := r.FormValue("email")
	ages := r.Form["age"]
	recaptchaToken := r.FormValue("recaptchaResponse")

	if pin == "" || dateForm == "" || email == "" {
		log.Fatalln("Form Error: Some data is missing")
	}

	if !validateRecaptcha(recaptchaToken) {
		log.Fatalln("Recaptcha Validation Failed")
	}

	age := 0
	ageGroup := ""

	if len(ages) == 2 {
		age = -1
		ageGroup = "all"
	} else if contains(ages, "18") {
		age = 18
		ageGroup = "18-44"
	} else if contains(ages, "45") {
		age = 45
		ageGroup = "45+"
	} else {
		age = -1 // Check for both criteria (18+ and 45+)
		ageGroup = "all"
	}

	dates := strings.Split(dateForm, ",")
	for i := range dates {
		dates[i] = strings.TrimSpace(dates[i])
	}

	pincodes := strings.Split(pin, ",")
	for i := range pincodes {
		pincodes[i] = strings.TrimSpace(pincodes[i])
	}

	emails := strings.Split(email, ",")
	for i := range emails {
		emails[i] = strings.TrimSpace(emails[i])
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("../static/thankyou.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}

	details := Details{
		Pincode: pin,
		Date:    dateForm,
		Email:   email,
		Age:     ageGroup,
	}

	t.Execute(w, details)

	mailSubject := "Covid vaccine availability!\n"
	body := fmt.Sprintf("<html><body> Thank you for registering.<br>You will be notified when the slot opens for the pincode(s) <b>%v</b> on <b>%v</b> for <b>%v</b> years of age.<br>Till then stay home, stay safe.</body></html>", pin, dateForm, ageGroup)

	go sendMail(mailSubject, true, body, emails)

	for _, pincode := range pincodes {
		for _, date := range dates {
			if dateIsAhead(date) {
				uuid := uuid.NewString()
				dbAdd(uuid, emails, date, pincode, age)
				// go worker(uuid)

			}
		}
	}

}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func dateIsAhead(date string) bool {
	format := "2006-01-02"

	dateTmp := strings.Split(date, "-")
	date = dateTmp[2] + "-" + dateTmp[1] + "-" + dateTmp[0]

	today := time.Now()
	dateChk, err := time.Parse(format, date)
	if err != nil {
		fmt.Printf("Error parsing time: %v", err)
	}

	if today.Equal(dateChk) || today.Before(dateChk) {
		return true
	}

	return false
}

func dateIsTomorrow(date string) bool {
	format := "2006-01-02"

	dateTmp := strings.Split(date, "-")
	date = dateTmp[2] + "-" + dateTmp[1] + "-" + dateTmp[0]

	today := time.Now()
	tomorrow := today.Add(24 * time.Hour).Format(format)

	return tomorrow == date

}
