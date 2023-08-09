package gl

import (
	"github.com/go-gl/gl/all-core/gl"
	"log"
)

func init() {
	if err := gl.Init(); err != nil {
		log.Fatalf("failed to initialize gl package: %v\n", err.Error())
	}
	gl.Enable(gl.MULTISAMPLE)
}
