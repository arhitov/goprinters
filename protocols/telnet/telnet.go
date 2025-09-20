package telnet

import (
	//"app/pkg/gologger"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

type Telnet struct {
	address string
	timeout int // Second
}

func NewTelnet(address string, timeout int) *Telnet {
	return &Telnet{
		address: address,
		timeout: timeout,
	}
}

func (t *Telnet) Write(printData string) error {
	conn, err := t.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Отправляем данные для печати
	_, err = conn.Write([]byte(printData))
	if err != nil {
		//gologger.Error(err).Message("failed to send print data").Done()
		return fmt.Errorf("failed to send print data: %w", err)
	}

	//gologger.Debugf("Print command sent successfully").Done()
	return nil
}

func (t *Telnet) Ping() error {
	conn, err := t.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Если соединение установлено, то значит доступ есть
	return nil
}

func (t *Telnet) connect() (net.Conn, error) {
	address := t.address
	// Если нет в адресе порта, добавляем его
	if !strings.Contains(t.address, ":") {
		address = net.JoinHostPort(t.address, "23")
	}
	host, port, err := net.SplitHostPort(address)
	//gologger.Debugf("checkTelnetPrinter %s => %s:%s", t.address, host, port).Error(err).Done()
	if err != nil {
		return nil, errors.New("invalid address format")
	}

	var conn net.Conn
	if t.timeout > 0 {
		timeout := time.Duration(t.timeout) * time.Second
		conn, err = net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			return nil, err
		}

		// Устанавливаем таймаут на операции чтения/записи
		_ = conn.SetDeadline(time.Now().Add(1 * time.Second))
	} else {
		conn, err = net.Dial("tcp", net.JoinHostPort(host, port))
		if err != nil {
			return nil, err
		}
	}

	return conn, nil
}
