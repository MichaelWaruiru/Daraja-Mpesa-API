package main

import (
	"fmt"
	// "time"
)

func main() {
	token, err := accessToken()
	if err != nil {
		fmt.Println("Error getting access token:", err)
		return
	}

	// fmt.Println("Executing C2B payment...")
	// C2BPayment(token)
	// time.Sleep(2 * time.Second)
	// if err != nil {
	// 	fmt.Println("Error getting the access token: ", err)
	// 	return
	// }

	fmt.Println("Executing B2C payment...")
	B2CPayment(token)

	// fmt.Println("Pushing STK...")
	// simulate(token)
}
