package base

import "go_base_project/app/response"

type ServiceInterface interface {
	Do() response.ServiceResponse
}