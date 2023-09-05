//go:build windows

package window

import (
	"github.com/lxn/win"
)

const (
	globalStyle = win.GWL_EXSTYLE
)

type Style int

const (
	Default Style = iota
	Transparent
)

func (s Style) Uintptr(hwnd Hwnd) uintptr {
	windowLongPtr := win.GetWindowLongPtr(win.HWND(hwnd), globalStyle)
	switch s {
	case Transparent:
		windowLongPtr &= ^uintptr(win.WS_EX_APPWINDOW)
		windowLongPtr |= win.WS_EX_TOOLWINDOW | win.WS_EX_LAYERED | win.WS_EX_TRANSPARENT
		break
	default:
		windowLongPtr &= ^uintptr(win.WS_EX_TOOLWINDOW | win.WS_EX_LAYERED | win.WS_EX_TRANSPARENT)
		windowLongPtr |= uintptr(win.WS_EX_APPWINDOW)
		break
	}
	return windowLongPtr
}

func (s Style) String() string {
	return [...]string{
		`default`,
		`transparent`,
	}[s]
}

func (s Style) SetStyle(hwnd Hwnd) {
	win.SetWindowLongPtr(win.HWND(hwnd), globalStyle, s.Uintptr(hwnd))
}
