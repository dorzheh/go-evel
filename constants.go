package evel

/*
#cgo CFLAGS: -I/opt/evel-library/code/evel_library
#include <evel.h>
#include <stdlib.h>
*/
import "C"

// EVEL_ERR_CODES
//
// Error codes for EVEL low level interface
const (
	EVEL_SUCCESS                = C.EVEL_SUCCESS                // The operation was successful.
	EVEL_ERR_GEN_FAIL           = C.EVEL_ERR_GEN_FAIL           // Non-specific failure.
	EVEL_CURL_LIBRARY_FAIL      = C.EVEL_CURL_LIBRARY_FAIL      // A cURL library operation failed.
	EVEL_PTHREAD_LIBRARY_FAIL   = C.EVEL_PTHREAD_LIBRARY_FAIL   // A Posix threads operation failed.
	EVEL_OUT_OF_MEMORY          = C.EVEL_OUT_OF_MEMORY          // A memory allocation failure occurred.
	EVEL_EVENT_BUFFER_FULL      = C.EVEL_EVENT_BUFFER_FULL      // Too many events in the ring-buffer.
	EVEL_EVENT_HANDLER_INACTIVE = C.EVEL_EVENT_HANDLER_INACTIVE // Attempt to raise event when inactive.
	EVEL_NO_METADATA            = C.EVEL_NO_METADATA            // Failed to retrieve OpenStack metadata.
	EVEL_BAD_METADATA           = C.EVEL_BAD_METADATA           // OpenStack metadata invalid format.
	EVEL_BAD_JSON_FORMAT        = C.EVEL_BAD_JSON_FORMAT        // JSON failed to parse correctly.
	EVEL_JSON_KEY_NOT_FOUND     = C.EVEL_JSON_KEY_NOT_FOUND     // Failed to find the specified JSON key.
	EVEL_MAX_ERROR_CODES        = C.EVEL_MAX_ERROR_CODES        // Maximum number of valid error codes.
)

// EVEL_EVENT_DOMAINS

// Event domains for the various events we support.
// JSON equivalent field: domain
type EvelEventDomains int

const (
	EVEL_DOMAIN_INTERNAL    EvelEventDomains = iota // C.EVEL_DOMAIN_INTERNAL    - Internal event, not for external routing.
	EVEL_DOMAIN_HEARTBEAT                           // C.EVEL_DOMAIN_HEARTBEAT   - A Heartbeat event (event header only).
	EVEL_DOMAIN_FAULT                               // C.EVEL_DOMAIN_FAULT       - A Fault event.
	EVEL_DOMAIN_MEASUREMENT                         // C.EVEL_DOMAIN_MEASUREMENT - A Measurement for VF Scaling event.
	EVEL_DOMAIN_REPORT                              // C.EVEL_DOMAIN_REPORT      - A Measurement for VF Reporting event.
	EVEL_MAX_DOMAINS                                // C.EVEL_MAX_DOMAINS        - Maximum number of recognized Event types.
)

//EVEL_EVENT_PRIORITIES
//
//Event priorities.
//JSON equivalent field: priority
type EvelEventPriorities int

const (
	EVEL_PRIORITY_HIGH   EvelEventPriorities = iota // C.EVEL_PRIORITY_HIGH
	EVEL_PRIORITY_MEDIUM                            // C.EVEL_PRIORITY_MEDIUM
	EVEL_PRIORITY_NORMAL                            // C.EVEL_PRIORITY_NORMAL
	EVEL_PRIORITY_LOW                               // C.EVEL_PRIORITY_LOW
	EVEL_MAX_PRIORITIES                             // C.EVEL_MAX_PRIORITIES
)

//EVEL_FAULT_SEVERITIES
//
//  Fault severities.
//  JSON equivalent field: eventSeverity
type EvelFaultSeverities int

const (
	EVEL_SEVERITY_CRITICAL EvelFaultSeverities = iota // C.EVEL_SEVERITY_CRITICAL
	EVEL_SEVERITY_MAJOR                               // C.EVEL_SEVERITY_MAJOR
	EVEL_SEVERITY_MINOR                               // C.EVEL_SEVERITY_MINOR
	EVEL_SEVERITY_WARNING                             // C.EVEL_SEVERITY_WARNING
	EVEL_SEVERITY_NORMAL                              // C.EVEL_SEVERITY_NORMAL
	EVEL_MAX_SEVERITIES                               // C.EVEL_MAX_SEVERITIES
)

// EVEL_SOURCE_TYPES

// Fault source types.
// JSON equivalent field: eventSourceType
type EvelSourceType int

const (
	EVEL_SOURCE_OTHER           EvelSourceType = iota // C.EVEL_SOURCE_OTHER
	EVEL_SOURCE_ROUTER                                // C.EVEL_SOURCE_ROUTER
	EVEL_SOURCE_SWITCH                                // C.EVEL_SOURCE_SWITCH
	EVEL_SOURCE_HOST                                  // C.EVEL_SOURCE_HOST
	EVEL_SOURCE_CARD                                  // C.EVEL_SOURCE_CARD
	EVEL_SOURCE_PORT                                  // C.EVEL_SOURCE_PORT
	EVEL_SOURCE_SLOT_THRESHOLD                        // C.EVEL_SOURCE_SLOT_THRESHOLD
	EVEL_SOURCE_PORT_THRESHOLD                        // C.EVEL_SOURCE_PORT_THRESHOLD
	EVEL_SOURCE_VIRTUAL_MACHINE                       // C.EVEL_SOURCE_VIRTUAL_MACHINE
	EVEL_MAX_SOURCE_TYPES                             // C.EVEL_MAX_SOURCE_TYPES
)
