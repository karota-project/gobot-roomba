package roomba

import (
	"encoding/json"
	"github.com/hybridgroup/gobot"
	"strconv"
	"strings"
	"time"
)

type RoombaDriver struct {
	gobot.Driver
}

type RoombaInterface interface {
}

func NewRoombaDriver(a *RoombaAdaptor, name string) *RoombaDriver {
	r := &RoombaDriver{
		Driver: *gobot.NewDriver(
			name,
			"roomba.RoombaDriver",
			a,
		),
	}

	r.AddCommand("Start", func(params map[string]interface{}) interface{} {
		r.Start()
		return nil
	})

	r.AddCommand("Baud", func(params map[string]interface{}) interface{} {
		baudRate := uint8(params["baudRate"].(float64))
		r.Baud(baudRate)
		return nil
	})

	r.AddCommand("Safe", func(params map[string]interface{}) interface{} {
		r.Safe()
		return nil
	})

	r.AddCommand("Full", func(params map[string]interface{}) interface{} {
		r.sender(COMMAND_FULL, []uint8{})
		return nil
	})

	r.AddCommand("Clean", func(params map[string]interface{}) interface{} {
		r.sender(COMMAND_CLEAN, []uint8{})
		return nil
	})

	r.AddCommand("Max", func(params map[string]interface{}) interface{} {
		r.sender(COMMAND_MAX, []uint8{})
		return nil
	})

	r.AddCommand("Spot", func(params map[string]interface{}) interface{} {
		r.sender(COMMAND_SPOT, []uint8{})
		return nil
	})

	r.AddCommand("SeekDock", func(params map[string]interface{}) interface{} {
		r.sender(COMMAND_SEEK_DOCK, []uint8{})
		return nil
	})

	/**
	 * days[0] = sun ... days[6] = sat
	 */
	r.AddCommand("Schedule", func(params map[string]interface{}) interface{} {
		layout := "15:04"

		weekdays := []string{
			"sunday",
			"monday",
			"tuesday",
			"wednesday",
			"thursday",
			"friday",
			"saturday",
		}

		days := [7]time.Time{}

		for i, weekday := range weekdays {
			if t, ok := params[weekday].(string); ok {
        day, err := time.Parse(layout, t)
        if err != nil {
          days[i] = day
        }
			}
		}

    r.Schedule(days)
		return nil
	})

	r.AddCommand("DisableSchedule", func(params map[string]interface{}) interface{} {
		r.DisableSchedule()
		return nil
	})

	r.AddCommand("SetDateTime", func(params map[string]interface{}) interface{} {
		dateTime, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", params["dateTime"].(string))
		r.SetDateTime(dateTime)
		return nil
	})

	r.AddCommand("Power", func(params map[string]interface{}) interface{} {
		r.Power()
		return nil
	})

	r.AddCommand("Drive", func(params map[string]interface{}) interface{} {
		velocity := int16(params["velocity"].(float64))
		radius := int16(params["radius"].(float64))
		r.Drive(velocity, radius)
		return nil
	})

	r.AddCommand("DriveDirect", func(params map[string]interface{}) interface{} {
		rightVelocity := int16(params["rightVelocity"].(float64))
		leftVelocity := int16(params["leftVelocity"].(float64))
		r.DriveDirect(rightVelocity, leftVelocity)
		return nil
	})

	r.AddCommand("DrivePwm", func(params map[string]interface{}) interface{} {
		rightPwm := int16(params["rightPwm"].(float64))
		leftPwm := int16(params["leftPwm"].(float64))

		r.DrivePwm(rightPwm, leftPwm)
		return nil
	})

	r.AddCommand("Motors", func(params map[string]interface{}) interface{} {
		mainBrushDirection := params["mainBrushDirection"].(bool)
		sideBrushClockwise := params["sideBrushClockwise"].(bool)
		mainBrush := params["mainBrush"].(bool)
		vacuum := params["vacuum"].(bool)
		sideBrush := params["sideBrush"].(bool)

		r.Motors(mainBrushDirection, sideBrushClockwise, mainBrush, vacuum, sideBrush)
		return nil
	})

	r.AddCommand("PwmMotors", func(params map[string]interface{}) interface{} {
		mainBrushPwm := int8(params["mainBrushPwm"].(float64))
		sideBrushPwm := int8(params["sideBrushPwm"].(float64))
		vacuumPwm := uint8(params["vacuumPwm"].(float64))

		r.PwmMotors(mainBrushPwm, sideBrushPwm, vacuumPwm)
		return nil
	})

	r.AddCommand("Leds", func(params map[string]interface{}) interface{} {
		checkRobot := params["checkRobot"].(bool)
		dock := params["dock"].(bool)
		spot := params["spot"].(bool)
		debris := params["debris"].(bool)
		cleanPowerColor := uint8(params["cleanPowerColor"].(float64))
		cleanPowerIntensity := uint8(params["cleanPowerIntensity"].(float64))

		r.Leds(checkRobot, dock, spot, debris, cleanPowerColor, cleanPowerIntensity)
		return nil
	})

	r.AddCommand("SchedulingLeds", func(params map[string]interface{}) interface{} {
		weekday := uint8(params["weekday"].(float64))
		schedule := params["schedule"].(bool)
		clock := params["clock"].(bool)
		am := params["am"].(bool)
		pm := params["pm"].(bool)
		colon := params["colon"].(bool)

		r.SchedulingLeds(weekday, schedule, clock, am, pm, colon)
		return nil
	})

	r.AddCommand("DigitLedsRaw", func(params map[string]interface{}) interface{} {
		digit := [4]uint8{
			uint8(params["digit1"].(float64)),
			uint8(params["digit2"].(float64)),
			uint8(params["digit3"].(float64)),
			uint8(params["digit4"].(float64)),
		}

		r.DigitLedsRaw(digit)
		return nil
	})

	r.AddCommand("DigitLedsAscii", func(params map[string]interface{}) interface{} {
		message := params["message"].(string)
		r.DigitLedsAscii(message)
		return nil
	})

	r.AddCommand("Buttons", func(params map[string]interface{}) interface{} {
		clock := params["clock"].(bool)
		schedule := params["schedule"].(bool)
		day := params["day"].(bool)
		hour := params["hour"].(bool)
		dock := params["dock"].(bool)
		spot := params["spot"].(bool)
		clean := params["clean"].(bool)

		r.Buttons(clock, schedule, day, hour, dock, spot, clean)
		return nil
	})

	r.AddCommand("Song", func(params map[string]interface{}) interface{} {
		songNumber := uint8(params["songNumber"].(float64))
		notesStr := params["notes"].(string)

		notes := []Note{}
		json.Unmarshal([]byte(notesStr), &notes)
		r.Song(songNumber, notes)
		return nil
	})

	r.AddCommand("Play", func(params map[string]interface{}) interface{} {
		songNumber := uint8(params["songNumber"].(float64))
		r.Play(songNumber)
		return nil
	})

	r.AddCommand("Sensors", func(params map[string]interface{}) interface{} {
		packetId := uint8(params["packetId"].(float64))
		r.Sensors(packetId)
		return nil
	})

	r.AddCommand("QueryList", func(params map[string]interface{}) interface{} {
    s := params["packetIds"].(string)
		packetIdsStr := strings.Split(strings.Replace(s, " ", "", -1), ",")

		packetIds := []uint8{}

		for _, packetIdStr := range packetIdsStr {
			packetId, err := strconv.Atoi(packetIdStr)
      if err != nil {
        packetIds = append(packetIds, uint8(packetId))
      }
		}

		r.QueryList(packetIds)
		return nil
	})

	r.AddCommand("Stream", func(params map[string]interface{}) interface{} {
    s := params["packetIds"].(string)
		packetIdsStr := strings.Split(strings.Replace(s, " ", "", -1), ",")

		packetIds := []uint8{}

		for _, packetIdStr := range packetIdsStr {
			packetId, err := strconv.Atoi(packetIdStr)
      if err != nil {
        packetIds = append(packetIds, uint8(packetId))
      }
		}

		r.Stream(packetIds)
		return nil
	})

	r.AddCommand("PauseStream", func(params map[string]interface{}) interface{} {
		r.PauseStream()
		return nil
	})

	r.AddCommand("ResumeStream", func(params map[string]interface{}) interface{} {
		r.ResumeStream()
		return nil
	})

	return r
}

