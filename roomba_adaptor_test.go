package roomba

import (
  "github.com/hybridgroup/gobot"
  "testing"
)

func initTestRoombaAdaptor() *RoombaAdaptor {
  return NewRoombaAdaptor("myAdaptor")
}

func TestRoombaAdaptorConnect(t *testing.T) {
  a := initTestRoombaAdaptor()
  gobot.Expect(t, a.Connect(), true)
}

func TestRoombaAdaptorFinalize(t *testing.T) {
  a := initTestRoombaAdaptor()
  gobot.Expect(t, a.Finalize(), true)
}
