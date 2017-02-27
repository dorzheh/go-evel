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

type ReportEvent struct {
	ptr *C.EVENT_REPORT
}

/**********************************************************************************
 * Create a new Report event.
 *
 * @note    The mandatory fields on the Report must be supplied to this
 *          factory function and are immutable once set.  Optional fields have
 *          explicit setter functions, but again values may only be set once so
 *          that the Report has immutable properties.
 *
 * @param   measurement_interval

 * @returns ReportEvent structure containing a pointer to C.EVENT_REPORT.
 *                      If the event isnot used (i.e. posted) it must be released
 *                      using EvelFreeEvent.
 * @retval  error       Failed to create the event and nil in case the event is created
 *                      successfully
 **********************************************************************************/
func NewEvelReport(measurement_interval float32) (ReportEvent, error) {
	reportPtr := C.evel_new_report(C.double(measurement_interval))
	if reportPtr == nil {
		errorStr := "Error in creating a new report event."
		EvelError(errorStr)
		return ReportEvent{},  errors.New(errorStr)
	}
	return ReportEvent{ptr: reportPtr}, nil
}

/*****************************************************************************
 * Set the Event Type property of the Report.
 *
 * @note  The property is treated as immutable: it is only valid to call
 *        the setter once.  However, we don't assert if the caller tries to
 *        overwrite, just ignoring the update instead.
 *
 * @param event_type  string The Event Type to be set. The caller does not need
 *                           to preserve the value once the function returns
 *****************************************************************************/
func (r *ReportEvent) EvelReportTypeSet(event_type string) {
	cEventType := C.CString(event_type)
	defer C.free(unsafe.Pointer(cEventType))

	C.evel_report_type_set(r.ptr, cEventType)
}

/****************************************************************************
 * Add a Feature usage value name/value pair to the Report.
 *
 * @param feature      string   Feature's name.
 * @param utilization  float32  Utilization of the feature.
 *****************************************************************************/
func (r *ReportEvent) EvelReportFeatureUseAdd(feature string, utilization float32) {
	cFeature := C.CString(feature)
	defer C.free(unsafe.Pointer(cFeature))

	C.evel_report_feature_use_add(r.ptr, cFeature, C.double(utilization))
}

/****************************************************************************
 * Add a Additional Measurement value name/value pair to the Report.
 *
 * The name is null delimited ASCII string.  The library takes
 * a copy so the caller does not have to preserve values after the function
 * returns.
 *
 * @param group  string  Measurement group's name.
 * @param name   string  Measurement's name.
 * @param value  string  Measurement's value.
 *****************************************************************************/
func (r *ReportEvent) EvelReportCustomMeasurementAdd(group, name, value string) {
	cGroup := C.CString(group)
	defer C.free(unsafe.Pointer(cGroup))
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.evel_report_custom_measurement_add(r.ptr, cGroup, cName, cValue)
}
