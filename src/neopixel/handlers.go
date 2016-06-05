package neopixel

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Color - json object for holding color connands
type Color struct {
	ColorName string  `json:"colorname"`
	Value     []uint8 `json:"colorvalue"`
	Hex       string  `json:"colorhex"`
}

// HexToValue - converts hex string to color values in uint8 array format
func (c *Color) HexToValue() {
	value := []uint8{0, 0, 0, 0}
	// check length, if not 6 characters, fail
	log.Printf("c.Hex = %v, length=%v\n", c.Hex, len(c.Hex))
	if len(c.Hex) != 7 {
		log.Println("ERROR: string is not a hex color value.")
	}
	red := "0x" + c.Hex[1:3]
	gre := "0x" + c.Hex[3:5]
	blu := "0x" + c.Hex[5:7]
	r, _ := strconv.ParseUint(red, 0, 8)
	g, _ := strconv.ParseUint(gre, 0, 8)
	b, _ := strconv.ParseUint(blu, 0, 8)
	value[0] = uint8(r)
	value[1] = uint8(g)
	value[2] = uint8(b)
	c.Value = value
}

// Action - generic action handeler object
type Action struct {
	Do string `json:"action"`
}

// ActionPOST -- Handler for Actions
func (l *LedStrip) ActionPOST(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering Action Handler")
	var action Action
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &action); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		log.Println(err.Error())
	}
	log.Println(action.Do)
	switch action.Do {
	case "test":
		l.TestStrip()
	case "test2":
		l.TestStrip2()
	case "theater":
		l.TheaterChase(white, 100)
	case "theaterred":
		l.TheaterChase(red, 100)
	case "theatergreen":
		l.TheaterChase(green, 100)
	case "theaterblue":
		l.TheaterChase(blue, 100)
	case "rainbow":
		l.Rainbow(100)
	}
}

// SetColorPOST -- Handler for setting static colors
func (l *LedStrip) SetColorPOST(w http.ResponseWriter, r *http.Request) {
	var color = new(Color)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &color); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		log.Println(err.Error())
	}

	// The only value that is important is Color.Value, everything else translates
	// into that value.

	// ColorValue is set; color name is not, color hex is not
	// ColorHex is set; color name is not, color value is not

	// ColorName is set; color value is not, color Hex is Not
	colorMap := GetColorMap()
	c := new(Rgbw)
	if color.ColorName != "" {
		c = colorMap[color.ColorName]
	} else {
		color.HexToValue()
		c.RgbwFromUintArray(color.Value)
	}
	log.Printf("Setting color %v on %v leds.\n", c, l.leds)
	l.SetStripColor(*c)
	l.Sync()
	w.WriteHeader(http.StatusOK)
}

/*
// ColorPOST -- post for setting colors of strip frames
func (np *NeoPixel) ColorPOST(w http.ResponseWriter, r *http.Request) {
	var leds Leds
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &leds); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		log.Println(err.Error())
	}
	log.Printf("Setting colors on %v leds\n", np.leds)
	np.SetColors(leds)
	np.Sync()
	w.WriteHeader(http.StatusOK)
}

// AnimPOST -- something animation related. Not sure yet.
func (np *NeoPixel) AnimPOST(w http.ResponseWriter, r *http.Request) {
	var anim Anim
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &anim); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		log.Println(err.Error())
	}
	np.Anim <- &anim
	w.WriteHeader(http.StatusOK)
}
*/