func (r *RoombaDriver) adaptor() *RoombaAdaptor {
	return r.Driver.Adaptor().(*RoombaAdaptor)
}

func (r *RoombaDriver) Start() bool {
	r.sender(COMMAND_START, []uint8{})
	return true
}

func (r *RoombaDriver) Halt() bool {
  r.sender(COMMAND_SAFE, []uint8{})
	return true
}

/*
func (r *RoombaDriver) Start() {
	r.sender(COMMAND_START, []uint8{})
}
*/

func (r *RoombaDriver) Baud(baudRate uint8) {
	r.sender(COMMAND_BAUD, []uint8{baudRate})
}

func (r *RoombaDriver) Safe() {
	r.sender(COMMAND_SAFE, []uint8{})
}

func (r *RoombaDriver) Full() {
	r.sender(COMMAND_FULL, []uint8{})
}

func (r *RoombaDriver) Clean() {
	r.sender(COMMAND_CLEAN, []uint8{})
}

func (r *RoombaDriver) Max() {
	r.sender(COMMAND_MAX, []uint8{})
}

func (r *RoombaDriver) Spot() {
	r.sender(COMMAND_SPOT, []uint8{})
}

func (r *RoombaDriver) SeekDock() {
	r.sender(COMMAND_SEEK_DOCK, []uint8{})
}

