package main

import "fmt"

func SendSMS(to string, message string) {
	fmt.Printf("СМС на %s: %s\n", to, message)
}
