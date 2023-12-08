wasm:
	tinygo build -o shizu.wasm -target wasm -opt=z cmd/wasm/main.go
	