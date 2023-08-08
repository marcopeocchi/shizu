wasm:
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o shizu.wasm cmd/wasm/main.go