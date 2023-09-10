package app

import "encoding/json"

// IService : base of service and have to implement on created new service
type IService interface {
	Do() ServiceResponse
}

// response : base response on this app
type response struct {
	err    error
	result any
}

// Error : getter error data
func (r response) Error() error { return r.err }

// Result : getter result data
func (r response) Result() any { return r.result }

// ResultParse : getter result data on parsed data on model reference
func (r response) ResultParse(data interface{}) error {
	if r.err != nil {
		return r.err
	}
	marshal, err := json.Marshal(r.result)
	if err != nil {
		return err
	}
	err2 := json.Unmarshal([]byte(marshal), &data)
	if err2 != nil {
		return err2
	}
	return nil
}

// ServiceResponse : for service response reason
type ServiceResponse = response

// SRSuccess : return result without error data
func SRSuccess(result any) ServiceResponse {
	return ServiceResponse{err: nil, result: result}
}

// SRFail : return result with error data
func SRFail(err error, result any) ServiceResponse {
	return ServiceResponse{err: err, result: result}
}

// Service : for service extends reason
type Service ServiceResponse

func (Service) Success(result any) ServiceResponse {
	return SRSuccess(result)
}
func (Service) Error(err error, result any) ServiceResponse {
	return SRFail(err, result)
}
func (s Service) DataParse(target any) error {
	return new(response).ResultParse(&target)
}
