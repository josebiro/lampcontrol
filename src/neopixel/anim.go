package neopixel

import (
	"time"
)

// Anim - construct for handling animation
type Anim struct {
	Loop   int
	Frames []Frame
}

// Frame -- object that contains animation frame data
type Frame struct {
	Leds  Leds
	Delay time.Duration
}
