wasm:
	tinygo build -o shizu.wasm -target wasm -opt=z cmd/wasm/main.go
	wasm-opt shizu.wasm --enable-bulk-memory -Oz -o shizu-opt.wasm
        
