package wasmtime

// #include <wasm.h>
import "C"
import "runtime"

type Extern struct {
	_ptr   *C.wasm_extern_t
	_owner interface{}
}

type AsExtern interface {
	AsExtern() *Extern
}

func mkExtern(ptr *C.wasm_extern_t, owner interface{}) *Extern {
	f := &Extern{_ptr: ptr, _owner: owner}
	if owner == nil {
		runtime.SetFinalizer(f, func(f *Extern) {
			C.wasm_extern_delete(f._ptr)
		})
	}
	return f
}

func (e *Extern) ptr() *C.wasm_extern_t {
	ret := e._ptr
	maybeGC()
	return ret
}

func (e *Extern) owner() interface{} {
	if e._owner != nil {
		return e._owner
	}
	return e
}

// Returns the type of this export
func (e *Extern) Type() *ExternType {
	ptr := C.wasm_extern_type(e.ptr())
	runtime.KeepAlive(e)
	return mkExternType(ptr, nil)
}

// Returns a Func if this export is a function or nil otherwise
func (e *Extern) Func() *Func {
	ret := C.wasm_extern_as_func(e.ptr())
	if ret == nil {
		return nil
	} else {
		return mkFunc(ret, e.owner())
	}
}

// Returns a Global if this export is a global or nil otherwise
func (e *Extern) Global() *Global {
	ret := C.wasm_extern_as_global(e.ptr())
	if ret == nil {
		return nil
	} else {
		return mkGlobal(ret, e.owner())
	}
}

// Returns a Memory if this export is a memory or nil otherwise
func (e *Extern) Memory() *Memory {
	ret := C.wasm_extern_as_memory(e.ptr())
	if ret == nil {
		return nil
	} else {
		return mkMemory(ret, e.owner())
	}
}

// Returns a Table if this export is a table or nil otherwise
func (e *Extern) Table() *Table {
	ret := C.wasm_extern_as_table(e.ptr())
	if ret == nil {
		return nil
	} else {
		return mkTable(ret, e.owner())
	}
}

func (e *Extern) AsExtern() *Extern {
	return e
}
