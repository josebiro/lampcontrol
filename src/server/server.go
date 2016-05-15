package main

import (
	"log"
	"neopixel"
	"net/http"
	"runtime"
	"time"
)

func main() {
	var usbport string
	var usbspeed int
	var hostport string

	if runtime.GOOS == "windows" {
		usbport = "COM3"
		usbspeed = 250000
	} else {
		usbport = "/dev/ttyACM0"
		usbspeed = 115200
	}

	// Initialize the strip with unmber of lights and usb parameters
	np := neopixel.New(60, usbport, usbspeed)
	log.Println("NeoPixel USB initialized.")

	time.Sleep(time.Second * 1)
	http.HandleFunc("/color", np.ColorPOST)
	http.HandleFunc("/setcolor", np.SetColorPOST)
	http.HandleFunc("/action", np.ActionPOST)
	http.HandleFunc("/anim", np.AnimPOST)
	hostport = ":8080"
	log.Printf("Starting server on %v\n", hostport)
	http.ListenAndServe(hostport, nil)
}