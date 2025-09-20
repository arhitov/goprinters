package goprinters

import (
	"errors"
	"github.com/arhitov/goprinters/entities"
	"github.com/arhitov/goprinters/os"
	"github.com/arhitov/goprinters/protocols/telnet"
	"github.com/arhitov/goprinters/types"
	"runtime"
)

func GetOS() types.OS {
	switch runtime.GOOS {
	case "windows":
		return types.OSWindows
	default:
		return types.OSUnknown
	}
}

// GetSystemPrinters получить список принтеров зарегистрированных в системе
func GetSystemPrinters() ([]entities.Printer, error) {
	if GetOS() == types.OSUnknown {
		return []entities.Printer{}, nil
	}
	return os.GetSystemPrinters()
}

// CheckPrinterAvailability проверить что принтер доступен для печати
func CheckPrinterAvailability(printer entities.Printer) error {
	if GetOS() == types.OSUnknown {
		return errors.New("no printer available")
	}

	if printer.Interface == types.PrinterInterfaceTelnet {
		return telnet.NewTelnet(printer.Location, 5).Ping()
	} else {
		return os.CheckPrinterAvailability(printer)
	}
}

// PrintRaw отправляет текст на принтер
func PrintRaw(printer entities.Printer, text string) error {
	if GetOS() == types.OSUnknown {
		return errors.New("no printer available")
	}
	if printer.Interface == types.PrinterInterfaceTelnet {
		return telnet.NewTelnet(printer.Location, 5).Write(text)
	} else {
		return os.PrintRaw(printer, text)
	}
}
