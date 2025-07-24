# Read WEB (HTML to Markdown)

```sh
make bin/read.wasm
wasmtime run -S http bin/read.wasm -url https://example.com
```

## from Container Registry

see: https://github.com/users/a-skua/packages/container/package/read-web

```sh
wkg oci pull -o read.wasm ghcr.io/a-skua/read-web:1.0.0
wasmtime run -S http read.wasm -url https://example.com
```
