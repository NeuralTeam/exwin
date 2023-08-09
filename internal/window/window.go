package window

import (
	"github.com/NeuralTeam/exdraw/internal/glfw"
	"reflect"
	"runtime"
	"time"
)

func New(width, height int) (window *glfw.Glfw) {
	window = new(glfw.Glfw)
	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		glfw.New(window, width, height)
	}()
	for {
		select {
		case <-time.After(time.Millisecond):
			if !reflect.ValueOf(window.GetProgram()).IsNil() {
				return
			}
		}
	}
}
