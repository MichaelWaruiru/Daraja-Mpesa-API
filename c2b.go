package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func C2BPayment(token string) {
	url := "https://sandbox.safaricom.co.ke/mpesa/c2b/v1/registerurl"
	method := "POST"

	payloadData := map[string]interface{}{
		"ShortCode":       600981,
		"ResponseType":    "Completed",
		"ConfirmationURL": "https://mydomain.com/confirmation",
		"ValidationURL":   "https://mydomain.com/validation",
	}
	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	payload := bytes.NewReader(payloadBytes)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}
