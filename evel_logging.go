package evel

/*
#cgo CFLAGS: -I/opt/evel-library/code/evel_library
#cgo LDFLAGS: -lcurl -L/opt/evel-library/output/x86_64
#include <evel.h>
#include <stdlib.h>
#
// A wrapper for the log_debug function due to the fact that
// Go does not support calling variadic C functions
static void cgo_evel_log(EVEL_LOG_LEVELS level, char * format)
{
   log_debug(level, format);

}
*/
import "C"
import (
	"unsafe"
	"fmt"
)

/*******************************************************************************
 * A wrapper to EVEL_ERROR macro
 *
 * @param format string     fmt.Sprintf fromat
 * @param i      interface   parameters that should be part of formatted message
 *******************************************************************************/
func EvelError(format string, i ...interface{}) {
	cStr := C.CString("ERROR: " + fmt.Sprintf(format, i...))
	defer C.free(unsafe.Pointer(cStr))

	C.cgo_evel_log(C.EVEL_LOG_ERROR, cStr)
}

/*******************************************************************************
 * A wrapper to EVEL_INFO macro
 *
 * @param format string     fmt.Sprintf fromat
 * @param i      interface   parameters that should be part of formatted message
 *******************************************************************************/
func EvelInfo(format string, i ...interface{}) {
	cStr := C.CString(fmt.Sprintf(format, i...))
	defer C.free(unsafe.Pointer(cStr))

	C.cgo_evel_log(C.EVEL_LOG_INFO, cStr)
}

/*******************************************************************************
 * A wrapper to EVEL_DEBUG macro
 *
 * @param format string     fmt.Sprintf fromat
 * @param i      interface   parameters that should be part of formatted message
 *******************************************************************************/
func EvelDebug(format string, i ...interface{}) {
	cStr := C.CString(fmt.Sprintf(format, i...))
	defer C.free(unsafe.Pointer(cStr))

	C.cgo_evel_log(C.EVEL_LOG_DEBUG, cStr)
}




