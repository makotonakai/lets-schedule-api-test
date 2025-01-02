package models

import (
	"time"
	"github.com/makotonakai/lets-schedule-api-test/config"
)

type User struct {	
	Id int `json:"id"`	
	UserName string `json:"user_name"`
	EmailAddress string `json:"email_address"`
	Password string `json:"password"`
	IsAdmin	bool `json:"is_admin"`
	CanLogin bool `json:"can_login"`	
	CreatedAt time.Time `json:"created_at"`	
	UpdatedAt time.Time `json:"updated_at"`	
}

func GetUserIdFromEmailAddress(users map[string]*User, address string) (int, error) {
	for _, user := range users {
		if user.EmailAddress == address {
			return user.Id, nil
		}
	}
	return -1, config.ErrEmailAddressNotFound
}
