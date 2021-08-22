TEST_TOOL="gotest"

build:
	go build -v

tidy:
	go mod tidy

test: 
	$(TEST_TOOL) -v -cover -parallel 5 -failfast  ./... 
