package model

import "time"

type UserTwoFaSms struct {
	ID        interface{} `json:"id" bson:"_id"`
	UserID    interface{} `json:"user_id" bson:"user_id"`
	PhoneID   interface{} `json:"phone_id" bson:"phone_id"`
	Confirmed bool        `json:"confirmed" bson:"confirmed"`
	CreatedAt *time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at" bson:"updated_at"`
}

func NewUserTwoFaSms(userID string, phoneID string, confirmed bool) *UserTwoFaSms{
	now := time.Now().UTC()
	return &UserTwoFaSms{
		UserID: userID,
		PhoneID: phoneID,
		Confirmed: confirmed,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}