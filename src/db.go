package main

import (
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func dbAdd(uuid string, emails []string, date string, pincode string, ageGroup int) {
	emailList := strings.Join(emails, ",") // Comma separated emails

	statement, _ := database.Prepare("INSERT INTO users (UUID, emails, date, pincode, ageGroup) VALUES (?, ?, ?, ?, ?)")
	statement.Exec(uuid, emailList, date, pincode, ageGroup)
}

func dbRemove(uuid string) {
	statement, _ := database.Prepare("DELETE FROM users WHERE UUID=?")
	statement.Exec(uuid)
}

func dbGetUser(uuid string) (emailList []string, date, pincode string, ageGroup int) {
	var emails string

	statement, _ := database.Prepare("SELECT emails, date, pincode, ageGroup FROM users WHERE UUID=?")
	rows, _ := statement.Query(uuid)

	for rows.Next() {
		rows.Scan(&emails, &date, &pincode, &ageGroup)
		emailList = strings.Split(emails, ",")
	}

	return
}

func initDB() {
	var uuid string

	rows, err := database.Query("SELECT uuid FROM users")
	if err != nil {
		fmt.Println("Statement InitDB Error: ", err)
		return
	}

	for rows.Next() {
		rows.Scan(&uuid)
		go worker(uuid)
	}

	return
}
