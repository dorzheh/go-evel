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
)

type Event struct {
	ptr *C.struct_event_header
}

/**********************************************************************************************
 * Create a new heartbeat event.
 *
 * @note that the heartbeat is just a "naked" commonEventHeader!
 *
 * @returns Event object containing a pointer to C.struct_event_header. If the event is
 *          not used it must be released using EvelFreeEvent
 * @retval  error  Failed to create the event and nil in case the event is created successfully
 *********************************************************************************************/
func NewEvelHeartbeat() (Event, error) {
	cPtr := C.evel_new_heartbeat()
	if cPtr == nil {
		errorStr := "New heartbeat failed"
		// Send to the EVEL log
		EvelError(errorStr)
		return Event{}, errors.New(errorStr)
	}
	return Event{ptr: cPtr}, nil
}
