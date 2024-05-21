package utils

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Rid     string      `json:"rid"`
}

func BuildResponse(status bool, message string, data interface{}, rid string) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func BuildSuccessResponse(data interface{}) Response {
	return Response{
		Data: data,
	}
}