/**
 * days[0] = sun ... days[6] = sat
 */
func (r *RoombaDriver) Schedule(days [7]time.Time) {
	var daybits uint8 = 0x00
	daybytes := []uint8{}

	for i, day := range days {
		if !day.IsZero() {
			daybits |= 0x01 << uint8(i)
		}
		daybytes = append(daybytes, uint8(day.Hour()), uint8(day.Minute()))
	}

	r.sender(COMMAND_SCHEDULE, append([]uint8{daybits}, daybytes...))
}

func (r *RoombaDriver) DisableSchedule() {
	r.sender(COMMAND_SCHEDULE, []uint8{
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00,
	})
}

func (r *RoombaDriver) SetDateTime(dateTime time.Time) {
	r.sender(COMMAND_SET_DAY_TIME, []uint8{
		uint8(dateTime.Weekday()),
		uint8(dateTime.Hour()),
		uint8(dateTime.Minute()),
	})
}

func (r *RoombaDriver) Power() {
	r.sender(COMMAND_POWER, []uint8{})
}

func (r *RoombaDriver) Drive(velocity int16, radius int16) {
	r.sender(COMMAND_DRIVE, []uint8{
		uint8((velocity >> 8) & 0xFF),
		uint8((velocity >> 0) & 0xFF),
		uint8((radius >> 8) & 0xFF),
		uint8((radius >> 0) & 0xFF),
	})
}

func (r *RoombaDriver) DriveDirect(rightVelocity int16, leftVelocity int16) {
	r.sender(COMMAND_DRIVE_DIRECT, []uint8{
		uint8((rightVelocity >> 8) & 0xFF),
		uint8((rightVelocity >> 0) & 0xFF),
		uint8((leftVelocity >> 8) & 0xFF),
		uint8((leftVelocity >> 0) & 0xFF),
	})
}

func (r *RoombaDriver) DrivePwm(rightPwm int16, leftPwm int16) {
	r.sender(COMMAND_DRIVE_PWM, []uint8{
		uint8((rightPwm >> 8) & 0xFF),
		uint8((rightPwm >> 0) & 0xFF),
		uint8((leftPwm >> 8) & 0xFF),
		uint8((leftPwm >> 0) & 0xFF),
	})
}

