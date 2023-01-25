BINARY_NAME=wasm-server-runner

build: 
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go
 
run-server: build
	cd .. && ./runner/${BINARY_NAME}

clean: 
	go clean
	rm ${BINARY_NAME}

dep:
	go mod download

lint:
	golangci-lint run --enable-all
