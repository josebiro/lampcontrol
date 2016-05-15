package neopixel

import (
	"log"
)

type Led struct {
	Color Rgbw
}

type LedStrip struct {
	Leds []*Led
}

func (l *LedStrip) SetColor(i int, r Rgbw) {
	pixel := l.Leds[i]
	pixel.Color = r
}

func (l *LedStrip) Reset() {
	off := Rgbw{0, 0, 0, 0}
	for i := range l.Leds {
		pixel := l.Leds[i]
		pixel.Color = off
	}
}

// NewStrip - initialize a number of LEDs as a strip and set to black (off)
func NewStrip(leds int) LedStrip {
	log.Println("Creating new strip: ", leds)
	off := Rgbw{0, 0, 0, 0}
	// make an array to hold the leds...
	ledArray := make([]*Led, leds)
	for i := range ledArray {
		ledArray[i] = &Led{Color: off}
	}
	var s LedStrip
	s.Leds = ledArray
	return s
}

// OneBigArray - Flatten the array structure to a uint8 array
func (s *LedStrip) OneBigArray() []uint8 {
	var output []uint8
	for i := range s.Leds {
		pixel := s.Leds[i]
		c := pixel.Color.Uint8()
		output = append(output, c...)
	}
	return output
}
