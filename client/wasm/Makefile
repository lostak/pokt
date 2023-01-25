BINARY_NAME=wasm

build: 
	GOOS=js GOARCH=wasm go build -o /home/lostak/dev/pokt/assets/main.wasm /home/lostak/dev/pokt/client/wasm/main.go

clean: 
	go clean
	rm ${BINARY_NAME}
