package window

import (
	_backend "github.com/NeuralTeam/exwin/internal/backend"
	"github.com/NeuralTeam/exwin/pkg/backend"
	"github.com/NeuralTeam/exwin/pkg/window"
	"github.com/go-gl/glfw/v3.3/glfw"
	"reflect"
	"runtime"
	"time"
)

func New(width, height int, title string) (window *Window, err error) {
	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		window = new(Window)
		if err = window.New(width, height, title); err != nil {
			return
		}
		window.Render()
	}()
	for reflect.ValueOf(window).IsNil() ||
		reflect.ValueOf(window.Backend).IsNil() {
		continue
	}
	return
}

func (wnd *Window) New(width, height int, title string) (err error) {
	if wnd.Window, err = glfw.CreateWindow(width, height, title, nil, nil); err != nil {
		return
	}
	wnd.Window.MakeContextCurrent()

	if wnd.Backend, err = _backend.New(wnd.Window); err != nil {
		return
	}
	return
}

func (wnd *Window) Render() {
	wnd.SetRefreshRate(0)

	defer wnd.RefreshRate.Stop()
	for !wnd.Window.ShouldClose() {
		select {
		case <-wnd.RefreshRate.C:
			wnd.Backend.Frame(wnd.Window)
		}
	}
}

func (wnd *Window) Hide() {
	wnd.Window.Hide()
}

func (wnd *Window) Show() {
	wnd.Window.Show()
}

func (wnd *Window) Close() {
	if wnd.Window.ShouldClose() {
		return
	}
	wnd.AppendAndDelete(func() any {
		wnd.Window.Destroy()
		return nil
	})
	wnd.Window.SetShouldClose(true)
}

func (wnd *Window) GetBackend() backend.Backend {
	return wnd.Backend
}

func (wnd *Window) GetOpacity() float32 {
	return wnd.Window.GetOpacity()
}

func (wnd *Window) GetSize() (w, h int) {
	return wnd.Window.GetSize()
}

func (wnd *Window) GetPosition() (x, y int) {
	return wnd.Window.GetPos()
}

func (wnd *Window) GetRefreshRate() *time.Ticker {
	return wnd.RefreshRate
}

func (wnd *Window) GetMonitorRefreshRate() *time.Ticker {
	return time.NewTicker(
		time.Second / time.Duration(glfw.GetPrimaryMonitor().GetVideoMode().RefreshRate),
	)
}

func (wnd *Window) SetOpacity(opacity float32) {
	wnd.Window.SetOpacity(opacity)
}

func (wnd *Window) SetSize(w, h int) {
	wnd.Window.SetSize(w, h)
}

func (wnd *Window) SetPosition(x, y int) {
	wnd.Window.SetPos(x, y)
}

func (wnd *Window) SetBounds(x, y, w, h int) {
	wnd.Backend.SetCanvasBounds(x, y, w, h)
}

func (wnd *Window) SetRefreshRate(refreshRate time.Duration) {
	if wnd.RefreshRate == nil {
		wnd.RefreshRate = wnd.GetMonitorRefreshRate()
	}
	if refreshRate <= 0 {
		return
	}
	wnd.RefreshRate.Reset(time.Second / refreshRate)
}

func (wnd *Window) SetSwapInterval(interval int) {
	wnd.AppendAndDelete(func() any {
		glfw.SwapInterval(interval)
		return nil
	})
}

func (wnd *Window) SetAttributes(attributes window.Attributes) {
	attributes.SetAttributes(wnd.Window)
}
