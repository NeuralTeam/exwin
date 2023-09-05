package backend

import "github.com/tfriedel6/canvas"

type Backend interface {
	Append(action Action) (result any)
	AppendAndDelete(action Action) (result any)

	GetCanvas() *canvas.Canvas
}
