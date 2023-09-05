package backend

import (
	"github.com/NeuralTeam/exwin/pkg/backend"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
	"sync"
)

type Backend struct {
	Actions,
	Returns sync.Map

	*goglbackend.GoGLBackend
	Canvas *canvas.Canvas
}

func New(window *glfw.Window) (c *Backend, err error) {
	w, h := window.GetFramebufferSize()

	c = new(Backend)
	if c.GoGLBackend, err = goglbackend.New(0, 0, w, h, nil); err != nil {
		return
	}
	c.Canvas = canvas.New(c.GoGLBackend)
	return
}

func (c *Backend) Frame(window *glfw.Window) {
	if window.ShouldClose() {
		return
	}
	window.MakeContextCurrent()

	glfw.PollEvents()
	c.Actions.Range(func(key, value interface{}) bool {
		if window.ShouldClose() {
			return false
		}
		c.Returns.Store(key, value.(backend.Action)())
		return true
	})
	window.SwapBuffers()
}

func (c *Backend) Append(action backend.Action) (result any) {
	c.Store(&action)
	return c.Load(&action)
}

func (c *Backend) AppendAndDelete(action backend.Action) (result any) {
	c.Store(&action)
	defer c.Delete(&action)
	return c.Load(&action)
}

func (c *Backend) Store(action *backend.Action) {
	c.Actions.Store(action, *action)
}

func (c *Backend) Load(action *backend.Action) (result any) {
	for ok := false; !ok; result, ok = c.Returns.Load(action) {
		continue
	}
	return
}

func (c *Backend) Delete(action *backend.Action) {
	c.Actions.Delete(action)
	c.Returns.Delete(action)
}

func (c *Backend) GetCanvas() *canvas.Canvas {
	return c.Canvas
}

func (c *Backend) SetCanvasBounds(x, y, w, h int) {
	c.GoGLBackend.SetBounds(x, y, w, h)
	c.Canvas = canvas.New(c.GoGLBackend)
}
