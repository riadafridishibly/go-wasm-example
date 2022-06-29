build: compile/wasm run/server
ifeq (,$(wildcard ./assets/wasm_exec.js))
	cp "${shell go env GOROOT}/misc/wasm/wasm_exec.js" assets/
endif

copy/wasm_exec:
	cp "${shell go env GOROOT}/misc/wasm/wasm_exec.js" assets/

compile/wasm:
	cd cmd/wasm && GOOS=js GOARCH=wasm go build -o ../../assets/main.wasm

run/server:
	go run cmd/server/server.go
