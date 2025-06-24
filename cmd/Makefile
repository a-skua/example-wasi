SRC := $(shell find . -name '*.rs' -or -name '*.wit')

compose.wasm: calculator.wasm target/wasm32-wasip2/release/cmd.wasm
	wac plug -o $@ --plug $^

calculator.wasm:
	@wkg oci pull -o $@ docker.io/askua9/calculator:2025.06.25

target/wasm32-wasip2/release/cmd.wasm: $(SRC)
	@cargo build -r --target wasm32-wasip2

.PHONY: init
init:
	@wkg wit fetch
