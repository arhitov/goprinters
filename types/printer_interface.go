package types

import "github.com/arhitov/goenum"

type PrinterInterface = goenum.StringNamedMeta[printerInterfaceMap, printerInterfaceMeta]

const (
	PrinterInterfaceUSB    PrinterInterface = "usb"
	PrinterInterfaceTelnet PrinterInterface = "telnet"
)

type printerInterfaceMap struct{}

type printerInterfaceMeta struct{}

func (l printerInterfaceMap) ValueMap() map[string]any {
	return map[string]any{
		string(PrinterInterfaceUSB):    PrinterInterfaceUSB,
		string(PrinterInterfaceTelnet): PrinterInterfaceTelnet,
	}
}

func (l printerInterfaceMap) NameMap() map[string]string {
	return map[string]string{
		string(PrinterInterfaceUSB):    "USB",
		string(PrinterInterfaceTelnet): "Telnet",
	}
}

func (l printerInterfaceMap) MetaMap() map[string]printerInterfaceMeta {
	return map[string]printerInterfaceMeta{
		string(PrinterInterfaceUSB):    {},
		string(PrinterInterfaceTelnet): {},
	}
}