func (r *RoombaDriver) Motors(mainBrushDirection bool, sideBrushClockwise bool, mainBrush bool, vacuum bool, sideBrush bool) {
	var motorbits uint8 = 0x00
	args := []bool{sideBrush, vacuum, mainBrush, sideBrushClockwise, mainBrushDirection}

	for i, v := range args {
		if v {
			motorbits |= 0x01 << (uint)(i)
		}
	}

	r.sender(COMMAND_MOTORS, []uint8{motorbits})
}

func (r *RoombaDriver) PwmMotors(mainBrushPwm int8, sideBrushPwm int8, vacuumPwm uint8) {
	r.sender(COMMAND_PWM_MOTORS, []uint8{
		uint8(mainBrushPwm),
		uint8(sideBrushPwm),
		uint8(vacuumPwm)})
}

func (r *RoombaDriver) Leds(checkRobot bool, dock bool, spot bool, debris bool, cleanPowerColor uint8, cleanPowerIntensity uint8) {
	var ledbits uint8 = 0x00
	args := []bool{debris, spot, dock, checkRobot}

	for i, v := range args {
		if v {
			ledbits |= 0x01 << uint(i)
		}
	}

	r.sender(COMMAND_LEDS, []uint8{ledbits, cleanPowerColor, cleanPowerIntensity})
}

func (r *RoombaDriver) SchedulingLeds(weekday uint8, schedule bool, clock bool, am bool, pm bool, colon bool) {
	var ledbits uint8 = 0x00
	args := []bool{colon, pm, am, clock, schedule}

	for i, v := range args {
		if v {
			ledbits |= 0x01 << uint(i)
		}
	}

	r.sender(COMMAND_SCHEDULING_LEDS, []uint8{weekday, ledbits})
}

func (r *RoombaDriver) DigitLedsRaw(digit [4]uint8) {
  r.sender(COMMAND_DIGIT_LEDS_RAW, digit[:4])
}

func (r *RoombaDriver) DigitLedsAscii(message string) {
	r.sender(COMMAND_DIGIT_LEDS_ASCII, []uint8(message[:4]))
}

func (r *RoombaDriver) Buttons(clock bool, schedule bool, day bool, hour bool, dock bool, spot bool, clean bool) {
	var buttonbits uint8 = 0x00
	args := []bool{clean, spot, dock, hour, day, schedule, clock}

	for i, v := range args {
		if v {
			buttonbits |= 0x01 << uint(i)
		}
	}

	r.sender(COMMAND_BUTTONS, []uint8{buttonbits})
}

func (r *RoombaDriver) Song(songNumber uint8, notes []Note) {
	datas := []uint8{songNumber, (uint8)(len(notes))}

	for _, note := range notes {
		datas = append(datas, note.Number, note.Duration)
	}

	r.sender(COMMAND_SONG, datas)
}

func (r *RoombaDriver) Play(songNumber uint8) {
	r.sender(COMMAND_PLAY, []uint8{songNumber})
}

func (r *RoombaDriver) Sensors(packetId uint8) {
	r.sender(COMMAND_SENSORS, []uint8{packetId})
}

func (r *RoombaDriver) QueryList(packetIds []uint8) {
	r.sender(COMMAND_QUERY_LIST, packetIds)
}

func (r *RoombaDriver) Stream(packetIds []uint8) {
	r.sender(COMMAND_STREAM, packetIds)
}

func (r *RoombaDriver) PauseStream() {
	r.sender(COMMAND_PAUSE_RESUME_STREAM, []uint8{0x00})
}

func (r *RoombaDriver) ResumeStream() {
	r.sender(COMMAND_PAUSE_RESUME_STREAM, []uint8{0x01})
}

func (r *RoombaDriver) sender(command uint8, datas []uint8) {
	d := append([]uint8{command}, datas...)
	r.adaptor().sp.Write(d)
}
