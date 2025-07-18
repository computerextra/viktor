dev:
	go mod tidy
	air

build:
	go mod tidy
	go generate
	go build