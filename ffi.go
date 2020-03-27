package wasmtime

// #cgo !windows LDFLAGS:-lwasmtime
// #cgo windows LDFLAGS:-lwasmtime.dll
// #include <wasm.h>
// #include <wasi.h>
// #include <wasmtime.h>
import "C"

func Foo() {
  C.wasm_engine_new()
}
