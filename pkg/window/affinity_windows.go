//go:build windows

package window

import (
	"syscall"
)

var (
	user32                   = syscall.NewLazyDLL(`user32.dll`)
	setWindowDisplayAffinity = user32.NewProc(`SetWindowDisplayAffinity`)
)

type DisplayAffinity int

const (
	None DisplayAffinity = iota
	Monitor
	ExcludeFromCapture
)

func (d DisplayAffinity) Uintptr() uintptr {
	return [...]uintptr{
		0x00000000,
		0x00000001,
		0x00000011,
	}[d]
}

func (d DisplayAffinity) String() string {
	return [...]string{
		`none`,
		`monitor`,
		`exclude from capture`,
	}[d]
}

func (d DisplayAffinity) SetDisplayAffinity(hwnd Hwnd) {
	_, _, _ = setWindowDisplayAffinity.Call(uintptr(hwnd), d.Uintptr())
}
