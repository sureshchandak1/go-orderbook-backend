package dtos

type SignUpReqDto struct {
	SignUpReqData `json:"data" validate:"required"`
	BaseRequest   `json:"request" validate:"required"`
}

type SignUpReqData struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	MobileNo string `json:"mobileNo" validate:"required"`
}

type SignUpResDto struct {
	BaseResponse  `json:"response"`
	SignUpResData `json:"data"`
}

type SignUpResData struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	MobileNo string `json:"mobileNo"`
}

type LoginReqDto struct {
	LoginReqData `json:"data" validate:"required"`
	BaseRequest  `json:"request" validate:"required"`
}

type LoginReqData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResDto struct {
	BaseResponse `json:"response"`
	LoginResData `json:"data"`
}

type LoginResData struct {
	Token string `json:"token"`
}
