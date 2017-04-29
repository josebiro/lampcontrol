package neopixel

import (
	"github.com/tarm/serial"
	"log"
	"sync"
)

// Led - object to hold specific LED color values
type Led struct {
	Color Rgbw
}

// LedStrip - a struct to hold all led strip related data
type LedStrip struct {
	sync.RWMutex
	s    *serial.Port
	leds int
	In   chan []byte
	Out  chan []byte
	Leds []*Led
}

// NewStrip - initialize a number of LEDs as a strip and set to black (off)
func NewStrip(leds int, port string, baud int) *LedStrip {
	log.Println("Creating new strip: ", leds)
	log.Println("Connecting on port: ", port, " at baud ", baud)

	c := &serial.Config{Name: port, Baud: baud}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	// make an array to hold the leds...
	ledArray := make([]*Led, leds)
	for i := range ledArray {
		ledArray[i] = &Led{Color: off}
	}

	// Setup the strip
	strip := &LedStrip{
		leds: leds,
		s:    s,
		In:   make(chan []uint8, 0),
		Out:  make(chan []uint8, 0),
		Leds: ledArray,
	}
	strip.Lock()
	go strip.reader()
	go strip.writer()

	return strip
}

// SetColor - set a specific pixel to a specific color
func (l *LedStrip) SetColor(i int, r Rgbw) {
	pixel := l.Leds[i]
	pixel.Color = r
}

// SetStripColor -- set entire strip to a solid color
func (l *LedStrip) SetStripColor(r Rgbw) {
	for i := range l.Leds {
		pixel := l.Leds[i]
		pixel.Color = r
	}
}

// Reset -- turn all LEDs off (set to black)
func (l *LedStrip) Reset() {
	off := Rgbw{0, 0, 0, 0}
	for i := range l.Leds {
		pixel := l.Leds[i]
		pixel.Color = off
	}
}

// Sync - write color values to strip
func (l *LedStrip) Sync() {
	l.Lock()
	l.Out <- l.OneBigArray()
}

func (l *LedStrip) reader() {
	buf := make([]byte, 128)
	for {
		_, err := l.s.Read(buf)
		if err != nil {
			log.Fatal(err.Error())
		}
		//log.Printf("Received %d bytes: [%s]", n, buf[:n])
		l.Unlock()
	}
}

func (l *LedStrip) writer() {
	for {
		b, ok := <-l.Out
		if !ok {
			log.Fatal("writer chan closed")
		}
		_, err := l.s.Write(b)
		if err != nil {
			log.Fatal(err.Error())
		}
		//log.Printf("Sent %d bytes", n)
	}
}

// OneBigArray - Flatten the array structure to a uint8 array
func (l *LedStrip) OneBigArray() []uint8 {
	var output []uint8
	for i := range l.Leds {
		pixel := l.Leds[i]
		c := pixel.Color.Uint8()
		output = append(output, c...)
	}
	return output
}
