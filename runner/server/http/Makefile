BINARY_NAME=server

build: 
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go

# Server is ran from base dir 
run-server: 
	cd ../../.. && ./runner/server/http/${BINARY_NAME}

build-and-run-server: build run-server

clean: 
	go clean
	rm ${BINARY_NAME}