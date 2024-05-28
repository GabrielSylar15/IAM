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
		Data:    data,
		Status:  true,
		Message: "success",
	}
}

func BuildErrorResponse(message string) Response {
	return Response{
		Data:    message,
		Status:  false,
		Message: "fail",
	}
}

func BuildInternalErrorResponse() Response {
	return Response{
		Data:    "Internal Error",
		Status:  false,
		Message: "fail",
	}
}
