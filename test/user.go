package user

import (
	"log"
	"time"
)

func Login(userName, password string) bool {
	log.Println(userName, password, "Login at", time.Now())
	return userName == password
}

func LogOut(userName, password string) bool {
	log.Println(userName, password, "LogOut at", time.Now())
	return userName == password
}
