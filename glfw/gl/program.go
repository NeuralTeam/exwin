package gl

import (
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
	"log"
	"sync"
	"time"
)

type Program struct {
	actions sync.Map
	returns sync.Map
	backend *goglbackend.GoGLBackend
	canvas  *canvas.Canvas
}

func (p *Program) init() {
	if err := gl.Init(); err != nil {
		log.Fatalf("failed to initialize gl package: %v\n", err.Error())
	}
	gl.Enable(gl.MULTISAMPLE)
}

func (p *Program) newCanvas(width, height int) {
	if backend, err := goglbackend.New(0, 0, width, height, nil); err == nil {
		p.backend = backend
		p.canvas = canvas.New(p.backend)
	} else {
		log.Fatalf("error loading gl backend: %v\n", err.Error())
	}
}

func New(width, height int) *Program {
	p := new(Program)
	p.init()
	p.newCanvas(width, height)
	return p
}

func (p *Program) Canvas() *canvas.Canvas {
	return p.canvas
}

func (p *Program) CanvasBounds(x, y, w, h int) {
	p.backend.SetBounds(x, y, w, h)
	p.canvas = canvas.New(p.backend)
}

func (p *Program) Frame(window *glfw.Window) *Program {
	// Poll for UI window events
	glfw.PollEvents()
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	p.actions.Range(func(key, value interface{}) bool {
		p.returns.Store(key, value.(func() any)())
		return true
	})
	// Swap in the previous rendered buffer
	window.SwapBuffers()

	return p
}

func (p *Program) Append(action func() any, once bool) (result any) {
	name := fmt.Sprintf("%p%p", &action, &once)
	p.actions.Store(name, action)
	if once {
		defer func() {
			p.actions.Delete(name)
			p.returns.Delete(name)
		}()
	}
	t := time.NewTicker(time.Nanosecond)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			if r, ok := p.returns.Load(name); ok {
				result = r
				return
			}
		}
	}
}
