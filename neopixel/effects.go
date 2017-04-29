package neopixel

import (
	"log"
	"time"
)

// ON and OFF contstants for LED color values
const (
	ON  = 250
	OFF = 0
)

// Setup some standard colors
var red = Rgbw{ON, OFF, OFF, OFF}
var orange = Rgbw{250, 127, 0, 0}
var yellow = Rgbw{250, 250, 0, 0}
var green = Rgbw{OFF, ON, OFF, OFF}
var blue = Rgbw{OFF, OFF, ON, OFF}
var indigo = Rgbw{75, 0, 130, 0}
var violet = Rgbw{138, 43, 226, 0}
var purple = Rgbw{127, 0, 127, 0}
var white = Rgbw{OFF, OFF, OFF, ON}
var black = Rgbw{OFF, OFF, OFF, OFF}
var off = Rgbw{OFF, OFF, OFF, OFF}

// Wheel - do the color dance
func Wheel(pos int) Rgbw {
	p := (uint8)(pos)
	p = 255 - p
	if p < 85 {
		return Rgbw{255 - p*3, 0, p * 3, 0}
	}
	if p < 170 {
		p -= 85
		return Rgbw{0, p * 3, 255 - p*3, 0}
	}
	p -= 170
	return Rgbw{p * 3, 255 - p*3, 0, 0}
}

// Rainbow - rainbow effect
func (l *LedStrip) Rainbow(wait int) {
	for j := 0; j < 256; j++ {
		for i := 0; i < l.leds; i++ {
			l.SetColor(i, Wheel((i+j)&255))
		}
		l.Sync()
		time.Sleep(time.Duration(wait) * time.Millisecond)
	}
	l.Reset()
	l.Sync()
	log.Println("Finished Rainbow sequence")
}

// TheaterChase -- another test sequece that does as the name says
func (l *LedStrip) TheaterChase(c Rgbw, wait int) {
	log.Println("Executing Theater Chase sequence")
	// Do 10 cycles
	for j := 0; j < 10; j++ {
		for q := 0; q < 3; q++ {
			for i := 0; i < l.leds; i = i + 3 {
				l.SetColor(i+q, c)
			}
			l.Sync()
			time.Sleep(time.Duration(wait) * time.Millisecond)
			for i := 0; i < l.leds; i = i + 3 {
				l.SetColor(i+q, black)
			}
		}
	}
	l.Reset()
	l.Sync()
	log.Println("Finished Theater Chase sequence")
}

// TestStrip2 - run a different strip test
func (l *LedStrip) TestStrip2() {
	log.Println("Executing test2 sequence")
	c := 0
	for p := 0; p < l.leds; p++ {
		if c == 0 {
			l.SetColor(p, red)
			c++
			continue
		}
		if c == 1 {
			l.SetColor(p, green)
			c++
			continue
		}
		if c == 2 {
			l.SetColor(p, blue)
			c++
			continue
		}
		if c == 3 {
			l.SetColor(p, white)
			c = 0
			continue
		}
	}
	l.Sync()
	time.Sleep(time.Duration(5) * time.Second)
	l.Reset()
	l.Sync()
	log.Println("Finished test2 sequence")
}

// TestStrip - run a test strip function
func (l *LedStrip) TestStrip() {
	log.Println("Executing test sequence")

	var color Rgbw

	// cycle through enough frames to test each LED
	for c := 0; c < 4; c++ {
		if c == 0 {
			color = red
		}
		if c == 1 {
			color = green
		}
		if c == 2 {
			color = blue
		}
		if c == 3 {
			color = white
		}
		count := 0
		for cycles := 0; cycles < l.leds; cycles++ {
			// Frame builder -- set a single LED on for the frame
			for i := 0; i < l.leds; i++ {
				if i == count {
					l.SetColor(i, color)
				} else {
					l.SetColor(i, off)
				}
			}
			l.Sync()
			//time.Sleep(time.Duration(50) * time.Millisecond)
			count++
		}
	}
	l.Reset()
	l.Sync()

	log.Println("Exiting test sequence")
}
