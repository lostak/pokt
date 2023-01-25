BINARY_NAME=pokt

build-proto:
	sh buildProto.sh

build-cobra: 
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go

build-wasm:
	sh buildWasm.sh

build-servers:
	sh buildServers.sh

build: build-proto build-cobra build-wasm build-servers 

build-no-proto: build-cobra build-wasm build-servers

run-grpc-server:
	cd runner/server/grpc && make run-server

build-and-run-grpc-server:
	cd runner/server/grpc && make build-and-run-server

run-http-server:
	cd runner/server/http && make run-server

build-and-run-http-server:
	cd runner/server/http && make build-and-run-server