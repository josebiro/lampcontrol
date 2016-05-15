package neopixel

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func (np *NeoPixel) ActionPOST(w http.ResponseWriter, r *http.Request) {
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
		np.TestStrip()
	}
}

func (np *NeoPixel) SetColorPOST(w http.ResponseWriter, r *http.Request) {
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

	colorMap := GetColorMap()
	log.Println(color.Value)
	if len(color.Value) == 0 {
		color.Value = colorMap[color.ColorName]
	}

	log.Printf("Setting color %v on %v leds.\n", color.Value, np.leds)
	np.SetColor(color.Value)
	np.Sync()
	w.WriteHeader(http.StatusOK)
}

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
