package request

type Admin struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique;not null"`
	AdminName string `json:"admin_name" gorm:"not null" binding:"omitempty,min=4,max=12"`
	Email     string `json:"email" gorm:"not null" binding:"omitempty,email"`
	Password  string `json:"password" gorm:"not null" binding:"required,min=8,max=16"`
}

type AdminLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
