#binary name
BINARY_NAME=rabbit
source: build
build:
		go build -o $(BINARY_NAME) -v ./src/$(BINARY_NAME)

linux:
		GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) -v ./src/$(BINARY_NAME)