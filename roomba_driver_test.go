package roomba

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestRoombaDriver() *RoombaDriver {
	return NewRoombaDriver(NewRoombaAdaptor("myAdaptor"), "myDriver")
}

func TestRoombaDriverStart(t *testing.T) {
	d := initTestRoombaDriver()
	gobot.Expect(t, d.Start(), true)
}

func TestRoombaDriverHalt(t *testing.T) {
	d := initTestRoombaDriver()
	gobot.Expect(t, d.Halt(), true)
}
