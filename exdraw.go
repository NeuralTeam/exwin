package exdraw

import (
	"github.com/NeuralTeam/exdraw/internal/glfw"
	"github.com/NeuralTeam/exdraw/internal/window"
	"github.com/lxn/win"
)

type Window interface {
	Close()
	Hide() *glfw.Glfw
	Show() *glfw.Glfw
	GetHwnd() win.HWND
	GetProgram() glfw.Program
	SetSize(w, h int) *glfw.Glfw
	SetPosition(x, y int) *glfw.Glfw
	SetBounds(x, y, w, h int) *glfw.Glfw
	SetRefreshRate(refreshRate int) *glfw.Glfw
	SetSwapInterval(interval int) *glfw.Glfw
	SetZOrder(order win.HWND) *glfw.Glfw
}

func NewWindow(width, height int) Window {
	return window.New(width, height)
}
