package entities

import "github.com/arhitov/goprinters/types"

type Printer struct {
	SystemName   string                 `json:"system_name"`
	FriendlyName string                 `json:"friendly_name"`
	Location     string                 `json:"location"`
	Interface    types.PrinterInterface `json:"interface"`
}
