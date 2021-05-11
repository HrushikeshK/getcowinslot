package main

import (
	"fmt"
	"time"
)

func worker(uuid string) {
	emailList, date, pincode, age := dbGetUser(uuid)
	for true {
		if dateIsTomorrow(date) || age == 45 {
			if checkCenter(age, pincode, date, emailList) || !dateIsAhead(date) {
				dbRemove(uuid)
				break
			}
		} else {
			fmt.Println("Later date: ", date)
		}

		time.Sleep(120 * time.Second)
	}

}
