package neopixel

import (
	"github.com/tarm/serial"
	"log"
	"sync"
	"time"
)

// ON - constant value for an on pixel
const ON = 250

// OFF - constrant value for an off pixel
const OFF = 0

// Leds - array for holding led color values
type Leds []uint8

// Color - json object for holding color connands
type Color struct {
	ColorName string  `json:"colorname"`
	Value     []uint8 `json:"colorvalue"`
}

// Action - generic action handeler object
type Action struct {
	Do string `json:"action"`
}

// NeoPixel - object to hold led strip data and serial connection
type NeoPixel struct {
	sync.RWMutex
	s      *serial.Port
	leds   int
	In     chan []byte
	Out    chan []byte
	Anim   chan *Anim
	Colors []uint8
}

// GetColorMap - utility function to translate names to color values
func GetColorMap() map[string][]uint8 {
	colorMap := make(map[string][]uint8)
	colorMap["red"] = []uint8{255, 0, 0, 0}
	colorMap["orange"] = []uint8{255, 127, 0, 0}
	colorMap["yellow"] = []uint8{255, 255, 0, 0}
	colorMap["green"] = []uint8{0, 255, 0, 0}
	colorMap["blue"] = []uint8{0, 0, 255, 0}
	colorMap["indigo"] = []uint8{75, 0, 130, 0}
	colorMap["violet"] = []uint8{138, 43, 226, 0}
	colorMap["purple"] = []uint8{127, 0, 127, 0}
	colorMap["white"] = []uint8{0, 0, 0, 255}
	colorMap["off"] = []uint8{0, 0, 0, 0}
	return colorMap
}

// New - setup a new NeoPixel strip
func New(leds int, port string, baud int) *NeoPixel {
	c := &serial.Config{Name: port, Baud: baud}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	np := &NeoPixel{
		leds:   leds,
		s:      s,
		In:     make(chan []uint8, 0),
		Out:    make(chan []uint8, 0),
		Colors: make([]uint8, leds*4),
		Anim:   make(chan *Anim),
	}
	np.Lock()
	go np.reader()
	go np.writer()
	go np.animator()

	return np
}

func (np *NeoPixel) animator() {
	for {
		// Get new animation
		a, ok := <-np.Anim
		if !ok {
			log.Fatal("animator chan closed")
		}

		// Remember previous state
		zero := np.Colors

		// Loop through frames
		for l := a.Loop; l > 0; l-- {
			for _, frame := range a.Frames {
				np.SetColors(frame.Leds)
				np.Sync()
				time.Sleep(frame.Delay * 1000000)
			}
		}

		// Reset leds
		np.SetColors(zero)
		np.Sync()
	}
}

// Sync - write color values to strip
func (np *NeoPixel) Sync() {
	np.Lock()
	np.Out <- np.Colors
}

func (np *NeoPixel) reader() {
	buf := make([]byte, 128)
	for {
		n, err := np.s.Read(buf)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("Received %d bytes: [%s]", n, buf[:n])
		np.Unlock()
		//np.In <- buf:n]
	}
}

func (np *NeoPixel) writer() {
	for {
		b, ok := <-np.Out
		if !ok {
			log.Fatal("writer chan closed")
		}
		n, err := np.s.Write(b)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("Sent %d bytes", n)
	}
}

// SetColors - set a strips color values
func (np *NeoPixel) SetColors(c []uint8) {
	if len(c) != np.leds*4 {
		return
	}
	np.Lock()
	np.Colors = c
	np.Unlock()
}

// SetColor - set a strip to a solid color
func (np *NeoPixel) SetColor(c []uint8) {
	if len(c) != 4 {
		return
	}
	var solid []uint8
	for i := 0; i < np.leds; i++ {
		solid = append(solid, c...)
	}
	np.Lock()
	np.Colors = solid
	np.Unlock()
}

// TestStrip - run a test strip function
func (np *NeoPixel) TestStrip() {
	log.Println("Executing test sequence")

	//var red = []uint8{ON, OFF, OFF, OFF}
	var red = Rgbw{ON, OFF, OFF, OFF}
	var green = Rgbw{OFF, ON, OFF, OFF}
	var blue = Rgbw{OFF, OFF, ON, OFF}
	var white = Rgbw{OFF, OFF, OFF, ON}
	var off = Rgbw{OFF, OFF, OFF, OFF}

	var color Rgbw

	// init the strip map
	strip := NewStrip(np.leds)

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
		for cycles := 0; cycles < np.leds; cycles++ {
			// Frame builder -- set a single LED on for the frame
			for i := 0; i < np.leds; i++ {
				if i == count {
					log.Println("Setting pixel ", i, " to ", color)
					strip.SetColor(i, color)
				} else {
					strip.SetColor(i, off)
				}
			}
			np.Colors = strip.OneBigArray()
			// frame should be built at this point. Display it.
			np.SetColors(np.Colors)
			np.Sync()

			time.Sleep(time.Duration(50) * time.Millisecond)
			count++
		}
	}
	strip.Reset()
	np.Colors = strip.OneBigArray()
	np.SetColor(np.Colors)
	np.Sync()

	log.Println("Exiting test sequence")
}
