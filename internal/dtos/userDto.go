package dtos

type UserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	MobileNo string `json:"mobileNo"`
}
