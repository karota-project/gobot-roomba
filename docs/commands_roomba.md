# Functions

## Start

This command starts the OI. You must always send the Start command before sending any other commands to the OI.

#### API Command

**Start**



## Baud

This command sets the baud rate in bits per second (bps) at which OI commands and data are sent according to the baud code sent in the data byte. The default baud rate at power up is 115200 bps, but the starting baud rate can be changed to 19200 by holding down the Clean button while powering on Roomba until you hear a sequence of descending tones. Once the baud rate is changed, it persists until Roomba is power cycled by pressing the power button or removing the battery, or when the battery voltage falls below the minimum required for processor operation. You must wait 100ms after sending this command before sending additional commands at the new baud rate.

#### Params

- **baudRate** - **uint8** - Baud code (0-11)

#### API Command

**Baud**

- **baudRate** - number (0-11)



## Safe

This command puts the OI into Safe mode, enabling user control of Roomba. It turns off all LEDs. The OI can be in Passive, Safe, or Full mode to accept this command. If a safety condition occurs (see above) Roomba reverts automatically to Passive mode.

#### API Command

**Safe**



## Full

This command gives you complete control over Roomba by putting the OI into Full mode, and turning off the cliff, wheel-drop and internal charger safety features. That is, in Full mode, Roomba executes any command that you send it, even if the internal charger is plugged in, or command triggers a cliff or wheel drop condition.

#### API Command

**Full**



## Clean

This command starts the default cleaning mode.

#### API Command

**Clean**



## Max

This command starts the Max cleaning mode.

#### API Command

**Max**



## Spot

This command starts the Spot cleaning mode.

#### API Command

**Spot**



## SeekDock

This command sends Roomba to the dock.

#### API Command

**SeekDock**



## Schedule

This command sends Roomba a new schedule.

#### Params

- **days** - **[7]time.Time** - Times are sent in 24 hour format.

#### API Command

**Schedule**

- **sunday** - string (00:00)
- **monday** - string (00:00)
- **tuesday** - string (00:00)
- **wednesday** - string (00:00)
- **thursday** - string (00:00)
- **friday** - string (00:00)
- **saturday** - string (00:00)



## DisableSchedule

This command sends Roomba to disable scheduled cleaning.

#### API Command

**DisableSchedule**



## SetDateTime

This command sets Roomba’s clock.

#### Params

- **datetime** - **time.Time** - Now date time

#### API Command

**SetDateTime**

- **dateTime** - string (2006-01-02 15:04:05 +0900 JST)



## Power

This command powers down Roomba. The OI can be in Passive, Safe, or Full mode to accept this command.

#### API Command

**Power**



## Drive

This command controls Roomba’s drive wheels. It takes four data bytes, interpreted as two 16-bit signed values using two’s complement. The first two bytes specify the average velocity of the drive wheels in millimeters per second (mm/s), with the high byte being sent first. The next two bytes specify the radius in millimeters at which Roomba will turn. The longer radii make Roomba drive straighter, while the shorter radii make Roomba turn more. The radius is measured from the center of the turning circle to the center of Roomba. A Drive command with a positive velocity and a positive radius makes Roomba drive forward while turning toward the left. A negative radius makes Roomba turn toward the right. Special cases for the radius make Roomba turn in place or drive straight, as specified below. A negative velocity makes Roomba drive backward.

#### Params

- **velocity** - **int16** - Velocity (-500 - 500 mm/s)
- **radius** - **int16** - Radius (-2000 - 2000 mm)

#### API Command

**Drive**

- **velocity** - number - Velocity (-500 - 500 mm/s)
- **radius** - number - Radius (-2000 - 2000 mm)



## DriveDirect

This command lets you control the forward and backward motion of Roomba’s drive wheels independently. It takes four data bytes, which are interpreted as two 16-bit signed values using two’s complement. The first two bytes specify the velocity of the right wheel in millimeters per second (mm/s), with the high byte sent first. The next two bytes specify the velocity of the left wheel, in the same format. A positive velocity makes that wheel drive forward, while a negative velocity makes it drive backward.

