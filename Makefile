build:
	go build -v

tidy:
	go mod tidy

test: 
	go test -v -cover -parallel 5 -failfast  ./... 
