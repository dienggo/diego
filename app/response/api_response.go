package response

func Api() ApiResponse {
	return ApiResponse{}
}

type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (ApiResponse) Error(message string, data interface{}) ApiResponse {
	return ApiResponse{
		Status:  false,
		Message: message,
		Data:    data,
	}
}

func (ApiResponse) Success(message string, data interface{}) ApiResponse {
	return ApiResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}
