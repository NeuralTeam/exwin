package glfw

import (
	"github.com/NeuralTeam/exdraw/internal/glfw/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/lxn/win"
	"time"
)

func New(window *Glfw, width, height int) {
	window.createWindow(width, height)
	window.program = gl.New(width, height)
	window.render()
}

func (g *Glfw) Close() {
	g.window.SetShouldClose(true)
}

func (g *Glfw) Hide() *Glfw {
	g.window.Hide()
	return g
}

func (g *Glfw) Show() *Glfw {
	g.window.Show()
	return g
}

func (g *Glfw) GetHwnd() win.HWND {
	return g.hwnd
}

func (g *Glfw) SetSize(w, h int) *Glfw {
	g.window.SetSize(w, h)
	return g
}

func (g *Glfw) SetPosition(x, y int) *Glfw {
	g.window.SetPos(x, y)
	return g
}

func (g *Glfw) SetBounds(x, y, w, h int) *Glfw {
	g.program.Append(func() any {
		g.program.SetCanvasBounds(x, y, w, h)
		return nil
	}, true)
	return g
}

func (g *Glfw) SetRefreshRate(refreshRate int) *Glfw {
	g.refreshRateTicker.Reset(time.Second / time.Duration(refreshRate))
	return g
}

func (g *Glfw) SetSwapInterval(interval int) *Glfw {
	g.program.Append(func() any {
		glfw.SwapInterval(interval)
		return nil
	}, true)
	return g
}

func (g *Glfw) SetZOrder(order win.HWND) *Glfw {
	g.window.SetAttrib(glfw.Floating, glfw.False)
	g.zOrder = order
	return g
}
