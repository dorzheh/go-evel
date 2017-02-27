package evel

/*
#cgo CFLAGS: -I/opt/evel-library/code/evel_library
#cgo LDFLAGS: -lcurl -level -L/opt/evel-library/libs/x86_64
#include <evel.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
	"errors"
)

/******************************************************************************
 * Post an event.
 *
 * @note  So far as the caller is concerned, successfully posting the event
 * relinquishes all responsibility for the event - the library will take care
 * of freeing the event in due course.

 * @param event   An interface representing the event to be posted.
 *
 * @retval  nil On success
 * @retval  error On failure.
 *****************************************************************************/
func EvelPostEvent(event interface{}) error {

	var retval C.EVEL_ERR_CODES
	switch t := event.(type) {
	case Event:
		retval = C.evel_post_event(event.(Event).ptr)
	case FaultEvent:
		retval = C.evel_post_event((*C.struct_event_header)(unsafe.Pointer(event.(FaultEvent).ptr)))
	case ReportEvent:
		C.evel_post_event((*C.struct_event_header)(unsafe.Pointer(event.(ReportEvent).ptr)))
	case MeasurementEvent:
		C.evel_post_event((*C.struct_event_header)(unsafe.Pointer(event.(MeasurementEvent).ptr)))
	default:
		return fmt.Errorf("Unsupported type: %T\n", t)
	}
	if retval != EVEL_SUCCESS {
		errorStr := fmt.Sprintf("Post failed %d (%s)", retval, C.GoString(C.evel_error_string()))
		// Send to the EVEL log
		EvelError(errorStr)
		return errors.New(errorStr)
	}
	return nil
}
