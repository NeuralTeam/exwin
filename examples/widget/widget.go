package main

import (
	_ "embed"
	"github.com/NeuralTeam/exwin"
	"github.com/NeuralTeam/exwin/pkg/window"
	"github.com/go-vgo/robotgo"
	"github.com/tfriedel6/canvas"
	"image/color"
	"time"
)

//go:embed Roboto-Light.ttf
var RobotoLight []byte

func LoadFont(c *canvas.Canvas) (err error) {
	var font *canvas.Font
	if font, err = c.LoadFont(RobotoLight); err != nil {
		return
	}
	c.SetFont(font, 70.0)
	c.SetTextAlign(canvas.Center)
	return
}

func NewWindow() (wnd exwin.Window, err error) {
	width, height := robotgo.GetScaleSize()
	width /= 6
	height /= 6

	if wnd, err = exwin.New(width, height, "widget"); err != nil {
		return
	}
	wnd.SetAttributes(window.TransparentAttributes)
	wnd.SetStyle(window.Transparent)

	wnd.SetZOrder(window.NoTopmost)
	wnd.SetPosition(width/4, height/4)
	return
}

func NewTimeWidget() (wnd exwin.Window, err error) {
	if wnd, err = NewWindow(); err != nil {
		return
	}
	b := wnd.GetBackend()
	c := b.GetCanvas()

	if err = LoadFont(c); err != nil {
		return
	}
	b.Append(func() any {
		w, h := wnd.GetSize()
		width, height := float64(w), float64(h)

		c.SetFillStyle(color.RGBA{A: 150})
		c.FillRect(
			0, 0,
			width, height,
		)

		c.SetFillStyle(color.RGBA{A: 200})
		c.FillText(
			time.Now().Format(time.Kitchen),
			width/2, height/1.5,
		)
		return nil
	})
	return
}

func main() {
	w, err := NewTimeWidget()
	if err != nil {
		panic(err)
	}

	<-time.After(time.Minute)
	w.Close()
}
