package main

import (
	"fmt"
	"time"
)

var balance int

func deposit() {
	currentBalance := balance
	fmt.Printf("balance before deposit %v\n", currentBalance)
	currentBalance += 1
	balance = currentBalance
	fmt.Printf("balance after deposit: %v\n", balance)
}

func withdraw() {
	currentBalance := balance
	fmt.Printf("balance before withdraw %v\n", currentBalance)
	currentBalance -= 10
	balance = currentBalance
	fmt.Printf("balance after withdraw: %v\n", balance)
}

func main() {
	balance = 10
	go deposit()
	go withdraw()
	time.Sleep(time.Millisecond)
}
