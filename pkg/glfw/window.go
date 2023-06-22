package glfw

import (
	"github.com/NeuralTeam/exdraw/pkg/glfw/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/lxn/win"
	"time"
)

func New(window *Glfw, width, height int) (render, terminate func()) {
	window.createWindow(width, height)
	window.program = gl.New(width, height)
	render, terminate = window.render, func() {
		glfw.Terminate()
		window.hwnd = win.HWND_MESSAGE
	}
	return
}

func (g *Glfw) Destroy() (destroyed bool) {
	g.window.SetShouldClose(true)

	for g.hwnd != win.HWND_MESSAGE {
		continue
	}
	destroyed = true
	return
}

func (g *Glfw) Hide() *Glfw {
	g.window.Hide()
	return g
}

func (g *Glfw) Show() *Glfw {
	g.window.Show()
	return g
}

func (g *Glfw) Size(w, h int) *Glfw {
	g.window.SetSize(w, h)
	return g
}

func (g *Glfw) Position(x, y int) *Glfw {
	g.window.SetPos(x, y)
	return g
}

func (g *Glfw) Bounds(x, y, w, h int) *Glfw {
	g.program.Append(func() any {
		g.program.CanvasBounds(x, y, w, h)
		return nil
	}, true)
	return g
}

func (g *Glfw) RefreshRate(refreshRate int) *Glfw {
	g.refreshRateTicker.Reset(time.Second / time.Duration(refreshRate))
	return g
}

func (g *Glfw) SwapInterval(interval int) *Glfw {
	g.program.Append(func() any {
		glfw.SwapInterval(interval)
		return nil
	}, true)
	return g
}

func (g *Glfw) ZOrder(order win.HWND) *Glfw {
	g.window.SetAttrib(glfw.Floating, glfw.False)
	g.zOrder = order
	return g
}

func (g *Glfw) Hwnd() win.HWND {
	return g.hwnd
}
