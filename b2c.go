package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func B2CPayment(token string) {
	url := "https://sandbox.safaricom.co.ke/mpesa/b2c/v3/paymentrequest"
	method := "POST"

	// Generate a unique OriginalConversationID
	originalConversationID := uuid.New().String()

	payload := map[string]interface{}{
		"OriginatorConversationID": originalConversationID,
		"InitiatorName":            "testapi",
		"SecurityCredential":       "R5yTD7xxAedBhIzeeDuKmsdGwVWMTyBQ7lrNirObp0DKIM0XXcJhBgOzWq3LjUGqQcLqrYOWvnAf/AdpRULycrj2uqsAkHFvtpi7FlMCV4JSNNVELmrNwieEKDmtQhrVfCkbSdX57zR/dJ3hLqGNU8cm+1taBY3m+PP4mVFQrUgk1pXHNEf03ob30XL6JiL44zbe31jWFfHmKCWoSNsg2nHUNh0mOLbi+bIxM6yAQ+0u4EhQYrs4vDnQic5fi/wPRnC2eaLUHGRVwF3cwdsFXLLqzSrcHHXoaiRHxRUZZ4a4WwJYd0ouGe9aa9Dne93HtxEZQPwX9RKn5Zs/mh9w4Q==",
		"CommandID":                "BusinessPayment",
		"Amount":                   10,
		"PartyA":                   600996,
		"PartyB":                   254708374149,
		"Remarks":                  "Test remarks",
		"QueueTimeOutURL":          "https://mydomain.com/b2c/queue",
		"ResultURL":                "https://mydomain.com/b2c/result",
		"Occasion":                 "",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
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
