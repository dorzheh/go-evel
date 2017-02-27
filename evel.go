package evel

/*
#cgo CFLAGS: -I/opt/evel-library/code/evel_library
#cgo LDFLAGS: -lcurl -level -L/opt/evel-library/libs/x86_64
#include <evel.h>
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

/****************************************************************************
 * Library initialization.
 *
 * Initialize the EVEL library.
 *
 * @note  This function initializes the cURL library.  Applications making use
 *        of libcurl may need to pull the initialization out of here.  Note
 *        also that this function is not threadsafe as a result - refer to
 *        libcurl's API documentation for relevant warnings.
 *
 * @sa  Matching Term function.
 *
 * @param   fqdn        string         The API's FQDN or IP address.
 * @param   port        int            The API's port.
 * @param   path        string         The optional path (may be NULL).
 * @param   topic       string         The optional topic part of the URL (may be NULL).
 * @param   secure      int            Whether to use HTTPS (0=HTTP, 1=HTTPS)
 * @param   username    string         Username for Basic Authentication of requests.
 * @param   password    string         Password for Basic Authentication of requests.
 * @param   source_type EvelSourceType The kind of node we represent.
 * @param   role        string         The role this node undertakes.
 * @param   verbosity   int            0 for normal operation, positive values for chattier logs.
 *
 * @returns
 * @retval  error  On failure and nil on success.
 *****************************************************************************/
func NewEvelInitialize(fqdn string, port int, path, topic string,
	secure int, username, password string,
	source_type EvelSourceType, role string, verbosity int) error {
	cFqdn := C.CString(fqdn)
	defer C.free(unsafe.Pointer(cFqdn))

	var cPath *C.char
	if path == "" {
		cPath = nil
	} else {
		cPath = C.CString(path)
		defer C.free(unsafe.Pointer(cPath))
	}

	var cTopic *C.char
	if topic == "" {
		cTopic = nil
	} else {
		cTopic = C.CString(topic)
		defer C.free(unsafe.Pointer(cTopic))
	}

	cUsername := C.CString(username)
	defer C.free(unsafe.Pointer(cUsername))

	cPassword := C.CString(password)
	defer C.free(unsafe.Pointer(cPassword))

	cRole := C.CString(role)
	defer C.free(unsafe.Pointer(cRole))

	res := int(C.evel_initialize(cFqdn, C.int(port), cPath, cTopic, C.int(secure),
		cUsername, cPassword, C.EVEL_SOURCE_TYPES(source_type), cRole, C.int(verbosity)))
	if res != 0 {
		return errors.New("Failed to initialize the EVEL library!")
	}
	return nil
}

/*****************************************************************************
 * Clean up the EVEL library.
 *
 * @note that at present don't expect Init/Term cycling not to leak memory!
 *
 * @returns
 * @retval  error On failure and nil on success.
 *****************************************************************************/
func EvelTerminate() error {
	if code := C.evel_terminate(); code != EVEL_SUCCESS {
		errorStr := "Failed to terminate the EVEL library!"
		EvelError(errorStr)
		return errors.New(errorStr)
	}
	return nil
}

/******************************************************************************
 * Free an event.
 *
 * Free off the event supplied.  Will recursively free all the contained
 * allocated memory.
 *
 * @note  It is safe to free a NULL pointer.
 *****************************************************************************/
func EvelFreeEvent(event interface{}) {
	cPtr := reflect.ValueOf(event).Elem()
	C.evel_free_event(unsafe.Pointer(&cPtr))
}

