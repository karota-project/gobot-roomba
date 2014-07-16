package roomba

import (
  "github.com/hybridgroup/gobot"
)

type RoombaAdaptor struct {
  gobot.Adaptor
}

func NewRoombaAdaptor(name string) *RoombaAdaptor {
  return &RoombaAdaptor{
    Adaptor: *gobot.NewAdaptor(
      name,
      "roomba.RoombaAdaptor",
    ),
  }
}

func (r *RoombaAdaptor) Connect() bool {
  return true
}

func (r *RoombaAdaptor) Finalize() bool {
  return true
}
