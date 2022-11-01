package domain

import (
	"github.com/zcalusic/sysinfo"
)

// SystemReport provides a comprehensive view of the system's hardware and software configuration.
type SystemReport struct {
	Meta    interface{}   `json:"sysinfo"`
	Node    interface{}   `json:"node"`
	OS      interface{}   `json:"os"`
	Kernel  interface{}   `json:"kernel"`
	Product interface{}   `json:"product"`
	Board   interface{}   `json:"board"`
	Chassis interface{}   `json:"chassis"`
	BIOS    interface{}   `json:"bios"`
	CPU     interface{}   `json:"cpu"`
	Memory  interface{}   `json:"memory"`
	Storage []interface{} `json:"storage,omitempty"`
	Network []interface{} `json:"network,omitempty"`
}

func NewSystemReport(si sysinfo.SysInfo) *SystemReport {
	storage := make([]interface{}, len(si.Storage))
	for i, s := range si.Storage {
		storage[i] = s
	}

	network := make([]interface{}, len(si.Network))
	for i, n := range si.Network {
		network[i] = n
	}
	return &SystemReport{
		Meta:    si.Meta,
		Node:    si.Node,
		OS:      si.OS,
		Kernel:  si.Kernel,
		Product: si.Product,
		Board:   si.Board,
		Chassis: si.Chassis,
		BIOS:    si.BIOS,
		CPU:     si.CPU,
		Memory:  si.Memory,
		Storage: storage,
		Network: network,
	}
}
