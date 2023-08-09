package glfw

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
	"time"
)

func init() {
	if err := glfw.Init(); err != nil {
		log.Fatalf("failed to initialize glfw package: %v\n", err.Error())
	}
	go func() {
		defer glfw.Terminate()
		for {
			select {
			case <-time.After(time.Hour):
				continue
			}
		}
	}()
}
