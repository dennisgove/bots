package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/ollie"
	"log"
	"os"
	"time"
)

func main() {

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])

	driver := ollie.NewDriver(bleAdaptor)

	work := func() {
		println("hi!")
		gobot.Every(1*time.Second, func() {
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			driver.SetRGB(r, g, b)
		})
	}

	log.Printf("creating robot")

	robot := gobot.NewRobot("ollieBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{driver},
		work,
	)

	log.Printf("starting robot")
	if err := robot.Start(); err != nil {
		panic(err)
	}

	log.Printf("ending robot")
}
