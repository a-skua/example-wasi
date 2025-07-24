SRC := $(shell find . -name '*.go' -not -path './cmd/*' -or -name 'go.mod')

.PHONY: build
build: bin/read.wasm

bin/%.wasm: cmd/%/main.go world.wasm $(SRC)
	tinygo build -o $@ \
		--target=wasip2 --no-debug \
		--wit-package $(word 2, $^) \
		--wit-world runner \
		$<

.PHONY: gen
gen: world.wasm
	@rm -rf internal/gen
	go tool wit-bindgen-go generate \
		--world wrapper \
		--out internal/gen $<

world.wasm: wit/world.wit
	wkg wit fetch
	wkg wit build -o $@

.PHONY: fmt
fmt:
	go fmt ./...
