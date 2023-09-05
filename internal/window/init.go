package window

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
	"time"
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
		log.Fatalf("failed to initialize glfw package: %v\n", err.Error())
	}
	go func() {
		defer glfw.Terminate()
		for range time.Tick(time.Hour) {
			continue
		}
	}()
	SetHints()
}
