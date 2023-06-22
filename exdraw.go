package exdraw

import (
	"github.com/NeuralTeam/exdraw/pkg/glfw"
	"reflect"
	"runtime"
	"time"
)

func Window(width, height int) (window *glfw.Glfw) {
	window = new(glfw.Glfw)
	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		render, terminate := glfw.New(window, width, height)
		defer terminate()
		render()
	}()
	t := time.NewTicker(time.Millisecond)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			if !reflect.ValueOf(window.Program()).IsNil() {
				return
			}
		}
	}
}
