package main

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/sys/windows"
	"io"
	"nokowebapi/nokocore"
	"nokowebapi/task"
	"syscall"
	"time"
	"unsafe"
)

type HandleIO struct {
	Handle *windows.Handle
	closed bool
}

func NewHandleIO(h *windows.Handle) *HandleIO {
	return &HandleIO{
		Handle: h,
	}
}

func (h *HandleIO) Read(p []byte) (int, error) {
	if h.closed {
		return 0, io.EOF
	}

	return windows.Read(*h.Handle, p)
}

func (h *HandleIO) Write(p []byte) (int, error) {
	if h.closed {
		return 0, errors.New("file already closed")
	}

	return windows.Write(*h.Handle, p)
}

func (h *HandleIO) Close() error {
	h.closed = true
	return windows.CloseHandle(*h.Handle)
}

type ConPty struct {
	Handle      *windows.Handle
	ProcessInfo *windows.ProcessInformation
	PtyIn       *HandleIO
	PtyOut      *HandleIO
	CmdIn       *HandleIO
	CmdOut      *HandleIO
}

func (c *ConPty) Read(p []byte) (int, error) {
	return c.CmdOut.Read(p)
}

func (c *ConPty) Write(p []byte) (int, error) {
	return c.CmdIn.Write(p)
}

func (c *ConPty) Close() error {
	var errorStack = []error{
		c.PtyIn.Close(),
		c.PtyOut.Close(),
		c.CmdIn.Close(),
		c.CmdOut.Close(),
	}

	for i, err := range errorStack {
		nokocore.KeepVoid(i)

		if err != nil {
			return err
		}
	}

	return nil
}

func GetProcThreadAttributeListContainer(hPCon *windows.Handle) (*windows.ProcThreadAttributeListContainer, error) {
	var err error

	procThreadAttributeList := &windows.ProcThreadAttributeListContainer{}
	if procThreadAttributeList, err = windows.NewProcThreadAttributeList(1); err != nil {
		return nil, err
	}

	err = procThreadAttributeList.Update(windows.PROC_THREAD_ATTRIBUTE_PSEUDOCONSOLE, unsafe.Pointer(hPCon), unsafe.Sizeof(hPCon))
	return procThreadAttributeList, err
}

func CloseHandles(handles ...windows.Handle) error {
	var err error
	for i, handle := range handles {
		nokocore.KeepVoid(i)

		if handle != windows.InvalidHandle {
			if err != nil {
				nokocore.KeepVoid(windows.CloseHandle(handle))
				continue
			}

			err = windows.CloseHandle(handle)
		}
	}

	return err
}

func main() {
	var n int
	var err error
	var hPCon windows.Handle
	var procThreadAttributeListContainer *windows.ProcThreadAttributeListContainer

	nokocore.KeepVoid(n, err, hPCon, procThreadAttributeListContainer)

	size := windows.Coord{
		X: 80,
		Y: 40,
	}

	nokocore.KeepVoid(size)

	var ptyIn, ptyOut, cmdIn, cmdOut windows.Handle

	conPty := &ConPty{
		Handle: &hPCon,
		PtyIn:  NewHandleIO(&ptyIn),
		PtyOut: NewHandleIO(&ptyOut),
		CmdIn:  NewHandleIO(&cmdIn),
		CmdOut: NewHandleIO(&cmdOut),
	}

	nokocore.KeepVoid(conPty)

	if procThreadAttributeListContainer, err = GetProcThreadAttributeListContainer(&hPCon); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", procThreadAttributeListContainer)

	if err = windows.CreatePipe(&ptyIn, &cmdIn, nil, 0); err != nil {
		panic(err)
	}

	if err = windows.CreatePipe(&cmdOut, &ptyOut, nil, 0); err != nil {
		nokocore.NoErr(CloseHandles(ptyIn, cmdIn))
		panic(err)
	}

	if err = windows.CreatePseudoConsole(size, ptyIn, ptyOut, 0, &hPCon); err != nil {
		panic(err)
	}

	if err = windows.ResizePseudoConsole(hPCon, size); err != nil {
		panic(err)
	}

	defer windows.ClosePseudoConsole(hPCon)

	startupInfo := &windows.StartupInfo{}

	if err = windows.GetStartupInfo(startupInfo); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", startupInfo)
	fmt.Println(windows.UTF16PtrToString(startupInfo.Title))

	process := nokocore.Unwrap(task.MakeProcess("vlc.exe", nil, nil, conPty.PtyIn, conPty.PtyOut, nil))
	nokocore.NoErr(process.Start())

	fmt.Println("Process started")

	go func() {
		if n, err = conPty.Write([]byte("dir\n")); err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)

		fmt.Println("Process exited")

		if n, err = conPty.Write([]byte("exit\n")); err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)

		nokocore.NoErr(process.Process.Signal(syscall.SIGKILL))
		nokocore.NoErr(conPty.Close())
	}()

	scanner := bufio.NewScanner(conPty)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println(input)
		if err = scanner.Err(); err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
	}

	nokocore.NoErr(process.Wait())
}
