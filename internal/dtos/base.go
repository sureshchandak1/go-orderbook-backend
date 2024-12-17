package dtos

type BaseRequest struct {
	Version   string  `json:"version"`
	DeviceId  string  `json:"deviceId"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type BaseResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func GetBaseResponse(status int, message string) *BaseResponse {
	return &BaseResponse{
		Status:  status,
		Message: message,
	}
}
