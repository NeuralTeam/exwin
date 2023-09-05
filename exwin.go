package exwin

import "github.com/NeuralTeam/exwin/internal/window"

func New(width, height int, title string) (Window, error) {
	return window.New(width, height, title)
}
