// Copyright 2023 The YockCloud Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cgo

import "C"
import (
	"encoding/json"
	"fmt"
	"unsafe"
)

func CastPtr[T any](v unsafe.Pointer) *T {
	return (*T)(unsafe.Pointer(v))
}

func Malloc[T any](stu T) *T {
	return (*T)(C.malloc(C.size_t(unsafe.Sizeof(stu))))
}

func JSON(a any) *C.char {
	data, err := json.Marshal(a)
	if err != nil {
		return C.CString(fmt.Sprintf(`{"err": "%s"}`, err.Error()))
	}
	return C.CString(string(data))
}

type MAP map[string]any

func Unmarshal[T any](c *C.char) (T, *C.char) {
	var opt T
	err := json.Unmarshal([]byte(C.GoString(c)), &opt)
	if err != nil {
		return opt, C.CString(err.Error())
	}
	return opt, nil
}
