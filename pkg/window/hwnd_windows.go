package window

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/lxn/win"
	"unsafe"
)

type Hwnd win.HWND

const (
	Bottom Hwnd = iota
	NoTopmost
	Top
	Topmost
)

func GetHwnd(window *glfw.Window) Hwnd {
	return Hwnd(unsafe.Pointer(window.GetWin32Window()))
}

func (h Hwnd) Hwnd() Hwnd {
	return [...]Hwnd{
		Hwnd(win.HWND_BOTTOM),
		Hwnd(win.HWND_NOTOPMOST),
		Hwnd(win.HWND_TOP),
		Hwnd(win.HWND_TOPMOST),
	}[h]
}

func (h Hwnd) String() string {
	return [...]string{
		`bottom`,
		`no topmost`,
		`top`,
		`topmost`,
	}[h]
}

func (h Hwnd) BringToTop() {
	var zOrder Hwnd
	for hwnd := win.GetForegroundWindow(); hwnd != 0; hwnd = win.GetWindow(hwnd, win.GW_HWNDPREV) {
		zOrder++
	}
	if h.GetZOrder() >= zOrder {
		win.BringWindowToTop(win.HWND(h))
	}
}

func (h Hwnd) GetZOrder() Hwnd {
	var zOrder Hwnd
	for hwnd := h; hwnd != 0; hwnd = Hwnd(win.GetWindow(win.HWND(hwnd), win.GW_HWNDPREV)) {
		zOrder++
	}
	return zOrder
}

func (h Hwnd) SetZOrder(zOrder Hwnd) {
	win.SetWindowPos(
		win.HWND(h),
		win.HWND(zOrder),
		0, 0,
		0, 0,
		win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_NOACTIVATE,
	)
}
