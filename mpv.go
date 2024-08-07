package mpv

/*
#if defined(__has_include) && __has_include(<libmpv/client.h>)
#include <libmpv/client.h>
#else
#include <mpv/client.h>
#endif
#include <stdlib.h>
#cgo LDFLAGS: -lmpv

char** makeCharArray(int size) {
    return calloc(sizeof(char*), size);
}

void setStringArray(char** a, int i, char* s) {
    a[i] = s;
}
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

// Format .
type Format int

const (
	FORMAT_NONE       Format = C.MPV_FORMAT_NONE
	FORMAT_STRING     Format = C.MPV_FORMAT_STRING
	FORMAT_OSD_STRING Format = C.MPV_FORMAT_OSD_STRING
	FORMAT_FLAG       Format = C.MPV_FORMAT_FLAG
	FORMAT_INT64      Format = C.MPV_FORMAT_INT64
	FORMAT_DOUBLE     Format = C.MPV_FORMAT_DOUBLE
	FORMAT_NODE       Format = C.MPV_FORMAT_NODE
	FORMAT_NODE_ARRAY Format = C.MPV_FORMAT_NODE_ARRAY
	FORMAT_NODE_MAP   Format = C.MPV_FORMAT_NODE_MAP
	FORMAT_BYTE_ARRAY Format = C.MPV_FORMAT_BYTE_ARRAY
)

// Mpv represents an mpv client.
type Mpv struct {
	handle *C.mpv_handle
}

// Create creates a new MPV struct.
func Create() *Mpv {
	return &Mpv{C.mpv_create()}
}

// MPVHandle returns the pointer to the mpv_handle struct
// for invoking native MPV APIs
func (m *Mpv) MPVHandle() unsafe.Pointer {
	return unsafe.Pointer(m.handle)
}

// ClientName .
func (m *Mpv) ClientName() string {
	return C.GoString(C.mpv_client_name(m.handle))
}

// GetTimeUS .
func (m *Mpv) GetTimeUS() int64 {
	return int64(C.mpv_get_time_us(m.handle))
}

// Initialize runs mpv_initialize and returns any errors that might occur.
func (m *Mpv) Initialize() error {
	return newError(C.mpv_initialize(m.handle))
}

// Command runs the specified command, returning an error if something goes wrong.
func (m *Mpv) Command(command []string) error {
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

	if err := newError(C.mpv_command(m.handle, arr)); err != nil {
		return fmt.Errorf("failed to execute command %v: %w", command, err)
	}
	return nil
}

// CommandString runs the given command string, this string is parsed internally by mpv.
func (m *Mpv) CommandString(command string) error {
	return newError(C.mpv_command_string(m.handle, C.CString(command)))
}

// CommandNode runs the given command node.
func (m *Mpv) CommandNode(args Node, result *Node) error {
	return newError(C.mpv_command_node(m.handle, args.CNode(), result.CNode()))
}

// LoadConfigFile loads the given config file.
func (m *Mpv) LoadConfigFile(fn string) error {
	return newError(C.mpv_load_config_file(m.handle, C.CString(fn)))
}

// SetProperty sets the client property according to the given format.
func (m *Mpv) SetProperty(name string, format Format, data interface{}) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	dataPtr, freeData := convertData(format, data)
	if freeData != nil {
		defer freeData()
	}

	return newError(C.mpv_set_property(m.handle, cName, C.mpv_format(format), dataPtr))
}

// GetProperty returns the value of the property according to the given format.
func (m *Mpv) GetProperty(name string, format Format) (interface{}, error) {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	switch format {
	case FORMAT_NONE:
		err := newError(C.mpv_get_property(m.handle, n, C.mpv_format(format), nil))
		if err != nil {
			return nil, err
		}
		return nil, nil
	case FORMAT_STRING, FORMAT_OSD_STRING:
		var result *C.char
		err := newError(C.mpv_get_property(m.handle, n, C.mpv_format(format), unsafe.Pointer(&result)))
		if err != nil {
			return nil, err
		}
		defer C.mpv_free(unsafe.Pointer(result))
		return C.GoString(result), nil
	case FORMAT_FLAG:
		var result C.int
		err := newError(C.mpv_get_property(m.handle, n, C.mpv_format(format), unsafe.Pointer(&result)))
		if err != nil {
			return nil, err
		}
		return result == 1, nil
	case FORMAT_INT64:
		var result C.int64_t
		err := newError(C.mpv_get_property(m.handle, n, C.mpv_format(format), unsafe.Pointer(&result)))
		if err != nil {
			return nil, err
		}
		return int64(result), nil
	case FORMAT_DOUBLE:
		var result C.double
		err := newError(C.mpv_get_property(m.handle, n, C.mpv_format(format), unsafe.Pointer(&result)))
		if err != nil {
			return nil, err
		}
		return float64(result), nil
	case FORMAT_NODE:
		var result C.mpv_node
		err := newError(C.mpv_get_property(m.handle, n, C.mpv_format(format), unsafe.Pointer(&result)))
		if err != nil {
			return nil, err
		}
		defer C.mpv_free_node_contents(&result)
		return NewNode(&result), nil
	case FORMAT_NODE_ARRAY:
		fallthrough
	case FORMAT_NODE_MAP:
		return nil, errors.New("unsupported format for mpv_get_property")
	default:
		return nil, ERROR_UNKNOWN_FORMAT

	}
}

// SetPropertyString sets the property to the given string.
func (m *Mpv) SetPropertyString(name, value string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	return newError(C.mpv_set_property_string(m.handle, cName, cValue))
}

// GetPropertyString returns the value of the property as a string. If the property is empty,
// an empty string is returned.
func (m *Mpv) GetPropertyString(name string) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	str := C.mpv_get_property_string(m.handle, cName)
	if str == nil {
		return ""
	}
	defer C.mpv_free(unsafe.Pointer(str))
	return C.GoString(str)
}

// GetPropertyOsdString returns the value of the property as a string formatted for mpv's
// on-screen display.
func (m *Mpv) GetPropertyOsdString(name string) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	str := C.mpv_get_property_osd_string(m.handle, cName)
	defer C.mpv_free(unsafe.Pointer(str))
	if str != nil {
		return C.GoString(str)
	} else {
		return ""
	}
}

// ObserveProperty .
func (m *Mpv) ObserveProperty(replyUserdata uint64, name string, format Format) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return newError(C.mpv_observe_property(m.handle, C.uint64_t(replyUserdata), cName, C.mpv_format(format)))
}

// UnobserveProperty .
func (m *Mpv) UnobserveProperty(replyUserdata uint64) error {
	return newError(C.mpv_unobserve_property(m.handle, C.uint64_t(replyUserdata)))
}

// SetOption sets the given option according to the given format.
func (m *Mpv) SetOption(name string, format Format, data interface{}) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	dataPtr, freeData := convertData(format, data)
	if freeData != nil {
		defer freeData()
	}

	return newError(C.mpv_set_option(m.handle, cName, C.mpv_format(format), dataPtr))
}

// SetOptionString sets the option to the given string.
func (m *Mpv) SetOptionString(name, value string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	return newError(C.mpv_set_option_string(m.handle, cName, cValue))
}

// WaitEvent calls mpv_wait_event and returns the result as an Event struct.
func (m *Mpv) WaitEvent(timeout float32) *Event {
	event := C.mpv_wait_event(m.handle, C.double(timeout))
	if event == nil {
		return nil
	}
	return &Event{
		Event_Id:       EventId(event.event_id),
		Data:           unsafe.Pointer(event.data),
		Reply_Userdata: uint64(event.reply_userdata),
		Error:          newError(event.error),
	}
}

// RequestEvent .
func (m *Mpv) RequestEvent(event EventId, enable bool) error {
	var enable_ C.int = 0
	if enable {
		enable_ = 1
	}
	return newError(C.mpv_request_event(m.handle, C.mpv_event_id(event), enable_))
}

// RequestLogMessages .
func (m *Mpv) RequestLogMessages(level string) error {
	cLevel := C.CString(level)
	defer C.free(unsafe.Pointer(cLevel))
	return newError(C.mpv_request_log_messages(m.handle, cLevel))
}

// Wakeup .
func (m *Mpv) Wakeup() {
	C.mpv_wakeup(m.handle)
}

// GetWakeupPipe .
func (m *Mpv) GetWakeupPipe() int {
	return int(C.mpv_get_wakeup_pipe(m.handle))
}

// TerminateDestroy terminates mpv and destroys the client.
func (m *Mpv) TerminateDestroy() {
	C.mpv_terminate_destroy(m.handle)
}

// convertData converts the data according to the given format and returns an unsafe.Pointer
// for use in SetOption and SetProperty.
// Note: If conversion lead to a memory allocation, you need to defer call the returned function.
func convertData(format Format, data interface{}) (unsafe.Pointer, func()) {
	switch format {
	case FORMAT_NONE:
		return nil, nil
	case FORMAT_STRING, FORMAT_OSD_STRING:
		str := C.CString(data.(string))
		return unsafe.Pointer(str), func() { C.free(unsafe.Pointer(str)) }
	case FORMAT_FLAG:
		var val C.int
		if data.(bool) {
			val = 1
		} else {
			val = 0
		}
		return unsafe.Pointer(&val), nil
	case FORMAT_INT64:
		i, ok := data.(int64)
		if !ok {
			i = int64(data.(int))
		}
		val := C.int64_t(i)
		return unsafe.Pointer(&val), nil
	case FORMAT_DOUBLE:
		val := C.double(data.(float64))
		return unsafe.Pointer(&val), nil
	case FORMAT_NODE:
		node := data.(*Node)
		return unsafe.Pointer(node.CNode()), nil
	case FORMAT_NODE_MAP:
		return unsafe.Pointer(CNodeMap(data.(map[string]*Node))), nil
	case FORMAT_NODE_ARRAY:
		return unsafe.Pointer(CNodeList(data.([]*Node))), nil
	default:
		return nil, nil
	}
}
