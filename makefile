run:
	go run main.go

tidy:
	go mod tidy

test:
	go test -v ./...