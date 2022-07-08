package main

import (
	"fmt"
	"testing-app/services"
)

func main() {
	smsService := &services.SMSService{}
	mp := services.NewMessageProcessor(smsService)
	fmt.Println(mp.Process("Hello World"))
}
