//go:build !windows

package os

import (
	"errors"
	"goprinters/types"
)

func GetOS() types.OS {
	return types.OSUnknown
}

// GetPrinters получить список принтеров
func GetPrinters() ([]types.Printer, error) {
	return []types.Printer{}, nil
}

// CheckPrinterAvailability проверить что принтер доступен для печати
func CheckPrinterAvailability(printer types.Printer) error {
	return errors.New("no printer available")
}

// PrintRaw отправляет текст на принтер
func PrintRaw(printer types.Printer, text string) error {
	return errors.New("no printer available")
}
