package types

type Printer struct {
	SystemName   string
	FriendlyName string
	VendorID     string
	ProductID    string
	Port         string
	Interface    PrinterInterface
}
