package roomba

///**
// * packetId 7
// */
//type BumpsAndWheelDrops struct {
//	WheelDropLeft  bool
//	WheelDropRight bool
//	BumpLeft       bool
//	BumpRight      bool
//}
//
///**
// * packetId 8
// */
//type Wall struct {
//	IsExist bool
//}
//
///**
// * packetId 9
// */
//type CliffLeft struct {
//	IsCliff bool
//}
//
///**
// * packetId 10
// */
//type CliffFrontLeft struct {
//	IsCliff bool
//}
//
///**
// * packetId 11
// */
//type CliffFrontRight struct {
//	IsCliff bool
//}
//
///**
// * packetId 12
// */
//type CliffRight struct {
//	IsCliff bool
//}
//
///**
// * packetId 13
// */
//type VirtualWall struct {
//	IsDetect bool
//}
//
///**
// * packetId 14
// */
//type WheelOvercurrents struct {
//	LeftWheel  bool
//	RightWheel bool
//	MainBrush  bool
//	SideBrush  bool
//}
//
///**
// * packetId 15
// */
//type DirtDetect struct {
//	dirtLevel uint8
//}
//
///**
// * packetId 16
// */
///*
//type UnusedByte struct {
//  // Unused
//}
//*/
//
///**
// * packetId 17
// */
//type InfraredCharacterOmni struct {
//	value uint8
//}
//
///**
// * packetId 52
// */
//type InfraredCharacterLeft struct {
//	value uint8
//}
//
///**
// * packetId 53
// */
//type InfraredCharacterRight struct {
//	value uint8
//}
//
//const (
//	// IR Remote Control
//	LEFT        uint8 = 129
//	FORWARD     uint8 = 130
//	RIGHT       uint8 = 131
//	SPOT        uint8 = 132
//	MAX         uint8 = 133
//	SMALL       uint8 = 134
//	MEDIUM      uint8 = 135
//	LARGE_CLEAN uint8 = 136
//	STOP        uint8 = 137
//	POWER       uint8 = 138
//	ARC_LEFT    uint8 = 139
//	ARC_RIGHT   uint8 = 140
//	STOP        uint8 = 141
//
//	// Scheduling Remote
//	DOWNLOAD  uint8 = 142
//	SEEK_DOCK uint8 = 143
//
//	// Roomba Discovery Drive-on Charger
//	RESERVED                        uint8 = 240
//	RED_BUOY                        uint8 = 248
//	GREEN_BUOY                      uint8 = 244
//	FORCE_FIELD                     uint8 = 242
//	RED_BUOY_AND_GREEN_BUOY         uint8 = 252
//	RED_BUOY_AND_FORCE_FIELD        uint8 = 250
//	GREEN_BUOY_AND_FORCE_FIELD      uint8 = 246
//	RED_BUOY_GREEN_BUOY_FORCE_FIELD uint8 = 254
//
//	// Roomba 500 Drive-on Charger
//	RESERVED                        uint8 = 160
//	FORCE_FIELD                     uint8 = 161
//	GREEN_BUOY                      uint8 = 164
//	GREEN_BUOY_AND_FORCE_FIELD      uint8 = 165
//	RED_BUOY                        uint8 = 168
//	RED_BUOY_AND_FORCE_FIELD        uint8 = 169
//	RED_BUOY_AND_GREEN_BUOY         uint8 = 172
//	RED_BUOY_GREEN_BUOY_FORCE_FIELD       = 173
//
//	// Roomba 500 Virtual Wall
//	VIRTUAL_WALL uint8 = 162
//
//	// Roomba 500 Virtual Wall Lightouse
//
//)
//
///**
// * packetId 18
// */
//type Buttons struct {
//	Clock    bool
//	Schedule bool
//	Day      bool
//	Hour     bool
//	Minute   bool
//	Dock     bool
//	Spot     bool
//	Clean    bool
//}
//
///**
// * packetId 19
// */
//type Distance struct {
//	value int16
//}
//
///**
// * packetId 20
// */
//type Angle struct {
//	value int16
//}
//
///**
// * packetId 21
// */
//type ChargingState struct {
//	state uint8
//}
//
//const (
//	NOT_CHARGING = iota
//	RECONDITIONING_CHARGING
//	FULL_CHARGING
//	TRICKLE_CHARGING
//	WAITING
//	CHARGING_FAULT_CONDITION
//)
//
///**
// * packetId 22
// */
//type Voltage struct {
//	value uint16
//}
//
///**
// * packetId 23
// */
//type Current struct {
//	value int16
//}
//
///**
// * packetId 24
// */
//type Temperature struct {
//	value int8
//}
//
///**
// * packetId 25
// */
//type BatteryCharge struct {
//	value uint16
//}
//
///**
// * packetId 26
// */
//type BatteryCapacity struct {
//	value uint16
//}
//
///**
// * packetId 27
// */
//type WallSignal struct {
//	value uint16
//}
//
///**
// * packetId 28
// */
//type CliffLeftSignal struct {
//	value uint16
//}
//
///**
// * packetId 29
// */
//type CliffFrontLeftSignal struct {
//	value uint16
//}
//
///**
// * packetId 30
// */
//type CliffFrontRightSignal struct {
//	value uint16
//}
//
///**
// * packetId 31
// */
//type CliffRightSignal struct {
//	value uint16
//}
//
///**
// * packetId 31
// */
//type CliffRightSignal struct {
//	value uint16
//}
//
///**
// * packetId 32
// */
///*
//type UnusedByte struct {
//  // Unused
//}
//*/
//
///**
// * packetId 33
// */
///*
//type UnusedByte struct {
//  // Unused
//}
//*/
//
///**
// * packetId 34
// */
//type ChargingSourcesAvailable struct {
//	HomeBase        bool
//	InternalCharger bool
//}
//
///**
// * packetId 35
// */
//type OIMode struct {
//	mode uint8
//}
//
//const (
//	OFF = iota
//	PASSIVE
//	SAFE
//	FULL
//)
//
///**
// * packetId 36
// */
//type SongNumber struct {
//	number uint8
//}
//
///**
// * packetId 37
// */
//type SongPlaying struct {
//	isPlaying bool
//}
//
///**
// * packetId 38
// */
//type NumberOfStreamPackets struct {
//	length uint8
//}
//
///**
// * packetId 39
// */
//type RequestedVelocity struct {
//	value int16
//}
//
///**
// * packetId 40
// */
//type RequestedRadius struct {
//	value int16
//}
//
///**
// * packetId 41
// */
//type RequestedRightVelocity struct {
//	value int16
//}
//
///**
// * packetId 42
// */
//type RequestedLeftVelocity struct {
//	value int16
//}
//
///**
// * packetId 43
// */
//type RightEncoderCounts struct {
//	value uint16
//}
//
///**
// * packetId 44
// */
//type LeftEncoderCounts struct {
//	value uint16
//}
//
///**
// * packetId 45
// */
//type LightBumper struct {
//	Right       bool
//	FrontRight  bool
//	CenterRight bool
//	CenterLeft  bool
//	FrontLeft   bool
//	Left        bool
//}
//
///**
// * packetId 46
// */
//type LightBumpLeftSignal struct {
//	value uint16
//}
//
///**
// * packetId 47
// */
//type LightBumpFrontLeftSignal struct {
//	value uint16
//}
//
///**
// * packetId 48
// */
//type LightBumpCenterLeftSignal struct {
//	value uint16
//}
//
///**
// * packetId 49
// */
//type LightBumpCenterRightSignal struct {
//	value uint16
//}
//
///**
// * packetId 50
// */
//type LightBumpFrontRightSignal struct {
//	value uint16
//}
//
///**
// * packetId 51
// */
//type LightBumpRightSignal struct {
//	value uint16
//}
//
///**
// * packetId 54
// */
//type LeftMotorCurrent struct {
//	value int16
//}
//
///**
// * packetId 55
// */
//type RightMotorCurrent struct {
//	value int16
//}
//
///**
// * packetId 56
// */
//type MainBrushMotorCurrent struct {
//	value int16
//}
//
///**
// * packetId 57
// */
//type SideBrushMotorCurrent struct {
//	value int16
//}
//
///**
// * packetId 58
// */
//type Stasis struct {
//	IsStasis bool
//}
