package evel

import (
	"log"
	"testing"
)

func TestScalingMeasurementEvent(t *testing.T) {

	if err := NewEvelInitialize(api_fqdn, api_port, api_path,
		api_topic, api_secure, api_username, api_password,
		client_source_type, client_role, verbose); err != nil {
		t.Fatal(err)
	}

	defer EvelTerminate()

	m, err := NewEvelMeasurement(1, 2, 3.3, 4.4, 5.5, 6.6, 7)
	if err != nil {
		t.Fatal(err)
	}

	m.EvelMeasurementTypeSet("Perf management...")
	m.EvelMeasurementAggCpuUseSet(8.8)
	m.EvelMeasurementCpuUseAdd("cpu1", 11.11)
	m.EvelMeasurementCpuUseAdd("cpu2", 22.22)

	m.EvelMeasurementFsysUseAdd("00-11-22", 100.11, 100.22, 33, 200.11, 200.22, 44)
	m.EvelMeasurementFsysUseAdd("33-44-55", 300.11, 300.22, 55, 400.11, 400.22, 66)

	m.EvelMeasurementLatencyAdd(0.0, 10.0, 20)
	m.EvelMeasurementLatencyAdd(10.0, 20.0, 30)

	m.EvelMeasurementVnicUseAdd("eth0", 1, 2, 3, 4, 5, 6, 7, 8)
	m.EvelMeasurementVnicUseAdd("eth1", 11, 12, 13, 14, 15, 16, 17, 18)

	m.EvelMeasurementFeatureUseAdd("FeatureA", 123.4)
	m.EvelMeasurementFeatureUseAdd("FeatureB", 567.8)

	m.EvelMeasurementCodecUseAdd("G711a", 91)
	m.EvelMeasurementCodecUseAdd("G729ab", 92)

	m.EvelMeasurementMediaPortUseSet(1234)

	m.EvelMeasurementVnfcScalingMetricSet(1234.5678)

	m.EvelMeasurementCustomMeasurementAdd("Group1", "Name1", "Value1")
	m.EvelMeasurementCustomMeasurementAdd("Group1", "Name2", "Value2")
	m.EvelMeasurementCustomMeasurementAdd("Group2", "Name1", "Value1")
	if err := EvelPostEvent(m); err != nil {
		log.Fatal(err)
	}

}
