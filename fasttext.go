package fasttext

// #cgo CXXFLAGS: -I${SRCDIR}/fastText/src -I${SRCDIR} -std=c++14
// #cgo LDFLAGS: -lstdc++
// #include <stdio.h>
// #include <stdlib.h>
// #include "cbits.h"
import "C"

import (
	"encoding/json"
	"unsafe"
)

type Model struct {
	path   string
	handle C.FastTextHandle
}

func Open(path string) *Model {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	return &Model{
		path:   path,
		handle: C.NewHandle(cpath),
	}
}

func (handle *Model) Close() error {
	if handle == nil {
		return nil
	}
	C.DeleteHandle(handle.handle)
	return nil
}

func (handle *Model) Predict(query string) (Predictions, error) {
	cquery := C.CString(query)
	defer C.free(unsafe.Pointer(cquery))

	r := C.Predict(handle.handle, cquery)

	defer C.free(unsafe.Pointer(r))
	js := C.GoString(r)

	predictions := []Prediction{}
	err := json.Unmarshal([]byte(js), &predictions)
	if err != nil {
		return nil, err
	}

	return predictions, nil
}
