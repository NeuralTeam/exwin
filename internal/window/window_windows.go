//go:build windows

package window

import (
	"github.com/NeuralTeam/exwin/internal/backend"
	"github.com/NeuralTeam/exwin/pkg/window"
	"github.com/go-gl/glfw/v3.3/glfw"
	"time"
)

type Window struct {
	Hwnd window.Hwnd

	RefreshRate,
	AlwaysOnTop *time.Ticker

	*backend.Backend
	*glfw.Window
}

func (wnd *Window) BringToTop() {
	wnd.GetHwnd().BringToTop()
}

func (wnd *Window) GetHwnd() window.Hwnd {
	if wnd.Hwnd == 0 {
		wnd.Hwnd = window.GetHwnd(wnd.Window)
	}
	return wnd.Hwnd
}

func (wnd *Window) GetZOrder() window.Hwnd {
	return wnd.GetHwnd().GetZOrder()
}

func (wnd *Window) SetZOrder(zOrder window.Hwnd) {
	wnd.GetHwnd().SetZOrder(zOrder)
}

func (wnd *Window) SetAlwaysOnTop(duration time.Duration) {
	if duration <= 0 {
		if wnd.AlwaysOnTop != nil {
			wnd.AlwaysOnTop.Stop()
		}
		return
	}
	wnd.AlwaysOnTop = time.NewTicker(duration)
	go func() {
		defer wnd.AlwaysOnTop.Stop()
		for range wnd.AlwaysOnTop.C {
			wnd.BringToTop()
		}
	}()
	return
}

func (wnd *Window) SetStyle(style window.Style) {
	style.SetStyle(wnd.GetHwnd())
}

func (wnd *Window) SetDisplayAffinity(affinity window.DisplayAffinity) {
	affinity.SetDisplayAffinity(wnd.GetHwnd())
}
