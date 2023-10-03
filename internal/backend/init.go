package backend

import (
	"github.com/go-gl/gl/all-core/gl"
)

func SetCapabilities() {
	gl.Enable(gl.MULTISAMPLE)
	gl.Enable(gl.SMOOTH)
	gl.Enable(gl.BLEND)
	gl.Enable(gl.DEPTH)
}

func init() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	SetCapabilities()
}
