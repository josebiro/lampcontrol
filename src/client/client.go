package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"flag"
	"log"
	"neopixel"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <endpoint> <arg>", os.Args[0])
		os.Exit(1)
	}

	endpoint := os.Args[1]

	var hostport string
	hostport = "localhost:8080"

	url := "http://" + hostport + "/" + endpoint
	log.Println("URL:>", url)

	var jsonStr []byte

	switch endpoint {
	case "setcolor":
		command := new(neopixel.Color)
		command.ColorName = os.Args[2]
		jsonStr, _ = json.Marshal(command)
	case "action":
		command := new(neopixel.Action)
		command.Do = os.Args[2]
		jsonStr, _ = json.Marshal(command)
	default:
		command := new(neopixel.Color)
		command.ColorName = "off"
		jsonStr, _ = json.Marshal(command)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "lampcontrol")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
