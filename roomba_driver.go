package roomba

import (
  "github.com/hybridgroup/gobot"
)

type RoombaDriver struct {
  gobot.Driver
}

type RoombaInterface interface {
}

func NewRoombaDriver(a *RoombaAdaptor, name string) *RoombaDriver {
  return &RoombaDriver{
    Driver: *gobot.NewDriver(
      name,
      "roomba.RoombaDriver",
      a,
    ),
  }
}

func (r *RoombaDriver) adaptor() *RoombaAdaptor {
  return r.Driver.Adaptor().(*RoombaAdaptor)
}

func (r *RoombaDriver) Start() bool { return true }
func (r *RoombaDriver) Halt() bool { return true }
