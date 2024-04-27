BUILD_FLAG = -trimpath -ldflags "-s -w"
BUILD_DIR = bin

default: build

build:
	go mod tidy
	env CGO_ENABLED=0  GOOS=windows GOARCH=amd64 go build $(BUILD_FLAG) -o $(BUILD_DIR)/EndpointSearch.exe main/main.go
	env CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build $(BUILD_FLAG) -o $(BUILD_DIR)/EndpointSearch main/main.go

clean:
	rm -rf ./$(BUILD_FLAG)