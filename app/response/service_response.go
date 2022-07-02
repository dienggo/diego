package response

func Service() ServiceResponse {
	return ServiceResponse{}
}

type ServiceResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (ServiceResponse) Error(message string, data interface{}) ServiceResponse {
	return ServiceResponse{
		Status:  false,
		Message: message,
		Data:    data,
	}
}

func (ServiceResponse) Success(message string, data interface{}) ServiceResponse {
	return ServiceResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}