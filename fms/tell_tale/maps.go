package tell_tale

type TtsStatusType uint8

type NumberToNameMap map[uint8]string
type StatusNameMap map[TtsStatusType]string

const (
	TtsStatusOff          TtsStatusType = 0
	TtsStatusRed          TtsStatusType = 1
	TtsStatusYellow       TtsStatusType = 2
	TtsStatusInfo         TtsStatusType = 3
	TtsStatusNotAvailable TtsStatusType = 0x7
)

var numberToNameMap NumberToNameMap
var statusNameMap StatusNameMap

func initializeNTNMap() {
	numberToNameMap[1] = "Cooling Air Conditioning"
	numberToNameMap[2] = "High beam, main beam"
	numberToNameMap[3] = "Low beam, dipped beam"
	numberToNameMap[4] = "Turn Signals"
	numberToNameMap[5] = "Hazard Warning"
	numberToNameMap[6] = "Provision for the disabled or handicapped persons"
	numberToNameMap[7] = "Parking Brake"
	numberToNameMap[8] = "Brake failure/brake system malfunction"
	numberToNameMap[9] = "Hatch open"
	numberToNameMap[10] = "Fuel level"
	numberToNameMap[11] = "Engine coolant temperature"
	numberToNameMap[12] = "Battery charging condition"
	numberToNameMap[13] = "Engine oil"
	numberToNameMap[14] = "Position lights,side lights"
	numberToNameMap[15] = "Front fog light"
	numberToNameMap[16] = "Rear fog light"
	numberToNameMap[17] = "Park Heating"
	numberToNameMap[18] = "Engine / Mil indicator"
	numberToNameMap[19] = "Service, call for maintenance"
	numberToNameMap[20] = "Transmission fluid temperature"
	numberToNameMap[21] = "Transmission failure/malfunction"
	numberToNameMap[22] = "Anti-lock brake system failure"
	numberToNameMap[23] = "Worn brake linings"
	numberToNameMap[24] = "Windscreen washer fluid/windshield"
	numberToNameMap[25] = "Tire failure/malfunction"
	numberToNameMap[26] = "Malfunction/general failure"
	numberToNameMap[27] = "Engine oil temperature"
	numberToNameMap[28] = "Engine oil level"
	numberToNameMap[29] = "Engine coolant level"
	numberToNameMap[30] = "Steering fluid level"
	numberToNameMap[31] = "Steering failure"
	numberToNameMap[32] = "Height Control (Levelling)"
	numberToNameMap[33] = "Retarder"
	numberToNameMap[34] = "Engine Emission system failure (Mil indicator)"
	numberToNameMap[35] = "ESC indication"
	numberToNameMap[36] = "Brake lights"
	numberToNameMap[37] = "Articulation"
	numberToNameMap[38] = "Stop Request"
	numberToNameMap[39] = "Pram request"
	numberToNameMap[40] = "Bus stop brake"
	numberToNameMap[41] = "AdBlue level"
	numberToNameMap[42] = "Raising"
	numberToNameMap[43] = "Lowering"
	numberToNameMap[44] = "Kneeling"
	numberToNameMap[45] = "Engine compartment temperature"
	numberToNameMap[46] = "Auxiliary air pressure"
	numberToNameMap[47] = "Air filter clogged"
	numberToNameMap[48] = "Fuel filter differential pressure"
	numberToNameMap[49] = "Seat belt"
	numberToNameMap[50] = "EBS"
	numberToNameMap[51] = "Lane departure indication"
	numberToNameMap[52] = "Advanced emergency braking system"
	numberToNameMap[53] = "ACC"
	numberToNameMap[54] = "Trailer connected  washer fluid"
	numberToNameMap[55] = "ABS Trailer"
	numberToNameMap[56] = "Airbag"
	numberToNameMap[57] = "EBS Trailer"
	numberToNameMap[58] = "Tachograph indication"
	numberToNameMap[59] = "ESC switched off"
	numberToNameMap[60] = "Lane departure warning switched off"
}

func initializeSNMap() {
	statusNameMap[TtsStatusOff] = "OFF"
	statusNameMap[TtsStatusRed] = "RED"
	statusNameMap[TtsStatusYellow] = "YELLOW"
	statusNameMap[TtsStatusInfo] = "INFO"
	statusNameMap[TtsStatusNotAvailable] = "NOT AVAILABLE"
}

func GetNameForTTSNumber(number uint8) string {
	if len(numberToNameMap) == 0 {
		initializeNTNMap()
	}

	str, ok := numberToNameMap[number]
	if !ok {
		return ""
	}
	return str
}

func GetStatusName(status TtsStatusType) string {

	if len(statusNameMap) == 0 {
		initializeSNMap()
	}
	str, ok := statusNameMap[status]
	if !ok {
		return ""
	}
	return str
}
