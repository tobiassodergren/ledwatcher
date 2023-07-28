package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/tobiassodergren/ledwatcher/pkg/led"
	"github.com/tobiassodergren/ledwatcher/pkg/mqtt"
)

var verbose bool
var device, broker, username, password string
var gain, port int
var treshold uint64

func init() {
	flag.StringVar(&broker, "b", "192.168.1.163", "The MQTT broker host. Default 192.168.1.163")
	flag.StringVar(&device, "d", "/dev/i2c-1", "The device to read from. Default: /dev/i2c-1")
	flag.IntVar(&gain, "g", 1, "Set the gain used when reading from led. Default: 1")
	flag.IntVar(&port, "o", 1883, "The MQTT broker port. Default: 1833")
	flag.StringVar(&password, "p", "hapwd", "The MQTT password")
	flag.Uint64Var(&treshold, "t", 100, "Set the treshold for when light is lit. Default: 100")
	flag.StringVar(&username, "u", "ha", "The MQTT username")
	flag.BoolVar(&verbose, "v", false, "Set verbose command. Default: false")
}

func sout(a ...interface{}) {
	if verbose {
		fmt.Println(a...)
	}
}

func main() {

	c := make(chan led.Status)

	flag.Parse()

	if verbose {
		sout("*** Welcome to ledwatcher ***")
		sout("Monitoring device: ", device)
		sout("Gain: ", gain)
		sout("Treshold: ", treshold)
	}

	mqtt.Init(verbose, broker, port, username, password)

	go led.Read(verbose, gain, treshold, device, c)

	for {
		led, _ := <-c
		if verbose {
			sout("Got info that isLit: ", led.IsLit, "Value: ", led.Value, "Treshold: ", led.Treshold)
		}
		data, err := json.Marshal(led)
		if err == nil {
			mqtt.Send(string(data))
		} else {
			sout("Error converting value: " + err.Error())
		}
	}
}
