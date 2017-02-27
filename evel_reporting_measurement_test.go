package evel

import (
	"log"
	"testing"
)

func TestReportingMeasurementEvent(t *testing.T) {

	if err := NewEvelInitialize(api_fqdn, api_port, api_path,
		api_topic, api_secure, api_username, api_password,
		client_source_type, client_role, verbose); err != nil {
		t.Fatal(err)
	}

	defer EvelTerminate()

	r, err := NewEvelReport(1.1)
	if err != nil {
		t.Fatal(err)
	}

	r.EvelReportTypeSet("Perf reporting...")
	r.EvelReportFeatureUseAdd("FeatureA", 123.4)
	r.EvelReportFeatureUseAdd("FeatureB", 567.8)
	r.EvelReportCustomMeasurementAdd("Group1", "Name1", "Value1")
	r.EvelReportCustomMeasurementAdd("Group1", "Name2", "Value2")
	r.EvelReportCustomMeasurementAdd("Group2", "Name1", "Value1")

	if err := EvelPostEvent(r); err != nil {
		log.Fatal(err)
	}
}
