package beta

import (
	"errors"
	"fmt"
	"go.uber.org/zap/buffer"
	"io"
	"nokowebapi/nokocore"
	"nokowebapi/task"
	"os"
	"strconv"
)

type PrinterCommand struct {
	buffer *buffer.Buffer
	name   string
}

func NewPrinter(name string) io.ReadWriteCloser {
	return &PrinterCommand{
		buffer: &buffer.Buffer{},
		name:   name,
	}
}

func (p *PrinterCommand) Read(data []byte) (int, error) {
	panic("not implemented")
	return 0, nil
}

func (p *PrinterCommand) Write(data []byte) (int, error) {
	return p.buffer.Write(data)
}

func createTempFile(buff *buffer.Buffer, pattern string) (*os.File, error) {
	tempDir := os.TempDir()
	tempFile, err := os.CreateTemp(tempDir, pattern)
	if err != nil {
		return nil, errors.New("error creating temporary file")
	}

	temp := buff.Bytes()
	n, err := tempFile.Write(temp)
	nokocore.KeepVoid(n)

	if err != nil {
		return nil, errors.New("error writing to temporary file")
	}

	return tempFile, nil
}

func (p *PrinterCommand) Close() error {
	path := "cmd.exe"
	computerName := os.Getenv("COMPUTERNAME")
	printerName := p.name

	tempFile, err := createTempFile(p.buffer, "escpos-*")
	if err != nil {
		return err
	}

	defer nokocore.NoErr(tempFile.Close())
	//defer nokocore.NoErr(os.Remove(tempFile.Name()))

	ns := fmt.Sprintf("\\\\%s\\%s", computerName, strconv.Quote(printerName))

	scriptExec := fmt.Sprintf("@echo off\nprint /D:%s %s", ns, strconv.Quote(tempFile.Name()))

	buffExec := &buffer.Buffer{}
	nokocore.Unwrap(buffExec.WriteString(scriptExec))

	tempExecFile, err := createTempFile(buffExec, "escpos-*.bat")
	if err != nil {
		return err
	}

	defer nokocore.NoErr(tempExecFile.Close())
	//defer nokocore.NoErr(os.Remove(tempExecFile.Name()))

	args := []string{"/C", tempExecFile.Name()}
	var environ []string

	process, err := task.MakeProcess(path, args, environ, nil, os.Stdout, os.Stderr)
	if err != nil {
		return err
	}

	return process.Run()
}
