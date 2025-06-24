SRC := $(shell find . -name '*.go')

calculator.wasm: world.wasm $(SRC)
	@tinygo build -o $@ --target wasip2 --no-debug --wit-package $< --wit-world calculator .

world.wasm: wit/world.wit
	@wkg wit fetch
	@wkg wit build -o $@

.PHONY: init
init: world.wasm
	@go tool wit-bindgen-go generate --world calculator --out internal $<

.PHONY: build
build: calculator.wasm

.PHONY: push
push: calculator.wasm
	@wkg oci push askua9/calculator:2025.05.06 $<
