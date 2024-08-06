package mpv

/*
#include <mpv/client.h>
#include <stdlib.h>
char** makeCharArray(int size);
void setStringArray(char** a, int i, char* s);
*/
import "C"

import (
	"unsafe"
)

// SetPropertyAsync .
func (m *Mpv) SetPropertyAsync(name string, replyUserdata uint64, format Format, data interface{}) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	dataPtr, freeData := convertData(format, data)
	if freeData != nil {
		defer freeData()
	}

	return newError(C.mpv_set_property_async(m.handle, C.uint64_t(replyUserdata), cName, C.mpv_format(format), dataPtr))
}

// GetPropertyAsync .
func (m *Mpv) GetPropertyAsync(name string, replyUserdata uint64, format Format) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return newError(C.mpv_get_property_async(m.handle, C.uint64_t(replyUserdata), cName, C.mpv_format(format)))
}

// CommandAsync .
func (m *Mpv) CommandAsync(replyUserdata uint64, command []string) error {
	arr := C.makeCharArray(C.int(len(command) + 1))
	if arr == nil {
		return ERROR_NOMEM
	}
	defer C.free(unsafe.Pointer(arr))

	cStrings := make([]*C.char, len(command))
	for i, s := range command {
		val := C.CString(s)
		cStrings[i] = val
		C.setStringArray(arr, C.int(i), val)
	}

	defer func() {
		for _, cStr := range cStrings {
			C.free(unsafe.Pointer(cStr))
		}
	}()

	return newError(C.mpv_command_async(m.handle, C.uint64_t(replyUserdata), arr))
}

// CommandNodeAsync .
func (m *Mpv) CommandNodeAsync(replyUserdata uint64, args Node) error {
	return newError(C.mpv_command_node_async(m.handle, C.uint64_t(replyUserdata), args.CNode()))
}
