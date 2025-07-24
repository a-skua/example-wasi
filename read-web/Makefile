SRC := $(shell find . -name '*.go' -not -path './cmd/*' -or -name 'go.mod')

.PHONY: build
build: bin/read.wasm bin/omit-outgoing-handler.wasm bin/omit-read.wasm

bin/omit-read.wasm: bin/read.wasm bin/omit-outgoing-handler.wasm
	wac plug $< --plug $(word 2, $^) -o $@

bin/%.wasm: cmd/%/main.go world.wasm gen-% $(SRC)
	tinygo build -o $@ \
		--target=wasip2 --no-debug \
		--wit-package $(word 2, $^) \
		--wit-world $* \
		$<

.PHONY: gen-%
gen-%: world.wasm
	go tool wit-bindgen-go generate \
		--world $* \
		--out internal/gen $<

world.wasm: wit/world.wit
	wkg wit fetch
	wkg wit build -o $@
	@rm -rf internal/gen

.PHONY: fmt
fmt:
	go fmt ./...
