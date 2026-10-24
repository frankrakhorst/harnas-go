package tui

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

// RawMode enables raw terminal mode on fd and returns a restore function.
func RawMode(fd int) (restore func(), err error) {
	var old syscall.Termios
	if _, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&old)), 0, 0, 0); errno != 0 {
		return nil, fmt.Errorf("tcgets: %w", errno)
	}
	newt := old
	newt.Iflag &^= syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK | syscall.ISTRIP | syscall.INLCR | syscall.IGNCR | syscall.ICRNL | syscall.IXON
	newt.Oflag &^= syscall.OPOST
	newt.Lflag &^= syscall.ECHO | syscall.ECHONL | syscall.ICANON | syscall.ISIG | syscall.IEXTEN
	newt.Cflag &^= syscall.CSIZE | syscall.PARENB
	newt.Cflag |= syscall.CS8
	newt.Cc[syscall.VMIN] = 1
	newt.Cc[syscall.VTIME] = 0
	if _, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(&newt)), 0, 0, 0); errno != 0 {
		return nil, fmt.Errorf("tcsets: %w", errno)
	}
	return func() {
		syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(&old)), 0, 0, 0)
	}, nil
}

// StdinFD returns file descriptor for os.Stdin.
func StdinFD() int { return int(os.Stdin.Fd()) }