package model

type OTP struct {
	OTP    string `json:"otp"`
	Key    string `json:"key"`
	UserID uint   `json:"userID"`
}
