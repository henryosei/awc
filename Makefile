run:
	go run cmd/cli/main.go
clean:
	rm awc
build:
	go build -o awc cmd/cli/main.go