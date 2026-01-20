test:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

run:
	go run main.go

mockgen:
	mockgen -source=internal/repository/interface.go -destination=internal/repository/mocks/repository_mock.go -package=mock_repository