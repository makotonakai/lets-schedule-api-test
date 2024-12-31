package models

import (
	"time"
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
