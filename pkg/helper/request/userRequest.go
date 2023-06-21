package request

type UserSign struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Phone           string `json:"phone_number" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type OTPreq struct {
	Phone string `json:"phoneNumber,omitempty" validate:"required"`
}
type Otpverifier struct {
	Pin   string `json:"pin,omitempty" validate:"required"`
	Phone string `json:"phoneNumber,omitempty" validate:"required"`
}

type BlockUser struct {
	UserID int    `json:"user_id"`
	Reason string `json:"reason"`
}
type Password struct {
	UserID      int    `json:"-"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type AddressReq struct {
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	Pincode     string `json:"pincode"`
	Landmark    string `json:"landmark"`
	IsDefault   *bool  `json:"is_default"`
}

/*type OTPverify struct{
	Phone string `json:"mobile_number,omitempty" validate:"required"`
	Pin string `json:"pin,omitempty" validate:"required"`
}*/
