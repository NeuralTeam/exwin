package main

import (
	_ "embed"
	"exdraw"
	"exdraw/glfw"
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/tfriedel6/canvas"
	"image/color"
	"time"
)

//go:embed Roboto-Light.ttf
var robotoLight []byte

func timeWidget() (window *glfw.Glfw) {
	width, height := robotgo.GetScaleSize()
	w, h := float64(width/7), float64(height/4)

	window = exdraw.Window(int(w), int(h)).
		Position(int(w/4), int(h/4)).
		SwapInterval(100).
		ZOrder(1)

	p := window.Program()
	c := p.Canvas()

	if font, err := c.LoadFont(robotoLight); err == nil {
		c.SetFont(font, 100.0)
		c.SetTextAlign(canvas.Center)
	}
	p.Append(func() any {
		c.SetFillStyle(color.RGBA{A: 150})
		c.FillRect(0, 0, w, h)
		c.SetFillStyle(color.RGBA{A: 200})
		c.FillText(time.Now().Format("15:04"), w/2, h/1.5)
		return nil
	}, false)

	return
}

func main() {
	_ = timeWidget()

	// Same goroutine
	t := time.NewTicker(time.Hour)
	for {
		select {
		case <-t.C:
			fmt.Println("one hour has passed")
		}
	}
}
