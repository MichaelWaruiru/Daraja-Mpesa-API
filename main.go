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

	// c2bRequest := C2B{
	// 	ShortCode:     "600981",
	// 	CommandID:     "CustomerPayBillOnline",
	// 	Amount:        "100",
	// 	Msisdn:        "254719453367",
	// 	BillRefNumber: "123456",
	// }

	// fmt.Println("Executing C2B payment...")
	// C2BPayment(token)
	// time.Sleep(2 * time.Second)
	// if err != nil {
	// 	fmt.Println("Error getting the access token: ", err)
	// 	return
	// }

	// // b2cRequest := B2C{
	// // 	InitiatorName:      "testapi",
	// // 	SecurityCredential: "R5yTD7xxAedBhIzeeDuKmsdGwVWMTyBQ7lrNirObp0DKIM0XXcJhBgOzWq3LjUGqQcLqrYOWvnAf",
	// // 	CommandID:          "BusinessPayment",
	// // 	Amount:             1,
	// // 	PartyA:             600996,
	// // 	PartyB:             254719453367,
	// // 	Remarks:            "Okay",
	// // 	QueueTimeOutURL:    "https://mydomain.com/b2c/queue",
	// // 	ResultURL:          "https://mydomain.com/b2c/result",
	// // }

	// fmt.Println("Executing B2C payment...")
	// B2CPayment(token)

	fmt.Println("Pushing STK...")
	simulate(token)
}
