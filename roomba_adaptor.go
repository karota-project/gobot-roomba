package roomba

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/tarm/goserial"
	"io"
)

type RoombaAdaptor struct {
	gobot.Adaptor
	sp io.ReadWriteCloser
}

func NewRoombaAdaptor(name string, port string) *RoombaAdaptor {
	return &RoombaAdaptor{
		Adaptor: *gobot.NewAdaptor(
			name,
			"roomba.RoombaAdaptor",
			port,
		),
	}
}

func (r *RoombaAdaptor) Connect() bool {
	if r.Connected() == true {
		disconnect(r)
	}
	connect(r)
	return true
}

func (r *RoombaAdaptor) Finalize() bool {
	disconnect(r)
	return true
}

func connect(r *RoombaAdaptor) {
	s, err := serial.OpenPort(&serial.Config{Name: r.Port(), Baud: 115200})
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	r.sp = s
	r.SetConnected(true)
}

func disconnect(r *RoombaAdaptor) {
	r.sp.Close()
	r.SetConnected(false)
}
