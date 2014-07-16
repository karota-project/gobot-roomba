package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/karota-project/gobot-roomba"
)

func main() {
	master := gobot.NewGobot()
	api.NewAPI(master).Start()

	roombaAdaptor := roomba.NewRoombaAdaptor("roomba-a01", "/dev/tty.usbserial-FTGNZ0WQ")
	roombaDriver := roomba.NewRoombaDriver(roombaAdaptor, "roomba-d01")

	master.AddRobot(
		gobot.NewRobot(
			"roomba",
			[]gobot.Connection{roombaAdaptor},
			[]gobot.Device{roombaDriver},
			func() {
				fmt.Println("work")

				roombaDriver.Full()

				roombaDriver.DigitLedsAscii("ABCD")

				fmt.Println("finish")
			}))

	master.Start()
}
