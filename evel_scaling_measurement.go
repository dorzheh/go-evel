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

type MeasurementEvent struct {
	ptr *C.EVENT_MEASUREMENT
}

/*****************************************************************************
 * Create a new Measurement event.
 *
 * @note    The mandatory fields on the Measurement must be supplied to this
 *          factory function and are immutable once set.  Optional fields have
 *          explicit setter functions, but again values may only be set once so
 *          that the Measurement has immutable properties.
 *
 * @param   concurrent_sessions
 * @param   configured_entities
 * @param   mean_request_latency
 * @param   measurement_interval
 * @param   memory_configured
 * @param   memory_used
 * @param   request_rate

 * @returns MeasurementEvent object containing pointer to C_EVENT_MEASUREMENT.
 *                            If the event is not used (i.e. posted) it must be
 *                            released using EvelFreeEvent.
 * @retval  error             Failed to create the event.
 *                            nil in case the event is created successfully.
 *****************************************************************************/
func NewEvelMeasurement(concurrent_sessions, configured_entries int,
	mean_request_latency, measurement_interval,
	memory_configured, memory_used float32,
	request_rate int) (MeasurementEvent, error) {

	cMeasPtr := C.evel_new_measurement(C.int(concurrent_sessions), C.int(configured_entries),
		C.double(mean_request_latency), C.double(measurement_interval),
		C.double(memory_configured), C.double(memory_used), C.int(request_rate))

	if cMeasPtr == nil {
		errorStr := "Error in creating a new measurement event."
		EvelError(errorStr)
		return MeasurementEvent{},  errors.New(errorStr)
	}
	return MeasurementEvent{ptr: cMeasPtr}, nil
}

/****************************************************************************
 * Set the Event Type property of the Measurement.
 *
 * @note  The property is treated as immutable: it is only valid to call
 *        the setter once.  However, we don't assert if the caller tries to
 *        overwrite, just ignoring the update instead.
 *
 * @param measurement_type  The Event Type to be set.
 *****************************************************************************/
func (m *MeasurementEvent) EvelMeasurementTypeSet(measurement_type string) {
	cMeasurementType := C.CString(measurement_type)
	defer C.free(unsafe.Pointer(cMeasurementType))

	C.evel_measurement_type_set(m.ptr, cMeasurementType)
}

/***************************************************************************
 * Add an additional CPU usage value name/value pair to the Measurement.
 *
 * @param name    string       CPU's name.
 * @param value   float32      CPU utilization.
 *****************************************************************************/
func (m *MeasurementEvent) EvelMeasurementCpuUseAdd(name string, value float32) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.evel_measurement_cpu_use_add(m.ptr, cName, C.double(value))
}

/*****************************************************************************
 * Add an additional File System usage value name/value pair to the
 * Measurement.
 *
 * @param vm_id                string   String with the file-system's UUID.
 * @param block_configured     float32  Block storage configured.
 * @param block_used           float32  Block storage in use.
 * @param block_iops           int      Block storage IOPS.
 * @param ephemeral_configured float32  Ephemeral storage configured.
 * @param ephemeral_used       float32  Ephemeral storage in use.
 * @param ephemeral_iops       int      Ephemeral storage IOPS.
 *****************************************************************************/
func (m *MeasurementEvent) EvelMeasurementFsysUseAdd(vm_id string, block_configured, block_used float32,
	block_iops int, ephemeral_configured, ephemeral_used float32,
	ephemeral_iops int) {
	cVmId := C.CString(vm_id)
	defer C.free(unsafe.Pointer(cVmId))

	C.evel_measurement_fsys_use_add(m.ptr, cVmId, C.double(block_configured), C.double(block_used),
		C.int(block_iops), C.double(ephemeral_configured),
		C.double(ephemeral_used), C.int(ephemeral_iops))
}

/*****************************************************************************
 * Add an additional Latency Distribution bucket to the Measurement.
 *
 * @param low_end  float32  Low end of the bucket's range.
 * @param high_end float32  High end of the bucket's range.
 * @param count    int      Count of events in this bucket.
 *****************************************************************************/
func (m *MeasurementEvent) EvelMeasurementLatencyAdd(low_end, hight_end float32, count int) {
	C.evel_measurement_latency_add(m.ptr, C.double(low_end), C.double(hight_end), C.int(count))
}

