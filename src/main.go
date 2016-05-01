package main

import (
	//"fmt"
	"log"

	ui "github.com/gizak/termui"
	"github.com/tarm/serial"
)

// ReadBuffer - some shit
func ReadBuffer(s *serial.Port) {
	buf := make([]byte, 128)
	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])
}

// SerialConnection - serial connection initializer
type SerialConnection struct {
	c *serial.Config
	s *serial.Port
}

func (s *SerialConnection) init() error {
	var err error
	s.c = &serial.Config{Name: "COM3", Baud: 250000}
	s.s, err = serial.OpenPort(s.c)
	if err != nil {
		return err
	}
	return nil
}

// SendCommand - send a serial command to the lamp
func (s *SerialConnection) SendCommand(command rune) error {
	_, err := s.s.Write([]byte(string(command)))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}

	l := ui.NewPar("Initializing Serial Connection...")
	l.Height = 3
	l.Width = 50
	l.TextFgColor = ui.ColorRed
	l.BorderLabel = "Loading..."
	l.BorderFg = ui.ColorRed

	ui.Render(l)

	s := new(SerialConnection)
	err = s.init()
	if err != nil {
		log.Fatal(err)
	}

	ui.Close()

	err = ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	strs := []string{
		"[1] White light",
		"[2] Red light",
		"[3] Green light",
		"[4] Blue light",
		"[5] Purple light",
		"[6] Violet light",
		"[7] Color Test",
		"[8] Rainbow",
		"[9] ...",
		"[0] Turn off lamp",
		"[Q] Quit Lamp Program",
	}

	ls := ui.NewList()
	ls.Items = strs
	ls.ItemFgColor = ui.ColorYellow
	ls.BorderLabel = "Lamp Control"
	ls.Height = ui.TermHeight()
	ls.Width = ui.TermWidth()
	ls.Y = 0

	ui.Render(ls)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		s.SendCommand('q')
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/Q", func(ui.Event) {
		s.SendCommand('q')
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/0", func(ui.Event) {
		s.SendCommand('q')
	})
	ui.Handle("/sys/kbd/1", func(ui.Event) {
		s.SendCommand('w')
	})
	ui.Handle("/sys/kbd/2", func(ui.Event) {
		s.SendCommand('r')
	})
	ui.Handle("/sys/kbd/3", func(ui.Event) {
		s.SendCommand('g')
	})
	ui.Handle("/sys/kbd/4", func(ui.Event) {
		s.SendCommand('b')
	})
	ui.Handle("/sys/kbd/5", func(ui.Event) {
		s.SendCommand('p')
	})
	ui.Handle("/sys/kbd/6", func(ui.Event) {
		s.SendCommand('v')
	})
	ui.Handle("/sys/kbd/7", func(ui.Event) {
		s.SendCommand('1')
	})
	ui.Handle("/sys/kbd/8", func(ui.Event) {
		s.SendCommand('3')
	})

	ui.Loop()

}
