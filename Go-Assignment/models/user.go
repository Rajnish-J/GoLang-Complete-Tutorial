package models

import "time"

type User struct {
	UserId                int        `json:"user_id"`
	FirstName             string     `json:"first_name"`
	LastName     		  string     `json:"last_name"`
	DateOfBirth  		  string  	 `json:"dob"`
	EmailId        		  string     `json:"email"`
	MobileNumber 		  string     `json:"mobile_number"`
	AlternateMobileNumber string     `json:"alternate_mobile_number"`
	Occupation  		  string     `json:"occupation"`
    Address               string     `json:"address"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}


