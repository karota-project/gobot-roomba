package roomba

const (
	/**
	 *This command starts the OI. You must always send the Start command before sending any other commands to the OI.
	 */
	COMMAND_START uint8 = 128

	/**
	 * This command sets the baud rate in bits per second (bps) at which OI commands and data are sent according to the baud code sent in the data byte. The default baud rate at power up is 115200 bps, but the starting baud rate can be changed to 19200 by holding down the Clean button while powering on Roomba until you hear a sequence of descending tones. Once the baud rate is changed, it persists until Roomba is power cycled by pressing the power button or removing the battery, or when the battery voltage falls below the minimum required for processor operation. You must wait 100ms after sending this command before sending additional commands at the new baud rate.
	 */
	COMMAND_BAUD uint8 = 129

	/**
	 * This command puts the OI into Safe mode, enabling user control of Roomba. It turns off all LEDs. The OI can be in Passive, Safe, or Full mode to accept this command. If a safety condition occurs (see above) Roomba reverts automatically to Passive mode.
	 */
	COMMAND_SAFE uint8 = 131

	/**
	 * This command gives you complete control over Roomba by putting the OI into Full mode, and turning off the cliff, wheel-drop and internal charger safety features. That is, in Full mode, Roomba executes any command that you send it, even if the internal charger is plugged in, or command triggers a cliff or wheel drop condition.
	 */
	COMMAND_FULL uint8 = 132

	/**
	 * This command starts the default cleaning mode.
	 */
	COMMAND_CLEAN uint8 = 135

	/**
	 * This command starts the Max cleaning mode.
	 */
	COMMAND_MAX uint8 = 136

	/**
	 * This command starts the Spot cleaning mode.
	 */
	COMMAND_SPOT uint8 = 134

	/**
	 * This command sends Roomba to the dock.
	 */
	COMMAND_SEEK_DOCK uint8 = 143

	/**
	 * This command sends Roomba a new schedule. To disable scheduled cleaning, send all 0s.
	 */
	COMMAND_SCHEDULE uint8 = 167

	/**
	 * This command sets Roomba’s clock.
	 */
	COMMAND_SET_DAY_TIME uint8 = 168

	/**
	 * This command powers down Roomba. The OI can be in Passive, Safe, or Full mode to accept this command.
	 */
	COMMAND_POWER uint8 = 133

	/**
	 * This command controls Roomba’s drive wheels. It takes four data bytes, interpreted as two 16-bit signed values using two’s complement. The first two bytes specify the average velocity of the drive wheels in millimeters per second (mm/s), with the high byte being sent first. The next two bytes specify the radius in millimeters at which Roomba will turn. The longer radii make Roomba drive straighter, while the shorter radii make Roomba turn more. The radius is measured from the center of the turning circle to the center of Roomba. A Drive command with a positive velocity and a positive radius makes Roomba drive forward while turning toward the left. A negative radius makes Roomba turn toward the right. Special cases for the radius make Roomba turn in place or drive straight, as specified below. A negative velocity makes Roomba drive backward.
	 */
	COMMAND_DRIVE uint8 = 137

	/**
	 * This command lets you control the forward and backward motion of Roomba’s drive wheels independently. It takes four data bytes, which are interpreted as two 16-bit signed values using two’s complement. The first two bytes specify the velocity of the right wheel in millimeters per second (mm/s), with the high byte sent first. The next two bytes specify the velocity of the left wheel, in the same format. A positive velocity makes that wheel drive forward, while a negative velocity makes it drive backward.
	 */
	COMMAND_DRIVE_DIRECT uint8 = 145

	/**
	 * This command lets you control the raw forward and backward motion of Roomba’s drive wheels independently. It takes four data bytes, which are interpreted as two 16-bit signed values using two’s complement. The first two bytes specify the PWM of the right wheel, with the high byte sent first. The next two bytes specify the PWM of the left wheel, in the same format. A positive PWM makes that wheel drive forward, while a negative PWM makes it drive backward.
	 */
	COMMAND_DRIVE_PWM uint8 = 146

	/**
	 * This command lets you control the forward and backward motion of Roomba’s main brush, side brush, and vacuum independently. Motor velocity cannot be controlled with this command, all motors will run at maximum speed when enabled. The main brush and side brush can be run in either direction. The vacuum only runs forward.
	 */
	COMMAND_MOTORS uint8 = 138

	/**
	 * This command lets you control the speed of Roomba’s main brush, side brush, and vacuum independently. With each data byte, you specify the duty cycle for the low side driver (max 128). For example, if you want to control a motor with 25% of battery voltage, choose a duty cycle of 128 * 25% = 32. The main brush and side brush can be run in either direction. The vacuum only runs forward. Positive speeds turn the motor in its default (cleaning) direction. Default direction for the side brush is counterclockwise. Default direction for the main brush/flapper is inward.
	 */
	COMMAND_PWM_MOTORS uint8 = 144

	/**
	  * This command controls the LEDs common to all models of Roomba 500. The Clean/Power LED is
	  specified by two data bytes: one for the color and the other for the intensity.
	*/
	COMMAND_LEDS uint8 = 139

	/**
	 * This command controls the state of the scheduling LEDs present on the Roomba 560 and 570.
	 */
	COMMAND_SCHEDULING_LEDS uint8 = 162

	/**
	 * This command controls the four 7 segment displays on the Roomba 560 and 570.
	 */
	COMMAND_DIGIT_LEDS_RAW uint8 = 163

	/**
	 * This command controls the four 7 segment displays on the Roomba 560 and 570 using ASCII character codes. Because a 7 segment display is not sufficient to display alphabetic characters properly, all characters are an approximation, and not all ASCII codes are implemented.
	 */
	COMMAND_DIGIT_LEDS_ASCII uint8 = 164

	/**
	 * This command lets you push Roomba’s buttons. The buttons will automatically release after 1/6th of a second.
	 */
	COMMAND_BUTTONS uint8 = 165

	/**
	 * This command lets you specify up to four songs to the OI that you can play at a later time. Each song is associated with a song number. The Play command uses the song number to identify your song selection. Each song can contain up to sixteen notes. Each note is associated with a note number that uses MIDI note definitions and a duration that is specified in fractions of a second. The number of data bytes varies, depending on the length of the song specified. A one note song is specified by four data bytes. For each additional note within a song, add two data bytes.
	 */
	COMMAND_SONG uint8 = 140

	/**
	 * This command lets you select a song to play from the songs added to Roomba using the Song command. You must add one or more songs to Roomba using the Song command in order for the Play command to work.
	 */
	COMMAND_PLAY uint8 = 141

	/**
	 * This command requests the OI to send a packet of sensor data bytes. There are 58 different sensor data packets. Each provides a value of a specific sensor or group of sensors.
	 */
	COMMAND_SENSORS uint8 = 142

	/**
	 * This command lets you ask for a list of sensor packets. The result is returned once, as in the Sensors command. The robot returns the packets in the order you specify.
	 */
	COMMAND_QUERY_LIST uint8 = 149

	/**
	 * This command starts a stream of data packets. The list of packets requested is sent every 15 ms, which is the rate Roomba uses to update data.
	 */
	COMMAND_STREAM uint8 = 148

	/**
	 * This command lets you stop and restart the steam without clearing the list of requested packets.
	 */
	COMMAND_PAUSE_RESUME_STREAM uint8 = 150
)