#### Params

- **rightVelocity** - **int16** - Right Velocity (-500 - 500 mm/s)
- **leftVelocity** - **int16** - Left Velocity (-500 - 500 mm/s)

#### API Command

**DriveDirect**

- rightVelocity - number - Right Velocity (-500 - 500 mm/s)
- leftVelocity - number - Left Velocity (-500 - 500 mm/s)



## DrivePwm

This command lets you control the raw forward and backward motion of Roomba’s drive wheels independently. It takes four data bytes, which are interpreted as two 16-bit signed values using two’s complement. The first two bytes specify the PWM of the right wheel, with the high byte sent first. The next two bytes specify the PWM of the left wheel, in the same format. A positive PWM makes that wheel drive forward, while a negative PWM makes it drive backward.

#### Params

- **rightPwm** - **int16** - Right wheel PWM (-255 - 255)
- **leftPwm** - **int16** - Left wheel PWM (-255 - 255)

#### API Command

**DrivePwm**

- rightPwm - number - Right wheel PWM (-255 - 255)
- leftPwm - number - Left wheel PWM (-255 - 255)



## Motors

This command lets you control the forward and backward motion of Roomba’s main brush, side brush, and vacuum independently. Motor velocity cannot be controlled with this command, all motors will run at maximum speed when enabled. The main brush and side brush can be run in either direction. The vacuum only runs forward.

#### Params

- **mainBrushDirection** - **bool** - Main brush direction
- **sideBrushClockwise** - **bool** - Side brush clockwise
- **mainBrush** - **bool** - Main brush
- **vacuum** - **bool** - Vacuum
- **sideBrush** - **bool** - Side Brush

#### API Command

**Motors**

- mainBrushDirection - boolean - Main brush direction
- sideBrushClockwise - boolean - Side brush clockwise
- mainBrush - boolean - Main brush
- vacuum - boolean - Vacuum
- sideBrush - boolean - Side Brush



## PwmMotors

This command lets you control the speed of Roomba’s main brush, side brush, and vacuum independently. With each data byte, you specify the duty cycle for the low side driver (max 128). For example, if you want to control a motor with 25% of battery voltage, choose a duty cycle of 128 * 25% = 32. The main brush and side brush can be run in either direction. The vacuum only runs forward. Positive speeds turn the motor in its default (cleaning) direction. Default direction for the side brush is counterclockwise. Default direction for the main brush/flapper is inward.

#### Params

- **mainBrushPwm** - **int8** - Main brush PWM (-127 - 127)
- **sideBrushPwm** - **int8** - Side brush PWM (-127 - 127)
- **vacuumPwm** - **uint8** - Vacuum PWM (0-127)

#### API Command

**PwmMotors**

- mainBrushPwm - number - Main brush PWM (-127 - 127)
- sideBrushPwm - number - Side brush PWM (-127 - 127)
- vacuumPwm - number - Vacuum PWM (0-127)



## Leds

This command controls the LEDs common to all models of Roomba 500. The Clean/Power LED is specified by two data bytes: one for the color and the other for the intensity.

#### Params

- **checkRobot** - **bool** - Uses an orange LED
- **dock** - **bool** - Uses a blue LED
- **spot** - **bool** - Uses green LED
- **debris** - **bool** - Uses Uses a blue LED
- **cleanPowerColor** - **uint8** - 0 = green, 255 = red. Intermediate values are intermediate colors (orange, yellow, etc).
- **cleanPowerIntensity** - **uint8** - 0 = off, 255 = full intensity. Intermediate values are intermediate intensities.

#### API Command

**Leds**

- checkRobot - boolean - Uses an orange LED
- dock - boolean - Uses a blue LED
- spot - boolean - Uses green LED
- debris - boolean - Uses Uses a blue LED
- cleanPowerColor - number - 0 = green, 255 = red. Intermediate values are intermediate colors (orange, yellow, etc).
- cleanPowerIntensity - number - 0 = off, 255 = full intensity. Intermediate values are intermediate intensities.



