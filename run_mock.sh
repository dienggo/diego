# Mocking Generator

mockgen -package mocks -destination=mocks/mock_http_example_repository.go -source=app/repositories/http_example_repository.go
mockgen -package mocks -destination=mocks/mock_user_repository.go -source=app/repositories/user_repository.go