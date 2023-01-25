# USAGE: bash ./crossCompile.sh [path-to-project-base-dir]
GOOS=js GOARCH=wasm go build -o $1/assets/json.wasm $1/client/wasm/main.go