## SchedulingLeds

This command controls the state of the scheduling LEDs present on the Roomba 560 and 570.

#### Params

- **wheekday** - **uint8** - Weekday LED bits (0-255)
- **schedule** - **bool** - Schedule LED
- **clock** - **bool** - Clock LED
- **am** - **bool** - AM LED
- **pm** - **bool** - PM LED
- **colon** - **bool** - Colon LED

#### API Command

**SchedulingLeds**

- wheekday - number - Weekday LED bits (0-255)
- schedule - boolean - Schedule LED
- clock - boolean - Clock LED
- am - boolean - AM LED
- pm - boolean - PM LED
- colon - boolean - Colon LED



## DigitLedsRaw

This command controls the four 7 segment displays on the Roomba 560 and 570.

#### Params

- **digit** - **[4]uint8** -

#### API Command

**DigitLedsRaw**

- digit1 - number -
- digit2 - number -
- digit3 - number -
- digit4 - number -



## DigitLedsAscii

This command controls the four 7 segment displays on the Roomba 560 and 570 using ASCII character codes. Because a 7 segment display is not sufficient to display alphabetic characters properly, all characters are an approximation, and not all ASCII codes are implemented.

#### Params

- **message** - **string** -

#### API Command

**DigitLedsAscii**

- message - string -



## Buttons

This command lets you push Roomba’s buttons. The buttons will automatically release after 1/6th of a second.

#### Params

- **clock** - **bool** -
- **schedule** - **bool** -
- **day** - **bool** -
- **hour** - **bool** -
- **dock** - **bool** -
- **spot** - **bool** -
- **clean** - **bool** -

#### API Command

**Buttons**

- clock - boolean -
- schedule - boolean -
- day - boolean -
- hour - boolean -
- dock - boolean -
- spot - boolean -
- clean - boolean -



## Song

This command lets you specify up to four songs to the OI that you can play at a later time. Each song is associated with a song number. The Play command uses the song number to identify your song selection. Each song can contain up to sixteen notes. Each note is associated with a note number that uses MIDI note definitions and a duration that is specified in fractions of a second. The number of data bytes varies, depending on the length of the song specified. A one note song is specified by four data bytes. For each additional note within a song, add two data bytes.

#### Params

- **songNumber** - **uint8** -
- **songNumber** - **gobotRoomba.Note** -

#### API Command

**Song**

- songNumber - number -
- songNumber - string - JSON [{"number": 100, "duration": 200"}, {"number": 101, "duration": 201"}, ...]



## Play

This command lets you select a song to play from the songs added to Roomba using the Song command. You must add one or more songs to Roomba using the Song command in order for the Play command to work.

#### Params

- **songNumber** - **uint8** -

#### API Command

**Play**

- songNumber - number -



## Sensors

This command requests the OI to send a packet of sensor data bytes. There are 58 different sensor data packets. Each provides a value of a specific sensor or group of sensors.

#### Params

- **packetId** - **uint8** -

#### API Command

**Sensors**

- packetId - number -



## QueryList

This command lets you ask for a list of sensor packets. The result is returned once, as in the Sensors command. The robot returns the packets in the order you specify.

#### Params

- **packetIds** - **[]uint8** -

#### API Command

**QueryList**

- packetIds - string - CSV 1,2,3,...



## Stream

This command starts a stream of data packets. The list of packets requested is sent every 15 ms, which is the rate Roomba uses to update data.
This method of requesting sensor data is best if you are controlling Roomba over a wireless network (which has poor real-time characteristics) with software running on a desktop computer.

#### Params

- **packetIds** - **[]uint8** -

#### API Command

**Stream**

- packetIds - string - CSV 1,2,3,...



## PauseStream

This command lets you stop the steam without clearing the list of requested packets.

#### API Command

**PauseStream**



## ResumeStream

This command lets you restart the steam without clearing the list of requested packets.

#### API Command

**ResumeStream**

