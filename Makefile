.PHONY: fmt test build run clear

fmt:
	@go fmt ./...

test:
	@go test ./... 

build:
	@go build -o ./dist/ot_calc-linux-amd64 ./

run:
	@go run ./

clear:
	@rm -f ./dist/ot_calc-linux-amd64