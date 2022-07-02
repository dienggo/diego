package response

func Api() apiResponse {
	return apiResponse{}
}

type apiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (apiResponse) Error(message string, data interface{}) apiResponse {
	return apiResponse{
		Status:  false,
		Message: message,
		Data:    data,
	}
}

func (apiResponse) Success(message string, data interface{}) apiResponse {
	return apiResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}