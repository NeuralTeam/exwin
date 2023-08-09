package glfw

import (
	"github.com/tfriedel6/canvas"
)

type Program interface {
	GetCanvas() *canvas.Canvas
	SetCanvasBounds(x, y, w, h int)
	Append(action func() any, once bool) (result any)
}

func (g *Glfw) GetProgram() Program {
	return g.program
}
