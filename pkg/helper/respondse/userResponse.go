package respondse

import "time"


type UserValue struct{
	ID uint `json:"id" gorm:"unique;not null"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Created time.Time `json:"created_time"`
}

