package j1939

// http://en.wikipedia.org/wiki/Table_of_OBD-II_Codes

// PIDs
// @see https://en.wikipedia.org/wiki/OBD-II_PIDs
const (
	SUPPORTED_PIDS_1_20              = 0  // 0x00 - bit encoded
	MONITOR_STATUS_SINCE_DTC_CLEARED = 1  // 0x01 - bit encoded
	FREEZE_DTC                       = 2  // 0x02 -
	FUEL_SYSTEM_STATUS               = 3  // 0x03 - bit encoded
	ENGINE_LOAD                      = 4  // 0x04 - %
	ENGINE_COOLANT_TEMP              = 5  // 0x05 - °C
	SHORT_TERM_FUEL_TRIM_BANK_1      = 6  // 0x06 - %
	LONG_TERM_FUEL_TRIM_BANK_1       = 7  // 0x07 - %
	SHORT_TERM_FUEL_TRIM_BANK_2      = 8  // 0x08 - %
	LONG_TERM_FUEL_TRIM_BANK_2       = 9  // 0x09 - %
	FUEL_PRESSURE                    = 10 // 0x0A - kPa
	INTAKE_MANIFOLD_ABS_PRESSURE     = 11 // 0x0B - kPa
	ENGINE_RPM                       = 12 // 0x0C - rpm
	VEHICLE_SPEED                    = 13 // 0x0D - km/h
	TIMING_ADVANCE                   = 14 // 0x0E - ° before TDC
	INTAKE_AIR_TEMP                  = 15 // 0x0F - °C
	MAF_FLOW_RATE                    = 16 // 0x10 - g/s
	THROTTLE_POSITION                = 17 // 0x11 - %
	COMMANDED_SECONDARY_AIR_STATUS   = 18 // 0x12 - bit encoded
	OXYGEN_SENSORS_PRESENT_2_BANKS   = 19 // 0x13 - bit encoded
	OXYGEN_SENSOR_1_A                = 20 // 0x14 - V %
	OXYGEN_SENSOR_2_A                = 21 // 0x15 - V %
	OXYGEN_SENSOR_3_A                = 22 // 0x16 - V %
	OXYGEN_SENSOR_4_A                = 23 // 0x17 - V %
	OXYGEN_SENSOR_5_A                = 24 // 0x18 - V %
	OXYGEN_SENSOR_6_A                = 25 // 0x19 - V %
	OXYGEN_SENSOR_7_A                = 26 // 0x1A - V %
	OXYGEN_SENSOR_8_A                = 27 // 0x1B - V %
	OBD_STANDARDS                    = 28 // 0x1C - bit encoded
	OXYGEN_SENSORS_PRESENT_4_BANKS   = 29 // 0x1D - bit encoded
	AUX_INPUT_STATUS                 = 30 // 0x1E - bit encoded
	RUN_TIME_SINCE_ENGINE_START      = 31 // 0x1F - sec

	SUPPORTED_PIDS_21_40          = 32 // 0x20 - bit encoded
	DISTANCE_TRAVELED_WITH_MIL_ON = 33 // 0x21 - km
	FUEL_RAIL_PRESSURE            = 34 // 0x22 - kPa
	FUEL_RAIL_GUAGE_PRESSURE      = 35 // 0x23 - kPa
	OXYGEN_SENSOR_1_B             = 36 // 0x24 - ratio V
	OXYGEN_SENSOR_2_B             = 37 // 0x25 - ratio V
	OXYGEN_SENSOR_3_B             = 38 // 0x26 - ratio V
	OXYGEN_SENSOR_4_B             = 39 // 0x27 - ratio V
	OXYGEN_SENSOR_5_B             = 40 // 0x28 - ratio V
	OXYGEN_SENSOR_6_B             = 41 // 0x29 - ratio V
	OXYGEN_SENSOR_7_B             = 42 // 0x2A - ratio V
	OXYGEN_SENSOR_8_B             = 43 // 0x2B - ratio V
	COMMANDED_EGR                 = 44 // 0x2C - %
	EGR_ERROR                     = 45 // 0x2D - %
	COMMANDED_EVAPORATIVE_PURGE   = 46 // 0x2E - %
	FUEL_TANK_LEVEL_INPUT         = 47 // 0x2F - %
	WARM_UPS_SINCE_CODES_CLEARED  = 48 // 0x30 - count
	DIST_TRAV_SINCE_CODES_CLEARED = 49 // 0x31 - km
	EVAP_SYSTEM_VAPOR_PRESSURE    = 50 // 0x32 - Pa
	ABS_BAROMETRIC_PRESSURE       = 51 // 0x33 - kPa
	OXYGEN_SENSOR_1_C             = 52 // 0x34 - ratio mA
	OXYGEN_SENSOR_2_C             = 53 // 0x35 - ratio mA
	OXYGEN_SENSOR_3_C             = 54 // 0x36 - ratio mA
	OXYGEN_SENSOR_4_C             = 55 // 0x37 - ratio mA
	OXYGEN_SENSOR_5_C             = 56 // 0x38 - ratio mA
	OXYGEN_SENSOR_6_C             = 57 // 0x39 - ratio mA
	OXYGEN_SENSOR_7_C             = 58 // 0x3A - ratio mA
	OXYGEN_SENSOR_8_C             = 59 // 0x3B - ratio mA
	CATALYST_TEMP_BANK_1_SENSOR_1 = 60 // 0x3C - °C
	CATALYST_TEMP_BANK_2_SENSOR_1 = 61 // 0x3D - °C
	CATALYST_TEMP_BANK_1_SENSOR_2 = 62 // 0x3E - °C
	CATALYST_TEMP_BANK_2_SENSOR_2 = 63 // 0x3F - °C

	SUPPORTED_PIDS_41_60             = 64 // 0x40 - bit encoded
	MONITOR_STATUS_THIS_DRIVE_CYCLE  = 65 // 0x41 - bit encoded
	CONTROL_MODULE_VOLTAGE           = 66 // 0x42 - V
	ABS_LOAD_VALUE                   = 67 // 0x43 - %
	FUEL_AIR_COMMANDED_EQUIV_RATIO   = 68 // 0x44 - ratio
	RELATIVE_THROTTLE_POSITION       = 69 // 0x45 - %
	AMBIENT_AIR_TEMP                 = 70 // 0x46 - °C
	ABS_THROTTLE_POSITION_B          = 71 // 0x47 - %
	ABS_THROTTLE_POSITION_C          = 72 // 0x48 - %
	ABS_THROTTLE_POSITION_D          = 73 // 0x49 - %
	ABS_THROTTLE_POSITION_E          = 74 // 0x4A - %
	ABS_THROTTLE_POSITION_F          = 75 // 0x4B - %
	COMMANDED_THROTTLE_ACTUATOR      = 76 // 0x4C - %
	TIME_RUN_WITH_MIL_ON             = 77 // 0x4D - min
	TIME_SINCE_CODES_CLEARED         = 78 // 0x4E - min
	MAX_VALUES_EQUIV_V_I_PRESSURE    = 79 // 0x4F - ratio V mA kPa
	MAX_MAF_RATE                     = 80 // 0x50 - g/s
	FUEL_TYPE                        = 81 // 0x51 - ref table
	ETHONOL_FUEL_PERCENT             = 82 // 0x52 - %
	ABS_EVAP_SYS_VAPOR_PRESSURE      = 83 // 0x53 - kPa
	EVAP_SYS_VAPOR_PRESSURE          = 84 // 0x54 - Pa
	SHORT_TERM_SEC_OXY_SENS_TRIM_1_3 = 85 // 0x55 - %
	LONG_TERM_SEC_OXY_SENS_TRIM_1_3  = 86 // 0x56 - %
	SHORT_TERM_SEC_OXY_SENS_TRIM_2_4 = 87 // 0x57 - %
	LONG_TERM_SEC_OXY_SENS_TRIM_2_4  = 88 // 0x58 - %
	FUEL_RAIL_ABS_PRESSURE           = 89 // 0x59 - kPa
	RELATIVE_ACCELERATOR_PEDAL_POS   = 90 // 0x5A - %
	HYBRID_BATTERY_REMAINING_LIFE    = 91 // 0x5B - %
	ENGINE_OIL_TEMP                  = 92 // 0x5C - °C
	FUEL_INJECTION_TIMING            = 93 // 0x5D - °
	ENGINE_FUEL_RATE                 = 94 // 0x5E - L/h
	EMISSION_REQUIREMENTS            = 95 // 0x5F - bit encoded

	SUPPORTED_PIDS_61_80           = 96  // 0x60 - bit encoded
	DEMANDED_ENGINE_PERCENT_TORQUE = 97  // 0x61 - %
	ACTUAL_ENGINE_TORQUE           = 98  // 0x62 - %
	ENGINE_REFERENCE_TORQUE        = 99  // 0x63 - Nm
	ENGINE_PERCENT_TORQUE_DATA     = 100 // 0x64 - %
	AUX_INPUT_OUTPUT_SUPPORTED     = 101 // 0x65 - bit encoded
)
