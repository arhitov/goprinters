//go:build windows

package os

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"goprinters/types"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)

type printerObject struct {
	DriverName string `json:"DriverName"`
	Name       string `json:"Name"`
	PortName   string `json:"PortName"`
	Datatype   string `json:"Datatype"`
}

// PnpDevice структура для хранения информации об устройстве PnP
type pnpDevice struct {
	Name       string `json:"Name"`
	Status     string `json:"Status"`
	InstanceId string `json:"InstanceId"`
}

func GetOS() types.OS {
	return types.OSWindows
}

// GetPrinters получить список принтеров
func GetPrinters() ([]types.Printer, error) {
	// Создаем PowerShell команду
	psCommand := `Get-Printer | Where-Object {$_.PortName -like 'USB*'} | Select-Object DriverName, Name, PortName, Datatype | ConvertTo-Json`
	printerObjs, err := parseOutputCommand[printerObject](psCommand)
	if err != nil {
		return nil, err
	}

	var printers []types.Printer
	for _, printerObj := range printerObjs {
		printers = append(printers, types.Printer{
			SystemName:   printerObj.DriverName,
			FriendlyName: printerObj.Name,
			Port:         printerObj.PortName,
			Interface:    types.PrinterInterfaceUSB,
		})
	}
	return printers, nil
}

// CheckPrinterAvailability проверить что принтер доступен для печати
func CheckPrinterAvailability(printer types.Printer) error {
	// Создаем PowerShell команду
	psCommand := fmt.Sprintf(
		`Get-PnpDevice -Class Printer | Where-Object {$_.Name -eq '%s'} | Select-Object FriendlyName, Status, InstanceId | ConvertTo-Json`,
		printer.SystemName,
	)
	//psCommand := `Get-PnpDevice -Class Printer | Where-Object {$_.InstanceId -like '*USB*'} | Select-Object FriendlyName, Status, InstanceId | ConvertTo-Json`
	deviceObjs, err := parseOutputCommand[pnpDevice](psCommand)
	if err != nil {
		return err
	}

	if len(deviceObjs) != 1 {
		return errors.New("printer not found")
	}

	if strings.ToLower(deviceObjs[0].Status) != "ok" {
		return errors.New("printer not availability")
	}

	return nil
}

// PrintRaw отправляет текст на принтер
func PrintRaw(printer types.Printer, text string) error {
	var (
		winspool         = syscall.NewLazyDLL("winspool.drv")
		openPrinter      = winspool.NewProc("OpenPrinterW")
		closePrinter     = winspool.NewProc("ClosePrinter")
		startDocPrinter  = winspool.NewProc("StartDocPrinterW")
		endDocPrinter    = winspool.NewProc("EndDocPrinter")
		startPagePrinter = winspool.NewProc("StartPagePrinter")
		endPagePrinter   = winspool.NewProc("EndPagePrinter")
		writePrinter     = winspool.NewProc("WritePrinter")
	)

	// Конвертируем строки в UTF-16
	pPrinterName, err := syscall.UTF16PtrFromString(printer.SystemName)
	if err != nil {
		return err
	}

	var hPrinter syscall.Handle
	ret, _, err := openPrinter.Call(
		uintptr(unsafe.Pointer(pPrinterName)),
		uintptr(unsafe.Pointer(&hPrinter)),
		0,
	)
	if ret == 0 {
		return fmt.Errorf("OpenPrinter failed: %v", err)
	}
	defer closePrinter.Call(uintptr(hPrinter))

	// DOC_INFO_1
	docName := "ZPL Print"
	pDocName, _ := syscall.UTF16PtrFromString(docName)
	pDataType, _ := syscall.UTF16PtrFromString("RAW")

	docInfo := struct {
		pDocName    uintptr
		pOutputFile uintptr
		pDataType   uintptr
	}{
		uintptr(unsafe.Pointer(pDocName)),
		0,
		uintptr(unsafe.Pointer(pDataType)),
	}

	ret, _, err = startDocPrinter.Call(
		uintptr(hPrinter),
		1,
		uintptr(unsafe.Pointer(&docInfo)),
	)
	if ret == 0 {
		return fmt.Errorf("StartDocPrinter failed: %v", err)
	}
	defer endDocPrinter.Call(uintptr(hPrinter))

	ret, _, err = startPagePrinter.Call(uintptr(hPrinter))
	if ret == 0 {
		return fmt.Errorf("StartPagePrinter failed: %v", err)
	}
	defer endPagePrinter.Call(uintptr(hPrinter))

	// Пишем данные
	bytes := []byte(text)
	var bytesWritten uint32

	ret, _, err = writePrinter.Call(
		uintptr(hPrinter),
		uintptr(unsafe.Pointer(&bytes[0])),
		uintptr(len(bytes)),
		uintptr(unsafe.Pointer(&bytesWritten)),
	)
	if ret == 0 {
		return fmt.Errorf("WritePrinter failed: %v", err)
	}

	return nil
}

func callCommand(command string) (string, error) {
	cmd := exec.Command("powershell", "-Command", command)

	// Захватываем вывод
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("ошибка выполнения PowerShell: %w", err)
	}

	output := out.String()

	// Проверяем, не пустой ли вывод
	if strings.TrimSpace(output) == "" {
		return "", nil
	}

	return output, nil
}

func parseOutputCommand[T any](command string) ([]T, error) {
	output, err := callCommand(command)
	if err != nil {
		return nil, err
	}
	var objs []T
	err = json.Unmarshal([]byte(output), &objs)
	if err != nil {
		// Пробуем распарсить как одиночный объект
		var obj T
		if err2 := json.Unmarshal([]byte(output), &obj); err2 == nil {
			objs = []T{obj}
		} else {
			return nil, fmt.Errorf("ошибка парсинга JSON: %w\nВывод: %s", err, output)
		}
	}
	return objs, nil
}
