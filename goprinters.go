package goprinters

import (
	"github.com/arhitov/goprinters/os"
	"github.com/arhitov/goprinters/types"
)

// GetPrinters получить список принтеров
func GetPrinters() ([]types.Printer, error) {
	return os.GetPrinters()
}

// CheckPrinterAvailability проверить что принтер доступен для печати
func CheckPrinterAvailability(printer types.Printer) error {
	return os.CheckPrinterAvailability(printer)
}

// PrintRaw отправляет текст на принтер
func PrintRaw(printer types.Printer, text string) error {
	return os.PrintRaw(printer, text)
}