/****************************************************************************
 * Add an additional vNIC usage record Measurement.
 *
 * @param vnic_id               string vNIC's ID.
 * @param broadcast_packets_in  int    Broadcast packets received.
 * @param broadcast_packets_out int    Broadcast packets transmitted.
 * @param bytes_in              int    Total bytes received.
 * @param bytes_out             int    Total bytes transmitted.
 * @param multicast_packets_in  int    Multicast packets received.
 * @param multicast_packets_out int    Multicast packets transmitted.
 * @param unicast_packets_in    int    Unicast packets received.
 * @param unicast_packets_out   int    Unicast packets transmitted.
 *****************************************************************************/
func (m *MeasurementEvent) EvelMeasurementVnicUseAdd(vnic_id string, broadcast_packets_in, broadcast_packets_out,
	bytes_in, bytes_out, multicast_packets_in, multicast_packets_out, unicast_packets_in, unicast_packets_out int) {
	cVnicId := C.CString(vnic_id)
	defer C.free(unsafe.Pointer(cVnicId))

	C.evel_measurement_vnic_use_add(m.ptr, cVnicId, C.int(broadcast_packets_in), C.int(broadcast_packets_out),
		C.int(bytes_in), C.int(bytes_out), C.int(multicast_packets_in), C.int(multicast_packets_out), C.int(unicast_packets_in),
		C.int(unicast_packets_out))
}

/****************************************************************************
 * Add a Feature usage value name/value pair to the Measurement.
 *
 * @param feature     string  Feature's name.
 * @param utilization float32 Utilization of the feature.
 *****************************************************************************/
func (m *MeasurementEvent) EvelMeasurementFeatureUseAdd(feature string, utilization float32) {
	cFeature := C.CString(feature)
	defer C.free(unsafe.Pointer(cFeature))

	C.evel_measurement_feature_use_add(m.ptr, cFeature, C.double(utilization))
}

/****************************************************************************
 * Add a Additional Measurement value name/value pair to the Report.
 *
 * @param group string  Measurement group's name.
 * @param name  string  Measurement's name.
 * @param value string  Measurement's value.
 *****************************************************************************/
func (m *MeasurementEvent) EvelMeasurementCustomMeasurementAdd(group, name, value string) {
	cGroup := C.CString(group)
	defer C.free(unsafe.Pointer(cGroup))
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	C.evel_measurement_custom_measurement_add(m.ptr, cGroup, cName, cValue)
}

/****************************************************************************
 * Add a Codec usage value name/value pair to the Measurement.
 *
 * @param codec       string  Codec's name.
 * @param utilization int     Utilization of the feature.
 *****************************************************************************/
func (m *MeasurementEvent) EvelMeasurementCodecUseAdd(codec string, utilization int) {
	cCodec := C.CString(codec)
	defer C.free(unsafe.Pointer(cCodec))

	C.evel_measurement_codec_use_add(m.ptr, cCodec, C.int(utilization))
}

/****************************************************************************
 * Set the Aggregate CPU Use property of the Measurement.
 *
 * @param cpu_use  float32  The CPU use to set.
 *****************************************************************************/
func (m *MeasurementEvent) EvelMeasurementAggCpuUseSet(cpu_use float32) {
	C.evel_measurement_agg_cpu_use_set(m.ptr, C.double(cpu_use))
}

/****************************************************************************
 * Set the Media Ports in Use property of the Measurement.
 *
 * @note  The property is treated as immutable: it is only valid to call
 *        the setter once.  However, we don't assert if the caller tries to
 *        overwrite, just ignoring the update instead.
 *
 * @param media_ports_in_use  int The media port usage to set.
 *****************************************************************************/
func (m *MeasurementEvent) EvelMeasurementMediaPortUseSet(media_ports_in_use int) {
	C.evel_measurement_media_port_use_set(m.ptr, C.int(media_ports_in_use))
}

/****************************************************************************
 * Set the VNFC Scaling Metric property of the Measurement.
 *
 * @note  The property is treated as immutable: it is only valid to call
 *        the setter once.  However, we don't assert if the caller tries to
 *        overwrite, just ignoring the update instead.
 *
 * @param scaling_metric  float32 The scaling metric to set.
 *****************************************************************************/
func (m *MeasurementEvent) EvelMeasurementVnfcScalingMetricSet(scaling_metric float32) {
	C.evel_measurement_vnfc_scaling_metric_set(m.ptr, C.double(scaling_metric))
}
