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
	"unsafe"
)

type FaultEvent struct {
	ptr *C.struct_event_fault
}

/**************************************************************************/ /**
 * Create a new fault event.
 *
 * @note    The mandatory fields on the Fault must be supplied to this factory
 *          function and are immutable once set.  Optional fields have explicit
 *          setter functions, but again values may only be set once so that the
 *          Fault has immutable properties.
 * @param   condition   The condition indicated by the Fault.
 * @param   specific_problem  The specific problem triggering the fault.
 * @param   priority    The priority of the event.
 * @param   severity    The severity of the Fault.
 * @returns pointer to the newly manufactured C.struct_event_fault.  If the event is
 *          not used (i.e. posted) it must be released using EvelFreeEvent.
 * @retval  FaultEvent  the new event
 * @retval  error       error generated during the event creation
 ********
 *********************************************************************/
func NewEvelFault(condition, specific_problem string, priority EvelEventPriorities,
	severity EvelFaultSeverities) (FaultEvent, error) {
	cCondition := C.CString(condition)
	defer C.free(unsafe.Pointer(cCondition))

	cSpecificProblem := C.CString(specific_problem)
	defer C.free(unsafe.Pointer(cSpecificProblem))

	faultPtr := C.evel_new_fault(cCondition, cSpecificProblem,
		C.EVEL_EVENT_PRIORITIES(priority),
		C.EVEL_FAULT_SEVERITIES(severity))
	if faultPtr == nil {
		errorStr := "Error in creating a new fault event."
		// Send to the EVEL log
		EvelError(errorStr)
		return FaultEvent{}, errors.New(errorStr)
	}
	return FaultEvent{ptr: faultPtr}, nil
}

/******************************************************************************
 * Add an additional value name/value pair to the Fault.
 *
 * The library takes a copy so the caller does not have to preserve values
 * after the function returns.
 *
 * @param name      Attribute's name.
 * @param value     Attribute's value.
 *****************************************************************************/
func (f *FaultEvent) EvelFaultAddlInfoAdd(name, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.evel_fault_addl_info_add(f.ptr, cName, cValue)
}

/******************************************************************************
 * Set the Alarm Interface A property of the Fault.
 *
 * @note  The property is treated as immutable: it is only valid to call
 *        the setter once.  However, we don't assert if the caller tries to
 *        overwrite, just ignoring the update instead.
 *
 * @param fault_iface  string  The Alarm Interface A to be set.
 *****************************************************************************/
func (f *FaultEvent) EvelFaultInterfaceSet(fault_iface string) {
	cFaultIface := C.CString(fault_iface)
	defer C.free(unsafe.Pointer(cFaultIface))

	C.evel_fault_interface_set(f.ptr, cFaultIface)
}

/*****************************************************************************
 * Set the Event Type property of the Fault.
 *
 * @note  The property is treated as immutable: it is only valid to call
 *        the setter once.  However, we don't assert if the caller tries to
 *        overwrite, just ignoring the update instead.
 *
 * @param falt_type  string  The Event Type to be set.
 *****************************************************************************/
func (f *FaultEvent) EvelFaultTypeSet(fault_type string) {
	cFaultType := C.CString(fault_type)
	defer C.free(unsafe.Pointer(cFaultType))

	C.evel_fault_type_set(f.ptr, cFaultType)
}
