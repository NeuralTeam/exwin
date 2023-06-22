package glfw

import (
	"github.com/tfriedel6/canvas"
)

type Program interface {
	Canvas() *canvas.Canvas
	CanvasBounds(x, y, w, h int)
	Append(action func() any, once bool) (result any)
}

func (g *Glfw) Program() Program {
	return g.program
}
