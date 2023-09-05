package window

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Attributes int

const (
	DefaultAttributes Attributes = iota
	TransparentAttributes
	FloatingAttributes
)

func (a Attributes) String() string {
	return [...]string{
		`default attributes`,
		`transparent attributes`,
		`floating attributes`,
	}[a]
}

func (a Attributes) SetAttributes(window *glfw.Window) {
	switch a {
	case FloatingAttributes:
		window.SetAttrib(glfw.Floating, glfw.True)
		fallthrough
	case TransparentAttributes:
		window.SetAttrib(glfw.Decorated, glfw.False)
		break
	default:
		window.SetAttrib(glfw.Decorated, glfw.True)
		break
	}
}
