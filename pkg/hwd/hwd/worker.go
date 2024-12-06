package hwd

import (
	"bufio"
	"fmt"
	"go.bug.st/serial"
	"nokowebapi/nokocore"
	"nokowebapi/task"
	"os"
	"os/exec"
	"strings"
	"time"
)

func watch(stdin *os.File, stdout *os.File) {
	nokocore.KeepVoid(stdin, stdout)

	scanner := bufio.NewScanner(stdout)

	for {
		if scanner.Scan() {
			output := scanner.Text()
			exited := false

			switch {
			case strings.HasPrefix(output, "ERROR"):
				exited = true
				break

			case strings.HasPrefix(output, "FATAL ERROR"):
				exited = true
				break

			case strings.HasPrefix(output, "EXIT"):
				exited = true
				break

			case strings.HasPrefix(output, "Test"):
				if data, ok := strings.CutPrefix(output, "Test:"); ok {
					fmt.Printf("Test: %s\n", data)
				}

			case strings.HasPrefix(output, "Data"):
				if data, ok := strings.CutPrefix(output, "Data:"); ok {
					fmt.Printf("Data: %s\n", data)
				}
			}

			if exited {
				break
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			break
		}
	}
}

func NewWorker() {
	var process *exec.Cmd

	serialName := "COM4"
	args := []string{"/scan", "/serial", serialName}

	pipeIn, stdin := nokocore.Unwrap2(os.Pipe())
	stdout, pipeOut := nokocore.Unwrap2(os.Pipe())
	nokocore.KeepVoid(pipeIn, pipeOut, stdin, stdout)

	go watch(stdin, stdout)

	for {
		mode := &serial.Mode{
			BaudRate: 9600,
			Parity:   serial.NoParity,
			DataBits: 8,
			StopBits: serial.OneStopBit,
		}

		fmt.Printf("(%s) Connecting...\n", serialName)

		port, err := serial.Open(serialName, mode)
		nokocore.KeepVoid(port)

		busy := false
		found := true
		if err != nil {
			switch {
			case strings.EqualFold(err.Error(), "serial port not found"):
				fmt.Printf("(%s) Serial port not found.\n", serialName)
				found = false
				busy = false
				break

			case strings.EqualFold(err.Error(), "serial port busy"):
				fmt.Printf("(%s) Serial port busy.\n", serialName)
				found = true
				busy = true
				break
			}
		}

		if port != nil {
			if err = port.Close(); err != nil {
				fmt.Printf("(%s) Failed to close serial port: %s\n", serialName, err.Error())
			}

			fmt.Printf("(%s) Exited...\n", serialName)
		}

		if found && !busy {
			fmt.Printf("(%s) Busy...\n", serialName)
		}

		if found && !busy && process == nil {
			fmt.Println("(NokoHwd) Started...")
			process = nokocore.Unwrap(task.MakeProcess("./bin/NokoHwd.exe", args, nil, pipeIn, pipeOut, nil))
			nokocore.NoErr(process.Start())
			pid := nokocore.Unwrap(task.GetProcessState(process)).Pid()
			fmt.Println("(NokoHwd) Pid:", pid)
		}

		if !found && process != nil {
			if !nokocore.Unwrap(task.GetProcessState(process)).Exited() {
				if err = process.Process.Kill(); err != nil {
					fmt.Printf("(NokoHwd) Failed to kill process: %s\n", err.Error())
				}
			}

			process = nil
			fmt.Println("(NokoHwd) Stopped...")
		}

		time.Sleep(time.Second)
	}
}
