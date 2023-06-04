package base

// Service : base of service and have to implement on created new service
type Service interface {
	Do() ServiceResponse
}
