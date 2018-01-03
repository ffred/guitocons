package guitocons

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

const ATTACH_PARENT_PROCESS = ^uint32(0) // (DWORD)-1

var (
	modkernel32       = syscall.NewLazyDLL("kernel32.dll")
	procAttachConsole = modkernel32.NewProc("AttachConsole")
)

func attachConsole(dwParentProcess uint32) (ok bool, lasterr error) {
	r1, _, lasterr := syscall.Syscall(procAttachConsole.Addr(), 1, uintptr(dwParentProcess), 0, 0)
	ok = bool(r1 != 0)
	return
}

// Guitocons attach the GUI process to the parent console, so we can send "logs" & "prints" if needed
func Guitocons() (err error) {
	ok, lasterr := attachConsole(ATTACH_PARENT_PROCESS)
	if ok {
		hout, err1 := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
		if err1 != nil {
			return fmt.Errorf("stdout connection error : %v", err1)
		}
		herr, err2 := syscall.GetStdHandle(syscall.STD_ERROR_HANDLE)
		if err2 != nil {
			return fmt.Errorf("stderr connection error : %v", err2)
		}
		os.Stdout = os.NewFile(uintptr(hout), "/dev/stdout")
		os.Stderr = os.NewFile(uintptr(herr), "/dev/stderr")
		log.SetOutput(os.Stderr)
		return
	}
	return fmt.Errorf("attachconsole error : %v", lasterr)
}
