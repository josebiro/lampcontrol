package neopixel

// Rgbw - LED Color scheme, just holding and setting color values
type Rgbw [4]uint8

// SetRed - set the red led value
func (rgbw *Rgbw) SetRed(r uint8) {
	rgbw[0] = r
}

// SetGreen - set the green led color value
func (rgbw *Rgbw) SetGreen(g uint8) {
	rgbw[1] = g
}

// SetBlue - set the blue led color value
func (rgbw *Rgbw) SetBlue(b uint8) {
	rgbw[2] = b
}

// SetWhite - set the white led color value
func (rgbw *Rgbw) SetWhite(w uint8) {
	rgbw[3] = w
}

// SetColor - set the pixel color value for all LEDs
func (rgbw *Rgbw) SetColor(r uint8, g uint8, b uint8, w uint8) {
	rgbw[0] = r
	rgbw[1] = g
	rgbw[2] = b
	rgbw[3] = w
}

// Uint8 - return string representation of the color object
func (rgbw *Rgbw) Uint8() []uint8 {
	return []uint8{rgbw[0], rgbw[1], rgbw[2], rgbw[3]}
}

// RgbwFromUintArray - Return a new RGBW from a uint8 array
func (rgbw *Rgbw) RgbwFromUintArray(a []uint8) {
	rgbw.SetRed(a[0])
	rgbw.SetGreen(a[1])
	rgbw.SetBlue(a[2])
	rgbw.SetWhite(a[3])
}
