package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func checkCenter(ageCriteria int, pincode string, date string, emailList []string) bool {
	foundStatus := false
	var mailMsg string

	url := "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByPin?pincode=" + pincode + "&date=" + date
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept-Language", "hi_IN")

	// for true {
	fmt.Printf("Polling pincode %v for email %v for date %v\n", pincode, emailList, date)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	// Read body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	var session Session
	err = json.Unmarshal([]byte(sb), &session)
	if err != nil {
		fmt.Printf("Error Unmarshalling API Response: %s\nOriginal Response: %v\n", err, sb)
	}

	if len(session.Sessions) != 0 {
		fmt.Println("Number of sessions: ", len(session.Sessions))

		for i := 0; i < len(session.Sessions); i++ {
			name := session.Sessions[i].Name
			state := session.Sessions[i].StateName
			district := session.Sessions[i].DistrictName
			block := session.Sessions[i].BlockName
			pin := session.Sessions[i].Pincode
			capacity := session.Sessions[i].Capacity
			vaccine := session.Sessions[i].Vaccine
			age := session.Sessions[i].Age
			date := session.Sessions[i].Date
			feeType := session.Sessions[i].FeeType

			if ageCriteria == age || ageCriteria == -1 {
				mailMsg = mailMsg + fmt.Sprintf("Date: %v<br>Name: %v<br>State: %v<br>District: %v<br>Block: %v<br>Pincode: %v<br>Capacity: %v<br>Vaccine: %v<br>Minimum Age: %v<br>Fee Type: %v<br>----<br>", date, name, state, district, block, pin, capacity, vaccine, age, feeType)

				foundStatus = true
			}
		}

	}

	if foundStatus {
		currTime := time.Now()
		timeFormat := currTime.Format("15:04:05")

		timeNow := fmt.Sprintf("Information confirmed at time: %v<br><br>", timeFormat)
		endMsg := "<br><br>Register/Login and schedule appointment at https://selfregistration.cowin.gov.in/ or login using Aarogya Setup app<br><br><b>Missed the slot? Register at <a href=\"https://getcowinslot.in\">getcowinslot.in</a> to notify again</b>"

		mailSubject := "Covid Vaccine Available!\n"
		mailMsg = timeNow + mailMsg + endMsg
		sendMail(mailSubject, true, mailMsg, emailList)
	}

	defer resp.Body.Close()

	return foundStatus
}
