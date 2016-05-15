package neopixel

import (
	"time"
)

type Anim struct {
	Loop   int
	Frames []Frame
}

type Frame struct {
	Leds  Leds
	Delay time.Duration
}
