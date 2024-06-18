package utils

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Rid     string      `json:"rid"`
	Code    int         `json:"code"`
}

func BuildResponse(status bool, message string, data interface{}, rid string) Response {
	return Response{
		Success: status,
		Message: message,
		Data:    data,
	}
}

func BuildSuccessResponse(data interface{}) Response {
	return Response{
		Data:    data,
		Success: true,
		Message: "success",
		Code:    200,
	}
}

func BuildErrorResponse(message string) Response {
	return Response{
		Data:    message,
		Success: false,
		Message: "fail",
		Code:    500,
	}
}

func BuildInternalErrorResponse() Response {
	return Response{
		Data:    "Internal Error",
		Success: false,
		Message: "fail",
		Code:    500,
	}
}
