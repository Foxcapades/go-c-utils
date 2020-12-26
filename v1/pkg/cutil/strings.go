package cutil

// #include <stdlib.h>
// #include "strings.h"
import "C"

import "unsafe"

func NewString(val string) String {
	return String{ val: val }
}

type String struct {
	val string
	ptr unsafe.Pointer
}

func (s *String) GoValue() string {
	return s.val
}

func (s *String) CValue() *C.char {
	if s.ptr != nil {
		return (*C.char)(s.ptr)
	} else {
		tmp := C.CString(s.val)

		s.ptr = unsafe.Pointer(tmp)

		return tmp
	}
}

func (s *String) Release() {
	if s.ptr != nil {
		C.free(s.ptr)
		s.ptr = nil
	}
}

func NewStringSlice(vals []string) (slice StringSlice) {
	slice.vals = make([]String, len(vals))

	for i := range vals {
		slice.vals[i] = NewString(vals[i])
	}

	return
}

type StringSlice struct {
	vals []String
	root unsafe.Pointer
}

// GoValue returns a copy of the internal string slice.
func (s *StringSlice) GoValue() []string {
	ln  := len(s.vals)
	out := make([]string, ln)

	for i := 0; i < ln; i++ {
		out[i] = s.vals[i].GoValue()
	}

	return out
}

func (s *StringSlice) CValue() **C.char {
	if s.root != nil {
		return (**C.char)(s.root)
	} else {
		ln  := len(s.vals)
		tmp := C.createStringArray(C.int(ln))

		s.root = unsafe.Pointer(tmp)

		for i := 0; i < ln; i++ {
			C.putString(tmp, C.int(i), s.vals[i].CValue())
		}

		return tmp
	}
}

func (s *StringSlice) Release() {
	for i := range s.vals {
		s.vals[i].Release()
	}

	if s.root != nil {
		C.free(s.root)
		s.root = nil
	}
}