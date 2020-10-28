// +build cosmwasm

package std

import (
	"github.com/cosmwasm/cosmwasm-go/std/ezjson"
	"unsafe"
)

func Build_ErrResponse(msg string) string {
	return `{"generic_err":{"msg":"` + msg + `","backtrace":null}}`
}

func StdErrResult(msg string) unsafe.Pointer {
	return Package_message([]byte(`{"Err":` + Build_ErrResponse(msg) + `}`))
}

func make_dependencies() Extern {
	return Extern{
		EStorage: ExternalStorage{},
		EApi:     ExternalApi{},
		EQuerier: ExternalQuerier{},
	}
}

// ========== init ==============
func DoInit(initFn func(*Extern, Env, MessageInfo, []byte) (*InitResultOk, *CosmosResponseError), envPtr, infoPtr, msgPtr uint32) unsafe.Pointer {
	msgData := Translate_range_custom(uintptr(msgPtr))

	env := Env{}
	envData := TranslateToSlice(uintptr(envPtr))
	err := ezjson.Unmarshal(envData, &env)
	if err != nil {
		return StdErrResult("Cannot Parse Env")
	}
	info := MessageInfo{}
	infoData := TranslateToSlice(uintptr(infoPtr))
	err := ezjson.Unmarshal(infoData, &info)
	if err != nil {
		return StdErrResult("Cannot Parse Info")
	}

	deps := make_dependencies()
	ok, ers := initFn(&deps, env, info, msgData)
	if ok == nil {
		b, e := ezjson.Marshal(ers)
		if e != nil {
			return StdErrResult("Marshal error failed : " + e.Error())
		}
		return Package_message(b)
	}

	data, err := ezjson.Marshal(*ok)
	if err != nil {
		return StdErrResult("Failed to marshal init response to []byte: " + err.Error())
	}
	return Package_message(data)
}

// ========= handler ============
func DoHandler(handlerFn func(*Extern, Env, MessageInfo, []byte) (*HandleResultOk, *CosmosResponseError), envPtr, infoPtr, msgPtr uint32) unsafe.Pointer {
	msgData := Translate_range_custom(uintptr(msgPtr))
	env := Env{}
	envData := TranslateToSlice(uintptr(envPtr))
	err := ezjson.Unmarshal(envData, &env)
	if err != nil {
		return StdErrResult("Cannot Parse Env")
	}
	info := MessageInfo{}
	infoData := TranslateToSlice(uintptr(infoPtr))
	err := ezjson.Unmarshal(infoData, &info)
	if err != nil {
		return StdErrResult("Cannot Parse Info")
	}

	deps := make_dependencies()
	ok, ers := handlerFn(&deps, env, info, msgData)
	if ok == nil {
		b, e := ezjson.Marshal(ers)
		if e != nil {
			return StdErrResult("Marshal error failed : " + e.Error())
		}
		return Package_message(b)
	}

	data, err := ezjson.Marshal(*ok)
	if err != nil {
		return StdErrResult("Failed to marshal init response to []byte: " + err.Error())
	}
	return Package_message(data)
}

// =========== query ===================
func DoQuery(queryFn func(*Extern, Env, []byte) (*QueryResponseOk, *CosmosResponseError), envPtr, msgPtr uint32) unsafe.Pointer {
	msgData := Translate_range_custom(uintptr(msgPtr))
	env := Env{}
	envData := TranslateToSlice(uintptr(envPtr))
	err := ezjson.Unmarshal(envData, &env)
	if err != nil {
		return StdErrResult("Cannot Parse Env")
	}

	deps := make_dependencies()
	ok, ers := queryFn(&deps, env, msgData)
	if ok == nil {
		b, e := ezjson.Marshal(ers)
		if e != nil {
			return StdErrResult("Marshal error failed : " + e.Error())
		}
		return Package_message(b)
	}

	data, err := ezjson.Marshal(*ok)
	if err != nil {
		return StdErrResult("Failed to marshal init response to []byte: " + err.Error())
	}
	return Package_message(data)
}

//export cosmwasm_vm_version_3
func cosmwasm_vm_version_3() {}

//export allocate
func allocate(size uint32) unsafe.Pointer {
	ptr, _ := Build_region(size, 0)
	return ptr
}

//export deallocate
func deallocate(pointer unsafe.Pointer) {
	Deallocate(pointer)
}
