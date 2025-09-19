package types

type Printer struct {
	SystemName   string
	FriendlyName string
	Location     string
	Interface    PrinterInterface
}
