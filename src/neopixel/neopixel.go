package neopixel

// Leds - array for holding led color values
type Leds []uint8

// GetColorMap - utility function to translate names to color values
func GetColorMap() map[string]*Rgbw {
	colorMap := make(map[string]*Rgbw)
	colorMap["red"] = &Rgbw{255, 0, 0, 0}
	colorMap["orange"] = &Rgbw{255, 127, 0, 0}
	colorMap["yellow"] = &Rgbw{255, 255, 0, 0}
	colorMap["green"] = &Rgbw{0, 255, 0, 0}
	colorMap["blue"] = &Rgbw{0, 0, 255, 0}
	colorMap["indigo"] = &Rgbw{75, 0, 130, 0}
	colorMap["violet"] = &Rgbw{138, 43, 226, 0}
	colorMap["purple"] = &Rgbw{127, 0, 127, 0}
	colorMap["white"] = &Rgbw{0, 0, 0, 255}
	colorMap["off"] = &Rgbw{0, 0, 0, 0}
	return colorMap
}

/*
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
*/
