package model

type ApiBody struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	EmailAddress string `json:"email_address"`
	Msisdn string `json:"msisdn"`
	Otp int `json:"otp"`
}

func (body ApiBody ) GetFirstName() string {
	return body.FirstName
}

func (body ApiBody) GetLastName()  string {
	return body.LastName
}

func (body ApiBody) GetEmailAddress() string {
	return body.EmailAddress
}

func (body ApiBody)  GetMsisdn() string {
	return body.Msisdn
}

func (body ApiBody) GetOtp() int  {
	return body.Otp
}

type Register interface {
	GetFirstName() string
	GetLastName() string
	GetMsisdn() string
	GetEmailAddress() string
}

type Login interface {
	GetMsisdn() string
	GetOtp() int
}

