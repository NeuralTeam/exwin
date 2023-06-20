package glfw

import (
	"github.com/NeuralTeam/exdraw/glfw/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/lxn/win"
	"log"
	"time"
	"unsafe"
)

type Glfw struct {
	hwnd,
	zOrder win.HWND
	refreshRateTicker *time.Ticker

	program *gl.Program
	window  *glfw.Window
}

func initGlfw() {
	if err := glfw.Init(); err != nil {
		log.Fatalf("failed to initialize glfw package: %v\n", err.Error())
	}
}

func (g *Glfw) hints() {
	glfw.WindowHint(glfw.DoubleBuffer, glfw.True)
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)

	glfw.WindowHint(glfw.Floating, glfw.True)
	glfw.WindowHint(glfw.ScaleToMonitor, glfw.True)
	glfw.WindowHint(glfw.AutoIconify, glfw.False)

	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Focused, glfw.False)

	glfw.WindowHint(glfw.StencilBits, 8)
	glfw.WindowHint(glfw.DepthBits, 0)
}

func (g *Glfw) mousePassthrough() {
	globalStyle := win.GWL_EXSTYLE
	extendedStyle := win.GetWindowLongPtr(g.hwnd, int32(globalStyle)) |
		win.WS_EX_TOOLWINDOW |
		win.WS_EX_LAYERED | win.WS_EX_TRANSPARENT
	extendedStyle &= ^uintptr(win.WS_EX_APPWINDOW)
	win.SetWindowLongPtr(g.hwnd, globalStyle, extendedStyle)
}

func (g *Glfw) insertAfter(order win.HWND) {
	if g.zOrder != win.HWND_MESSAGE {
		win.SetWindowPos(
			g.hwnd,
			order,
			0, 0,
			0, 0,
			win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_NOACTIVATE,
		)
		return
	}

	var zOrderOfGlfwWindow, zOrderOfActiveWindow int
	// Getting the Z-order of the window
	for hwnd := g.hwnd; hwnd != 0; hwnd = win.GetWindow(hwnd, win.GW_HWNDPREV) {
		zOrderOfGlfwWindow++
	}
	// Getting the Z-order of the active window
	for hwnd := win.GetForegroundWindow(); hwnd != 0; hwnd = win.GetWindow(hwnd, win.GW_HWNDPREV) {
		zOrderOfActiveWindow++
	}

	// If the window is below the active window,
	// then brings the window to top
	if zOrderOfGlfwWindow >= zOrderOfActiveWindow {
		win.BringWindowToTop(g.hwnd)
	}
}

func (g *Glfw) insertAfterByInterval() {
	go func() {
		t := time.NewTicker(1 * time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				g.insertAfter(g.zOrder)
			}
		}
	}()
}

func (g *Glfw) render() {
	defer g.refreshRateTicker.Stop()
	for !g.window.ShouldClose() {
		select {
		case <-g.refreshRateTicker.C:
			g.window.MakeContextCurrent()
			g.program.Frame(g.window)
		}
	}
}

func (g *Glfw) createWindow(width, height int) {
	initGlfw()

	g.hints()
	if window, err := glfw.CreateWindow(width, height, "", nil, nil); err == nil {
		window.MakeContextCurrent()
		glfw.SwapInterval(1)

		g.window = window
		g.hwnd = win.HWND(unsafe.Pointer(g.window.GetWin32Window()))
		g.refreshRateTicker = time.NewTicker(time.Second / time.Duration(
			glfw.GetPrimaryMonitor().GetVideoMode().RefreshRate))
	} else {
		glfw.Terminate()
		log.Fatalf("failed to create window: %v\n", err.Error())
	}
	g.mousePassthrough()
	g.insertAfterByInterval()
}
