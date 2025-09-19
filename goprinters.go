package goprinters

import (
	"errors"
	"github.com/arhitov/goprinters/os"
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

// GetPrinters получить список принтеров
func GetPrinters() ([]types.Printer, error) {
	if GetOS() == types.OSUnknown {
		return []types.Printer{}, nil
	}
	return os.GetPrinters()
}

// CheckPrinterAvailability проверить что принтер доступен для печати
func CheckPrinterAvailability(printer types.Printer) error {
	if GetOS() == types.OSUnknown {
		return errors.New("no printer available")
	}
	return os.CheckPrinterAvailability(printer)
}

// PrintRaw отправляет текст на принтер
func PrintRaw(printer types.Printer, text string) error {
	if GetOS() == types.OSUnknown {
		return errors.New("no printer available")
	}
	return os.PrintRaw(printer, text)
}
