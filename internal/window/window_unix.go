//go:build !windows

package window

import (
	"github.com/NeuralTeam/exwin/internal/backend"
	"github.com/go-gl/glfw/v3.3/glfw"
	"time"
)

type Window struct {
	RefreshRate *time.Ticker

	*backend.Backend
	*glfw.Window
}
