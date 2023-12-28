package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func simulate(token string) {
	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	method := "POST"

	// Puts amount dynamically
	var amount float64
	fmt.Print("Enter the amount: ")
	fmt.Scan(&amount)

	amountStr := strconv.FormatFloat(amount, 'f', 2, 64)

	var phoneNumber string
	fmt.Print("Enter Phone Number: ")
	fmt.Scan(&phoneNumber)

	payload := `{
		"BusinessShortCode": 174379,
		"Password": "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMjMxMjEzMTIwNTM2",
		"Timestamp": "20231213120536",
		"TransactionType": "CustomerPayBillOnline",
		"Amount": ` + amountStr + `,
		"PartyA": ` + phoneNumber + `,
		"PartyB": 174379,
		"PhoneNumber": ` + phoneNumber + `,
		"CallBackURL": "https://mydomain.com/path",
		"AccountReference": 123456789,
		"TransactionDesc": "Payment of Items" 
	}`

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBufferString(payload))

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
