package window

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"os"
	"os/signal"
)

func SetHints() {
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
	glfw.WindowHint(glfw.DoubleBuffer, glfw.True)

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ScaleToMonitor, glfw.True)

	glfw.WindowHint(glfw.Focused, glfw.False)

	glfw.WindowHint(glfw.StencilBits, 8)
	glfw.WindowHint(glfw.DepthBits, 0)
}

func init() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	go func() {
		defer glfw.Terminate()

		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		<-interrupt
	}()
	SetHints()
}
