package evel

import (
	"testing"
)

func TestFaultEvent(t *testing.T) {

	if err := NewEvelInitialize(api_fqdn, api_port, api_path,
		api_topic, api_secure, api_username, api_password,
		client_source_type, client_role, verbose); err != nil {
		t.Fatal(err)
	}

	defer EvelTerminate()

	condition := "My alarm condition"
	specific_problem := "Port error"
	priority := EVEL_PRIORITY_HIGH
	severity := EVEL_SEVERITY_CRITICAL
	f, err := NewEvelFault(condition, specific_problem, priority, severity)
	if err != nil {
		t.Fatal(err)
	}

	fault_event_type := "A failure occured"
	f.EvelFaultTypeSet(fault_event_type)

	alarm_interface := "My Interface Card"
	f.EvelFaultInterfaceSet(alarm_interface)

	key1 := "key1"
	value1 := "value1"
	f.EvelFaultAddlInfoAdd(key1, value1)

	key2 := "key2"
	value2 := "value2"
	f.EvelFaultAddlInfoAdd(key2, value2)
	if err := EvelPostEvent(f); err != nil {
		EvelFreeEvent(f.ptr)
		t.Fatal(err)
	}

	println("The fault event is processed sucessfully.")
}
