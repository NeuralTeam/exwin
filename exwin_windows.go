//go:build windows

package exwin

import (
	"github.com/NeuralTeam/exwin/pkg/backend"
	"github.com/NeuralTeam/exwin/pkg/window"
	"time"
)

type Window interface {
	Hide()
	Show()
	Close()

	BringToTop()

	GetBackend() backend.Backend

	GetOpacity() float32
	GetSize() (w, h int)
	GetPosition() (x, y int)
	GetRefreshRate() *time.Ticker
	GetMonitorRefreshRate() *time.Ticker

	GetHwnd() window.Hwnd
	GetZOrder() window.Hwnd

	SetOpacity(opacity float32)
	SetSize(w, h int)
	SetPosition(x, y int)
	SetBounds(x, y, w, h int)
	SetRefreshRate(refreshRate time.Duration)
	SetSwapInterval(interval int)
	SetAttributes(attributes window.Attributes)

	SetZOrder(order window.Hwnd)
	SetAlwaysOnTop(duration time.Duration)
	SetStyle(style window.Style)
	SetDisplayAffinity(affinity window.DisplayAffinity)
}
