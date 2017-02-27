package evel

import (
	"testing"
)

func TestHeartbeatEvent(t *testing.T) {

	if err := NewEvelInitialize(api_fqdn, api_port, api_path,
		api_topic, api_secure, api_username, api_password,
		client_source_type, client_role, verbose); err != nil {
		t.Fatal(err)
	}

	h, err := NewEvelHeartbeat()
	if err != nil {
		t.Fatal(err)
	}
	if err := EvelPostEvent(h); err != nil {
		EvelFreeEvent(h.ptr)
		t.Fatal(err)
	}

	println("The heartbeat event is processed sucessfully.")
	if err := EvelTerminate();err != nil {
		t.Fatal(err)
	}
}
