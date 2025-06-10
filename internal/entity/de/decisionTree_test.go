package de

import (
	"testing"
)

// ConfigRecord represents a single configuration record
type ConfigRecord struct {
	LenderID        int    `csv:"lender_id"`
	ProductCode     string `csv:"product_code,omitempty"`
	LeadSource      string `csv:"lead_source,omitempty"`
	FlowType        string `csv:"flow_type,omitempty"`
	UIVersion       string `csv:"ui_version,omitempty"`
	Status          string `csv:"status,omitempty"`
	OnlineOrOffline string `csv:"online_or_offline,omitempty"`
}

// TestDecisionTree_FindNodeByName Find a node by name
func TestDecisionTree_FindNodeByName(t *testing.T) {
	// TODO: Get lender_id
	// TODO: Get product_code
	// TODO: Get lead_source
	// TODO: Get flow_type
	// TODO: Get ui_flow
	// TODO: Get ui_flow_settings and ui_version (if have)
	// TODO: Get ui_version
	// TODO: Get active status

	/**
	table 1:
	lender_id, product_code, lead_source, flow_type, ui_version, status, online or offline
	9031, tpb_01, organic, normal, v1.0, active,TBD
	9031, tpb_01, evo, normal, v1.0, active,TBD
	9031, tpb_01, esms, normal, v1.0, active,TBD
	9031, tpb_01, evoapp, normal, v1.0, active,TBD
	9031, tpb_01, viettel_store_online, normal, v1.0, active,TBD
	*/

}
