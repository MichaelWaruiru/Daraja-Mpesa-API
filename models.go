package main

// type AccessTokenResponse struct {
// 	AccessToken string `json:"access_token"`
// 	TokenType   string `json:"token_type"`
// 	ExpiresIn   string    `json:"expires_in"`
// 	RefreshToken string `json:"refresh_token,omitempty"`
// }

// RegisterURL model
type RegisterURL struct {
	ShortCode       string
	ResponseType    string
	ConfirmationURL string
	ValidationURL   string
}

// C2B model
type C2B struct {
	ShortCode     string
	CommandID     string
	Amount        string
	Msisdn        string
	BillRefNumber string
}

// B2C model
type B2C struct {
	InitiatorName          string
	SecurityCredential     string
	CommandID              string
	SenderIdentifierType   string
	ReceiveridentifierType string
	Amount                 float32
	PartyA                 int
	PartyB                 int
	Remarks                string
	AccountReference       string
	QueueTimeOutURL        string
	ResultURL              string
}

type STKPush struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            string `json:"Amount"`
	PhoneNumber       string `json:"PhoneNumber"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	QueueTimeOutURL   string `json:"QueueTimeOutURL"`
	TransactionDesc   string `json:"TransactionDesc"`
	CheckOutRequestID string `json:"CheckOutRequestID"`
}
