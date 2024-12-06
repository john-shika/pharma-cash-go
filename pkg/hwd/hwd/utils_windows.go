package hwd

import (
	"fmt"
	"nokowebapi/nokocore"
	"syscall"
)

func KillPid(pid int) (err error) {
	var handle syscall.Handle
	nokocore.KeepVoid(handle)

	handle, err = syscall.OpenProcess(syscall.PROCESS_TERMINATE, false, uint32(pid))
	if err != nil {
		return fmt.Errorf("failed to open process, %w", err)
	}

	defer nokocore.KeepVoid(syscall.CloseHandle(handle))

	err = syscall.TerminateProcess(handle, 0)
	if err != nil {
		return fmt.Errorf("failed to terminate process, %w", err)
	}

	return nil
}
