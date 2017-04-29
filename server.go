package main

import (
	"flag"
	"log"
	"git.josebiro.com/josebiro/lampcontrol/neopixel"
	"net/http"
	"runtime"
	"time"
)

func main() {
	var testmode = flag.Bool("testmode", false, "Run in testmode (fakes serial connects)")

	var usbport string
	var usbspeed int
	var hostport string

	if runtime.GOOS == "windows" {
		usbport = "COM3"
		usbspeed = 115200
	} else {
		usbport = "/dev/ttyACM0"
		usbspeed = 115200
	}

	// Initialize the strip with unmber of lights and usb parameters
	strip := neopixel.NewStrip(60, usbport, usbspeed, *testmode)
	log.Println("NeoPixel Strip USB initialized.")

	time.Sleep(time.Second * 1)
	//http.HandleFunc("/color", np.ColorPOST)
	http.HandleFunc("/setcolor", strip.SetColorPOST)
	http.HandleFunc("/action", strip.ActionPOST)
	//http.HandleFunc("/anim", np.AnimPOST)
	hostport = ":8080"
	log.Printf("Starting server on %v\n", hostport)
	http.ListenAndServe(hostport, nil)
}
