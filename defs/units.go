package defs

import "fmt"

type EngineeringUnits uint16

// Engineering unit values.
const (
	UnitsMetersPerSecondPerSecond        EngineeringUnits = 0xA6
	UnitsSquareMeters                    EngineeringUnits = 0x00
	UnitsSquareCentimeters               EngineeringUnits = 0x74
	UnitsSquareFeet                      EngineeringUnits = 0x01
	UnitsSquareInches                    EngineeringUnits = 0x73
	UnitsMilliamperes                    EngineeringUnits = 0x02
	UnitsAmperes                         EngineeringUnits = 0x03
	UnitsAmperesPerMeter                 EngineeringUnits = 0xA7
	UnitsAmperesPerSquareMeter           EngineeringUnits = 0xA8
	UnitsAmpereSquareMeters              EngineeringUnits = 0xA9
	UnitsDecibels                        EngineeringUnits = 0xC7
	UnitsDecibelsMillivolt               EngineeringUnits = 0xC8
	UnitsDecibelsVolt                    EngineeringUnits = 0xC9
	UnitsFarads                          EngineeringUnits = 0xAA
	UnitsHenrys                          EngineeringUnits = 0xAB
	UnitsOhms                            EngineeringUnits = 0x04
	UnitsOhmMeters                       EngineeringUnits = 0xAC
	UnitsMilliohms                       EngineeringUnits = 0x91
	UnitsKilohms                         EngineeringUnits = 0x7A
	UnitsMegohms                         EngineeringUnits = 0x7B
	UnitsMicrosiemens                    EngineeringUnits = 0xBE
	UnitsMillisiemens                    EngineeringUnits = 0xCA
	UnitsSiemens                         EngineeringUnits = 0xAD // 1 mho equals 1 siemens
	UnitsSiemensPerMeter                 EngineeringUnits = 0xAE
	UnitsTeslas                          EngineeringUnits = 0xAF
	UnitsVolts                           EngineeringUnits = 0x05
	UnitsMillivolts                      EngineeringUnits = 0x7C
	UnitsKilovolts                       EngineeringUnits = 0x06
	UnitsMegavolts                       EngineeringUnits = 0x07
	UnitsVoltAmperes                     EngineeringUnits = 0x08
	UnitsKilovoltAmperes                 EngineeringUnits = 0x09
	UnitsMegavoltAmperes                 EngineeringUnits = 0x0A
	UnitsVoltAmperesReactive             EngineeringUnits = 0x0B
	UnitsKilovoltAmperesReactive         EngineeringUnits = 0x0C
	UnitsMegavoltAmperesReactive         EngineeringUnits = 0x0D
	UnitsVoltsPerDegreeKelvin            EngineeringUnits = 0xB0
	UnitsVoltsPerMeter                   EngineeringUnits = 0xB1
	UnitsDegreesPhase                    EngineeringUnits = 0x0E
	UnitsPowerFactor                     EngineeringUnits = 0x0F
	UnitsWebers                          EngineeringUnits = 0xB2
	UnitsJoules                          EngineeringUnits = 0x10
	UnitsKilojoules                      EngineeringUnits = 0x11
	UnitsKilojoulesPerKilogram           EngineeringUnits = 0x7D
	UnitsMegajoules                      EngineeringUnits = 0x7E
	UnitsWattHours                       EngineeringUnits = 0x12
	UnitsKilowattHours                   EngineeringUnits = 0x13
	UnitsMegawattHours                   EngineeringUnits = 0x92
	UnitsWattHoursReactive               EngineeringUnits = 0xCB
	UnitsKilowattHoursReactive           EngineeringUnits = 0xCC
	UnitsMegawattHoursReactive           EngineeringUnits = 0xCD
	UnitsBtus                            EngineeringUnits = 0x14
	UnitsKiloBtus                        EngineeringUnits = 0x93
	UnitsMegaBtus                        EngineeringUnits = 0x94
	UnitsTherms                          EngineeringUnits = 0x15
	UnitsTonHours                        EngineeringUnits = 0x16
	UnitsJoulesPerKilogramDryAir         EngineeringUnits = 0x17
	UnitsKilojoulesPerKilogramDryAir     EngineeringUnits = 0x95
	UnitsMegajoulesPerKilogramDryAir     EngineeringUnits = 0x96
	UnitsBtusPerPoundDryAir              EngineeringUnits = 0x18
	UnitsBtusPerPound                    EngineeringUnits = 0x75
	UnitsJoulesPerDegreeKelvin           EngineeringUnits = 0x7F
	UnitsKilojoulesPerDegreeKelvin       EngineeringUnits = 0x97
	UnitsMegajoulesPerDegreeKelvin       EngineeringUnits = 0x98
	UnitsJoulesPerKilogramDegreeKelvin   EngineeringUnits = 0x80
	UnitsNewton                          EngineeringUnits = 0x99
	UnitsCyclesPerHour                   EngineeringUnits = 0x19
	UnitsCyclesPerMinute                 EngineeringUnits = 0x1A
	UnitsHertz                           EngineeringUnits = 0x1B
	UnitsKilohertz                       EngineeringUnits = 0x81
	UnitsMegahertz                       EngineeringUnits = 0x82
	UnitsPerHour                         EngineeringUnits = 0x83
	UnitsGramsOfWaterPerKilogramDryAir   EngineeringUnits = 0x1C
	UnitsPercentRelativeHumidity         EngineeringUnits = 0x1D
	UnitsMicrometers                     EngineeringUnits = 0xC2
	UnitsMillimeters                     EngineeringUnits = 0x1E
	UnitsCentimeters                     EngineeringUnits = 0x76
	UnitsKilometers                      EngineeringUnits = 0xC1
	UnitsMeters                          EngineeringUnits = 0x1F
	UnitsInches                          EngineeringUnits = 0x20
	UnitsFeet                            EngineeringUnits = 0x21
	UnitsCandelas                        EngineeringUnits = 0xB3
	UnitsCandelasPerSquareMeter          EngineeringUnits = 0xB4
	UnitsWattsPerSquareFoot              EngineeringUnits = 0x22
	UnitsWattsPerSquareMeter             EngineeringUnits = 0x23
	UnitsLumens                          EngineeringUnits = 0x24
	UnitsLuxes                           EngineeringUnits = 0x25
	UnitsFootCandles                     EngineeringUnits = 0x26
	UnitsMilligrams                      EngineeringUnits = 0xC4
	UnitsGrams                           EngineeringUnits = 0xC3
	UnitsKilograms                       EngineeringUnits = 0x27
	UnitsPoundsMass                      EngineeringUnits = 0x28
	UnitsTons                            EngineeringUnits = 0x29
	UnitsGramsPerSecond                  EngineeringUnits = 0x9A
	UnitsGramsPerMinute                  EngineeringUnits = 0x9B
	UnitsKilogramsPerSecond              EngineeringUnits = 0x2A
	UnitsKilogramsPerMinute              EngineeringUnits = 0x2B
	UnitsKilogramsPerHour                EngineeringUnits = 0x2C
	UnitsPoundsMassPerSecond             EngineeringUnits = 0x77
	UnitsPoundsMassPerMinute             EngineeringUnits = 0x2D
	UnitsPoundsMassPerHour               EngineeringUnits = 0x2E
	UnitsTonsPerHour                     EngineeringUnits = 0x9C
	UnitsMilliwatts                      EngineeringUnits = 0x84
	UnitsWatts                           EngineeringUnits = 0x2F
	UnitsKilowatts                       EngineeringUnits = 0x30
	UnitsMegawatts                       EngineeringUnits = 0x31
	UnitsBtusPerHour                     EngineeringUnits = 0x32
	UnitsKiloBtusPerHour                 EngineeringUnits = 0x9D
	UnitsHorsepower                      EngineeringUnits = 0x33
	UnitsTonsRefrigeration               EngineeringUnits = 0x34
	UnitsPascals                         EngineeringUnits = 0x35
	UnitsHectopascals                    EngineeringUnits = 0x85
	UnitsKilopascals                     EngineeringUnits = 0x36
	UnitsMillibars                       EngineeringUnits = 0x86
	UnitsBars                            EngineeringUnits = 0x37
	UnitsPoundsForcePerSquareInch        EngineeringUnits = 0x38
	UnitsMillimetersOfWater              EngineeringUnits = 0xCE
	UnitsCentimetersOfWater              EngineeringUnits = 0x39
	UnitsInchesOfWater                   EngineeringUnits = 0x3A
	UnitsMillimetersOfMercury            EngineeringUnits = 0x3B
	UnitsCentimetersOfMercury            EngineeringUnits = 0x3C
	UnitsInchesOfMercury                 EngineeringUnits = 0x3D
	UnitsDegreesCelsius                  EngineeringUnits = 0x3E
	UnitsDegreesKelvin                   EngineeringUnits = 0x3F
	UnitsDegreesKelvinPerHour            EngineeringUnits = 0xB5
	UnitsDegreesKelvinPerMinute          EngineeringUnits = 0xB6
	UnitsDegreesFahrenheit               EngineeringUnits = 0x40
	UnitsDegreeDaysCelsius               EngineeringUnits = 0x41
	UnitsDegreeDaysFahrenheit            EngineeringUnits = 0x42
	UnitsDeltaDegreesFahrenheit          EngineeringUnits = 0x78
	UnitsDeltaDegreesKelvin              EngineeringUnits = 0x79
	UnitsYears                           EngineeringUnits = 0x43
	UnitsMonths                          EngineeringUnits = 0x44
	UnitsWeeks                           EngineeringUnits = 0x45
	UnitsDays                            EngineeringUnits = 0x46
	UnitsHours                           EngineeringUnits = 0x47
	UnitsMinutes                         EngineeringUnits = 0x48
	UnitsSeconds                         EngineeringUnits = 0x49
	UnitsHundredthsSeconds               EngineeringUnits = 0x9E
	UnitsMilliseconds                    EngineeringUnits = 0x9F
	UnitsNewtonMeters                    EngineeringUnits = 0xA0
	UnitsMillimetersPerSecond            EngineeringUnits = 0xA1
	UnitsMillimetersPerMinute            EngineeringUnits = 0xA2
	UnitsMetersPerSecond                 EngineeringUnits = 0x4A
	UnitsMetersPerMinute                 EngineeringUnits = 0xA3
	UnitsMetersPerHour                   EngineeringUnits = 0xA4
	UnitsKilometersPerHour               EngineeringUnits = 0x4B
	UnitsFeetPerSecond                   EngineeringUnits = 0x4C
	UnitsFeetPerMinute                   EngineeringUnits = 0x4D
	UnitsMilesPerHour                    EngineeringUnits = 0x4E
	UnitsCubicFeet                       EngineeringUnits = 0x4F
	UnitsCubicMeters                     EngineeringUnits = 0x50
	UnitsImperialGallons                 EngineeringUnits = 0x51
	UnitsMilliliters                     EngineeringUnits = 0xC5
	UnitsLiters                          EngineeringUnits = 0x52
	UnitsUSGallons                       EngineeringUnits = 0x53
	UnitsCubicFeetPerSecond              EngineeringUnits = 0x8E
	UnitsCubicFeetPerMinute              EngineeringUnits = 0x54
	UnitsCubicFeetPerHour                EngineeringUnits = 0xBF
	UnitsCubicMetersPerSecond            EngineeringUnits = 0x55
	UnitsCubicMetersPerMinute            EngineeringUnits = 0xA5
	UnitsCubicMetersPerHour              EngineeringUnits = 0x87
	UnitsImperialGallonsPerMinute        EngineeringUnits = 0x56
	UnitsMillilitersPerSecond            EngineeringUnits = 0xC6
	UnitsLitersPerSecond                 EngineeringUnits = 0x57
	UnitsLitersPerMinute                 EngineeringUnits = 0x58
	UnitsLitersPerHour                   EngineeringUnits = 0x88
	UnitsUSGallonsPerMinute              EngineeringUnits = 0x59
	UnitsUSGallonsPerHour                EngineeringUnits = 0xC0
	UnitsDegreesAngular                  EngineeringUnits = 0x5A
	UnitsDegreesCelsiusPerHour           EngineeringUnits = 0x5B
	UnitsDegreesCelsiusPerMinute         EngineeringUnits = 0x5C
	UnitsDegreesFahrenheitPerHour        EngineeringUnits = 0x5D
	UnitsDegreesFahrenheitPerMinute      EngineeringUnits = 0x5E
	UnitsJouleSeconds                    EngineeringUnits = 0xB7
	UnitsKilogramsPerCubicMeter          EngineeringUnits = 0xBA
	UnitsKwHoursPerSquareMeter           EngineeringUnits = 0x89
	UnitsKwHoursPerSquareFoot            EngineeringUnits = 0x8A
	UnitsMegajoulesPerSquareMeter        EngineeringUnits = 0x8B
	UnitsMegajoulesPerSquareFoot         EngineeringUnits = 0x8C
	UnitsNoUnits                         EngineeringUnits = 0x5F
	UnitsNewtonSeconds                   EngineeringUnits = 0xBB
	UnitsNewtonsPerMeter                 EngineeringUnits = 0xBC
	UnitsPartsPerMillion                 EngineeringUnits = 0x60
	UnitsPartsPerBillion                 EngineeringUnits = 0x61
	UnitsPercent                         EngineeringUnits = 0x62
	UnitsPercentObscurationPerFoot       EngineeringUnits = 0x8F
	UnitsPercentObscurationPerMeter      EngineeringUnits = 0x90
	UnitsPercentPerSecond                EngineeringUnits = 0x63
	UnitsPerMinute                       EngineeringUnits = 0x64
	UnitsPerSecond                       EngineeringUnits = 0x65
	UnitsPSIPerDegreeFahrenheit          EngineeringUnits = 0x66
	UnitsRadians                         EngineeringUnits = 0x67
	UnitsRadiansPerSecond                EngineeringUnits = 0xB8
	UnitsRevolutionsPerMinute            EngineeringUnits = 0x68
	UnitsSquareMetersPerNewton           EngineeringUnits = 0xB9
	UnitsWattsPerMeterPerDegreeKelvin    EngineeringUnits = 0xBD
	UnitsWattsPerSquareMeterDegreeKelvin EngineeringUnits = 0x8D
	UnitsPerMille                        EngineeringUnits = 0xCF
	UnitsGramsPerGram                    EngineeringUnits = 0xD0
	UnitsKilogramsPerKilogram            EngineeringUnits = 0xD1
	UnitsGramsPerKilogram                EngineeringUnits = 0xD2
	UnitsMilligramsPerGram               EngineeringUnits = 0xD3
	UnitsMilligramsPerKilogram           EngineeringUnits = 0xD4
	UnitsGramsPerMilliliter              EngineeringUnits = 0xD5
	UnitsGramsPerLiter                   EngineeringUnits = 0xD6
	UnitsMilligramsPerLiter              EngineeringUnits = 0xD7
	UnitsMicrogramsPerLiter              EngineeringUnits = 0xD8
	UnitsGramsPerCubicMeter              EngineeringUnits = 0xD9
	UnitsMilligramsPerCubicMeter         EngineeringUnits = 0xDA
	UnitsMicrogramsPerCubicMeter         EngineeringUnits = 0xDB
	UnitsNanogramsPerCubicMeter          EngineeringUnits = 0xDC
	UnitsGramsPerCubicCentimeter         EngineeringUnits = 0xDD
	UnitsBecquerels                      EngineeringUnits = 0xDE
	UnitsMegabecquerels                  EngineeringUnits = 0xE0
	UnitsGray                            EngineeringUnits = 0xE1
	UnitsMilligray                       EngineeringUnits = 0xE2
	UnitsMicrogray                       EngineeringUnits = 0xE3
	UnitsSieverts                        EngineeringUnits = 0xE4
	UnitsMillisieverts                   EngineeringUnits = 0xE5
	UnitsMicrosieverts                   EngineeringUnits = 0xE6
	UnitsMicrosievertsPerHour            EngineeringUnits = 0xE7
	UnitsDecibelsA                       EngineeringUnits = 0xE8
	UnitsNephelometricTurbidityUnit      EngineeringUnits = 0xE9
	UnitsPh                              EngineeringUnits = 0xEA
	UnitsGramsPerSquareMeter             EngineeringUnits = 0xEB
	UnitsMinutesPerDegreeKelvin          EngineeringUnits = 0xEC

	UnitsProprietaryRangeMin EngineeringUnits = 0x100
	UnitsProprietaryRangeMax EngineeringUnits = 0xFFFF
)

func (u EngineeringUnits) Name() string {
	switch u {
	case UnitsMetersPerSecondPerSecond:
		return "MetersPerSecondPerSecond"
	case UnitsSquareMeters:
		return "SquareMeters"
	case UnitsSquareCentimeters:
		return "SquareCentimeters"
	case UnitsSquareFeet:
		return "SquareFeet"
	case UnitsSquareInches:
		return "SquareInches"
	case UnitsMilliamperes:
		return "Milliamperes"
	case UnitsAmperes:
		return "Amperes"
	case UnitsAmperesPerMeter:
		return "AmperesPerMeter"
	case UnitsAmperesPerSquareMeter:
		return "AmperesPerSquareMeter"
	case UnitsAmpereSquareMeters:
		return "AmpereSquareMeters"
	case UnitsDecibels:
		return "Decibels"
	case UnitsDecibelsMillivolt:
		return "DecibelsMillivolt"
	case UnitsDecibelsVolt:
		return "DecibelsVolt"
	case UnitsFarads:
		return "Farads"
	case UnitsHenrys:
		return "Henrys"
	case UnitsOhms:
		return "Ohms"
	case UnitsOhmMeters:
		return "OhmMeters"
	case UnitsMilliohms:
		return "Milliohms"
	case UnitsKilohms:
		return "Kilohms"
	case UnitsMegohms:
		return "Megohms"
	case UnitsMicrosiemens:
		return "Microsiemens"
	case UnitsMillisiemens:
		return "Millisiemens"
	case UnitsSiemens:
		return "Siemens"
	case UnitsSiemensPerMeter:
		return "SiemensPerMeter"
	case UnitsTeslas:
		return "Teslas"
	case UnitsVolts:
		return "Volts"
	case UnitsMillivolts:
		return "Millivolts"
	case UnitsKilovolts:
		return "Kilovolts"
	case UnitsMegavolts:
		return "Megavolts"
	case UnitsVoltAmperes:
		return "VoltAmperes"
	case UnitsKilovoltAmperes:
		return "KilovoltAmperes"
	case UnitsMegavoltAmperes:
		return "MegavoltAmperes"
	case UnitsVoltAmperesReactive:
		return "VoltAmperesReactive"
	case UnitsKilovoltAmperesReactive:
		return "KilovoltAmperesReactive"
	case UnitsMegavoltAmperesReactive:
		return "MegavoltAmperesReactive"
	case UnitsVoltsPerDegreeKelvin:
		return "VoltsPerDegreeKelvin"
	case UnitsVoltsPerMeter:
		return "VoltsPerMeter"
	case UnitsDegreesPhase:
		return "DegreesPhase"
	case UnitsPowerFactor:
		return "PowerFactor"
	case UnitsWebers:
		return "Webers"
	case UnitsJoules:
		return "Joules"
	case UnitsKilojoules:
		return "Kilojoules"
	case UnitsKilojoulesPerKilogram:
		return "KilojoulesPerKilogram"
	case UnitsMegajoules:
		return "Megajoules"
	case UnitsWattHours:
		return "WattHours"
	case UnitsKilowattHours:
		return "KilowattHours"
	case UnitsMegawattHours:
		return "MegawattHours"
	case UnitsWattHoursReactive:
		return "WattHoursReactive"
	case UnitsKilowattHoursReactive:
		return "KilowattHoursReactive"
	case UnitsMegawattHoursReactive:
		return "MegawattHoursReactive"
	case UnitsBtus:
		return "Btus"
	case UnitsKiloBtus:
		return "KiloBtus"
	case UnitsMegaBtus:
		return "MegaBtus"
	case UnitsTherms:
		return "Therms"
	case UnitsTonHours:
		return "TonHours"
	case UnitsJoulesPerKilogramDryAir:
		return "JoulesPerKilogramDryAir"
	case UnitsKilojoulesPerKilogramDryAir:
		return "KilojoulesPerKilogramDryAir"
	case UnitsMegajoulesPerKilogramDryAir:
		return "MegajoulesPerKilogramDryAir"
	case UnitsBtusPerPoundDryAir:
		return "BtusPerPoundDryAir"
	case UnitsBtusPerPound:
		return "BtusPerPound"
	case UnitsJoulesPerDegreeKelvin:
		return "JoulesPerDegreeKelvin"
	case UnitsKilojoulesPerDegreeKelvin:
		return "KilojoulesPerDegreeKelvin"
	case UnitsMegajoulesPerDegreeKelvin:
		return "MegajoulesPerDegreeKelvin"
	case UnitsJoulesPerKilogramDegreeKelvin:
		return "JoulesPerKilogramDegreeKelvin"
	case UnitsNewton:
		return "Newton"
	case UnitsCyclesPerHour:
		return "CyclesPerHour"
	case UnitsCyclesPerMinute:
		return "CyclesPerMinute"
	case UnitsHertz:
		return "Hertz"
	case UnitsKilohertz:
		return "Kilohertz"
	case UnitsMegahertz:
		return "Megahertz"
	case UnitsPerHour:
		return "PerHour"
	case UnitsGramsOfWaterPerKilogramDryAir:
		return "GramsOfWaterPerKilogramDryAir"
	case UnitsPercentRelativeHumidity:
		return "PercentRelativeHumidity"
	case UnitsMicrometers:
		return "Micrometers"
	case UnitsMillimeters:
		return "Millimeters"
	case UnitsCentimeters:
		return "Centimeters"
	case UnitsKilometers:
		return "Kilometers"
	case UnitsMeters:
		return "Meters"
	case UnitsInches:
		return "Inches"
	case UnitsFeet:
		return "Feet"
	case UnitsCandelas:
		return "Candelas"
	case UnitsCandelasPerSquareMeter:
		return "CandelasPerSquareMeter"
	case UnitsWattsPerSquareFoot:
		return "WattsPerSquareFoot"
	case UnitsWattsPerSquareMeter:
		return "WattsPerSquareMeter"
	case UnitsLumens:
		return "Lumens"
	case UnitsLuxes:
		return "Luxes"
	case UnitsFootCandles:
		return "FootCandles"
	case UnitsMilligrams:
		return "Milligrams"
	case UnitsGrams:
		return "Grams"
	case UnitsKilograms:
		return "Kilograms"
	case UnitsPoundsMass:
		return "PoundsMass"
	case UnitsTons:
		return "Tons"
	case UnitsGramsPerSecond:
		return "GramsPerSecond"
	case UnitsGramsPerMinute:
		return "GramsPerMinute"
	case UnitsKilogramsPerSecond:
		return "KilogramsPerSecond"
	case UnitsKilogramsPerMinute:
		return "KilogramsPerMinute"
	case UnitsKilogramsPerHour:
		return "KilogramsPerHour"
	case UnitsPoundsMassPerSecond:
		return "PoundsMassPerSecond"
	case UnitsPoundsMassPerMinute:
		return "PoundsMassPerMinute"
	case UnitsPoundsMassPerHour:
		return "PoundsMassPerHour"
	case UnitsTonsPerHour:
		return "TonsPerHour"
	case UnitsMilliwatts:
		return "Milliwatts"
	case UnitsWatts:
		return "Watts"
	case UnitsKilowatts:
		return "Kilowatts"
	case UnitsMegawatts:
		return "Megawatts"
	case UnitsBtusPerHour:
		return "BtusPerHour"
	case UnitsKiloBtusPerHour:
		return "KiloBtusPerHour"
	case UnitsHorsepower:
		return "Horsepower"
	case UnitsTonsRefrigeration:
		return "TonsRefrigeration"
	case UnitsPascals:
		return "Pascals"
	case UnitsHectopascals:
		return "Hectopascals"
	case UnitsKilopascals:
		return "Kilopascals"
	case UnitsMillibars:
		return "Millibars"
	case UnitsBars:
		return "Bars"
	case UnitsPoundsForcePerSquareInch:
		return "PoundsForcePerSquareInch"
	case UnitsMillimetersOfWater:
		return "MillimetersOfWater"
	case UnitsCentimetersOfWater:
		return "CentimetersOfWater"
	case UnitsInchesOfWater:
		return "InchesOfWater"
	case UnitsMillimetersOfMercury:
		return "MillimetersOfMercury"
	case UnitsCentimetersOfMercury:
		return "CentimetersOfMercury"
	case UnitsInchesOfMercury:
		return "InchesOfMercury"
	case UnitsDegreesCelsius:
		return "DegreesCelsius"
	case UnitsDegreesKelvin:
		return "DegreesKelvin"
	case UnitsDegreesKelvinPerHour:
		return "DegreesKelvinPerHour"
	case UnitsDegreesKelvinPerMinute:
		return "DegreesKelvinPerMinute"
	case UnitsDegreesFahrenheit:
		return "DegreesFahrenheit"
	case UnitsDegreeDaysCelsius:
		return "DegreeDaysCelsius"
	case UnitsDegreeDaysFahrenheit:
		return "DegreeDaysFahrenheit"
	case UnitsDeltaDegreesFahrenheit:
		return "DeltaDegreesFahrenheit"
	case UnitsDeltaDegreesKelvin:
		return "DeltaDegreesKelvin"
	case UnitsYears:
		return "Years"
	case UnitsMonths:
		return "Months"
	case UnitsWeeks:
		return "Weeks"
	case UnitsDays:
		return "Days"
	case UnitsHours:
		return "Hours"
	case UnitsMinutes:
		return "Minutes"
	case UnitsSeconds:
		return "Seconds"
	case UnitsHundredthsSeconds:
		return "HundredthsSeconds"
	case UnitsMilliseconds:
		return "Milliseconds"
	case UnitsNewtonMeters:
		return "NewtonMeters"
	case UnitsMillimetersPerSecond:
		return "MillimetersPerSecond"
	case UnitsMillimetersPerMinute:
		return "MillimetersPerMinute"
	case UnitsMetersPerSecond:
		return "MetersPerSecond"
	case UnitsMetersPerMinute:
		return "MetersPerMinute"
	case UnitsMetersPerHour:
		return "MetersPerHour"
	case UnitsKilometersPerHour:
		return "KilometersPerHour"
	case UnitsFeetPerSecond:
		return "FeetPerSecond"
	case UnitsFeetPerMinute:
		return "FeetPerMinute"
	case UnitsMilesPerHour:
		return "MilesPerHour"
	case UnitsCubicFeet:
		return "CubicFeet"
	case UnitsCubicMeters:
		return "CubicMeters"
	case UnitsImperialGallons:
		return "ImperialGallons"
	case UnitsMilliliters:
		return "Milliliters"
	case UnitsLiters:
		return "Liters"
	case UnitsUSGallons:
		return "USGallons"
	case UnitsCubicFeetPerSecond:
		return "CubicFeetPerSecond"
	case UnitsCubicFeetPerMinute:
		return "CubicFeetPerMinute"
	case UnitsCubicFeetPerHour:
		return "CubicFeetPerHour"
	case UnitsCubicMetersPerSecond:
		return "CubicMetersPerSecond"
	case UnitsCubicMetersPerMinute:
		return "CubicMetersPerMinute"
	case UnitsCubicMetersPerHour:
		return "CubicMetersPerHour"
	case UnitsImperialGallonsPerMinute:
		return "ImperialGallonsPerMinute"
	case UnitsMillilitersPerSecond:
		return "MillilitersPerSecond"
	case UnitsLitersPerSecond:
		return "LitersPerSecond"
	case UnitsLitersPerMinute:
		return "LitersPerMinute"
	case UnitsLitersPerHour:
		return "LitersPerHour"
	case UnitsUSGallonsPerMinute:
		return "USGallonsPerMinute"
	case UnitsUSGallonsPerHour:
		return "USGallonsPerHour"
	case UnitsDegreesAngular:
		return "DegreesAngular"
	case UnitsDegreesCelsiusPerHour:
		return "DegreesCelsiusPerHour"
	case UnitsDegreesCelsiusPerMinute:
		return "DegreesCelsiusPerMinute"
	case UnitsDegreesFahrenheitPerHour:
		return "DegreesFahrenheitPerHour"
	case UnitsDegreesFahrenheitPerMinute:
		return "DegreesFahrenheitPerMinute"
	case UnitsJouleSeconds:
		return "JouleSeconds"
	case UnitsKilogramsPerCubicMeter:
		return "KilogramsPerCubicMeter"
	case UnitsKwHoursPerSquareMeter:
		return "KwHoursPerSquareMeter"
	case UnitsKwHoursPerSquareFoot:
		return "KwHoursPerSquareFoot"
	case UnitsMegajoulesPerSquareMeter:
		return "MegajoulesPerSquareMeter"
	case UnitsMegajoulesPerSquareFoot:
		return "MegajoulesPerSquareFoot"
	case UnitsNoUnits:
		return "NoUnits"
	case UnitsNewtonSeconds:
		return "NewtonSeconds"
	case UnitsNewtonsPerMeter:
		return "NewtonsPerMeter"
	case UnitsPartsPerMillion:
		return "PartsPerMillion"
	case UnitsPartsPerBillion:
		return "PartsPerBillion"
	case UnitsPercent:
		return "Percent"
	case UnitsPercentObscurationPerFoot:
		return "PercentObscurationPerFoot"
	case UnitsPercentObscurationPerMeter:
		return "PercentObscurationPerMeter"
	case UnitsPercentPerSecond:
		return "PercentPerSecond"
	case UnitsPerMinute:
		return "PerMinute"
	case UnitsPerSecond:
		return "PerSecond"
	case UnitsPSIPerDegreeFahrenheit:
		return "PSIPerDegreeFahrenheit"
	case UnitsRadians:
		return "Radians"
	case UnitsRadiansPerSecond:
		return "RadiansPerSecond"
	case UnitsRevolutionsPerMinute:
		return "RevolutionsPerMinute"
	case UnitsSquareMetersPerNewton:
		return "SquareMetersPerNewton"
	case UnitsWattsPerMeterPerDegreeKelvin:
		return "WattsPerMeterPerDegreeKelvin"
	case UnitsWattsPerSquareMeterDegreeKelvin:
		return "WattsPerSquareMeterDegreeKelvin"
	case UnitsPerMille:
		return "PerMille"
	case UnitsGramsPerGram:
		return "GramsPerGram"
	case UnitsKilogramsPerKilogram:
		return "KilogramsPerKilogram"
	case UnitsGramsPerKilogram:
		return "GramsPerKilogram"
	case UnitsMilligramsPerGram:
		return "MilligramsPerGram"
	case UnitsMilligramsPerKilogram:
		return "MilligramsPerKilogram"
	case UnitsGramsPerMilliliter:
		return "GramsPerMilliliter"
	case UnitsGramsPerLiter:
		return "GramsPerLiter"
	case UnitsMilligramsPerLiter:
		return "MilligramsPerLiter"
	case UnitsMicrogramsPerLiter:
		return "MicrogramsPerLiter"
	case UnitsGramsPerCubicMeter:
		return "GramsPerCubicMeter"
	case UnitsMilligramsPerCubicMeter:
		return "MilligramsPerCubicMeter"
	case UnitsMicrogramsPerCubicMeter:
		return "MicrogramsPerCubicMeter"
	case UnitsNanogramsPerCubicMeter:
		return "NanogramsPerCubicMeter"
	case UnitsGramsPerCubicCentimeter:
		return "GramsPerCubicCentimeter"
	case UnitsBecquerels:
		return "Becquerels"
	case UnitsMegabecquerels:
		return "Megabecquerels"
	case UnitsGray:
		return "Gray"
	case UnitsMilligray:
		return "Milligray"
	case UnitsMicrogray:
		return "Microgray"
	case UnitsSieverts:
		return "Sieverts"
	case UnitsMillisieverts:
		return "Millisieverts"
	case UnitsMicrosieverts:
		return "Microsieverts"
	case UnitsMicrosievertsPerHour:
		return "MicrosievertsPerHour"
	case UnitsDecibelsA:
		return "DecibelsA"
	case UnitsNephelometricTurbidityUnit:
		return "NephelometricTurbidityUnit"
	case UnitsPh:
		return "Ph"
	case UnitsGramsPerSquareMeter:
		return "GramsPerSquareMeter"
	case UnitsMinutesPerDegreeKelvin:
		return "MinutesPerDegreeKelvin"
	}

	if u >= UnitsProprietaryRangeMin {
		return fmt.Sprintf(
			"ProprietaryEngineeringUnits(%d)",
			uint16(u),
		)
	}

	return fmt.Sprintf(
		"EngineeringUnits(%d)",
		uint16(u),
	)
}

func (u EngineeringUnits) String() string {
	switch u {
	case UnitsSquareMeters:
		return "m²"
	case UnitsSquareCentimeters:
		return "cm²"
	case UnitsSquareFeet:
		return "ft²"
	case UnitsSquareInches:
		return "in²"

	case UnitsMilliamperes:
		return "mA"
	case UnitsAmperes:
		return "A"
	case UnitsAmperesPerMeter:
		return "A/m"
	case UnitsAmperesPerSquareMeter:
		return "A/m²"
	case UnitsAmpereSquareMeters:
		return "A·m²"

	case UnitsOhms:
		return "Ω"
	case UnitsMilliohms:
		return "mΩ"
	case UnitsKilohms:
		return "kΩ"
	case UnitsMegohms:
		return "MΩ"

	case UnitsMicrosiemens:
		return "µS"
	case UnitsMillisiemens:
		return "mS"
	case UnitsSiemens:
		return "S"
	case UnitsSiemensPerMeter:
		return "S/m"

	case UnitsVolts:
		return "V"
	case UnitsMillivolts:
		return "mV"
	case UnitsKilovolts:
		return "kV"
	case UnitsMegavolts:
		return "MV"

	case UnitsVoltAmperes:
		return "VA"
	case UnitsKilovoltAmperes:
		return "kVA"
	case UnitsMegavoltAmperes:
		return "MVA"

	case UnitsVoltAmperesReactive:
		return "var"
	case UnitsKilovoltAmperesReactive:
		return "kvar"
	case UnitsMegavoltAmperesReactive:
		return "Mvar"

	case UnitsDegreesPhase:
		return "°"
	case UnitsPowerFactor:
		return "PF"

	case UnitsJoules:
		return "J"
	case UnitsKilojoules:
		return "kJ"
	case UnitsMegajoules:
		return "MJ"

	case UnitsWattHours:
		return "Wh"
	case UnitsKilowattHours:
		return "kWh"
	case UnitsMegawattHours:
		return "MWh"

	case UnitsWattHoursReactive:
		return "varh"
	case UnitsKilowattHoursReactive:
		return "kvarh"
	case UnitsMegawattHoursReactive:
		return "Mvarh"

	case UnitsHertz:
		return "Hz"
	case UnitsKilohertz:
		return "kHz"
	case UnitsMegahertz:
		return "MHz"

	case UnitsMillimeters:
		return "mm"
	case UnitsCentimeters:
		return "cm"
	case UnitsKilometers:
		return "km"
	case UnitsMeters:
		return "m"
	case UnitsInches:
		return "in"
	case UnitsFeet:
		return "ft"

	case UnitsWatts:
		return "W"
	case UnitsMilliwatts:
		return "mW"
	case UnitsKilowatts:
		return "kW"
	case UnitsMegawatts:
		return "MW"

	case UnitsWattsPerSquareFoot:
		return "W/ft²"
	case UnitsWattsPerSquareMeter:
		return "W/m²"

	case UnitsPascals:
		return "Pa"
	case UnitsHectopascals:
		return "hPa"
	case UnitsKilopascals:
		return "kPa"
	case UnitsMillibars:
		return "mbar"
	case UnitsBars:
		return "bar"
	case UnitsPoundsForcePerSquareInch:
		return "psi"

	case UnitsDegreesCelsius:
		return "°C"
	case UnitsDegreesKelvin:
		return "K"
	case UnitsDegreesFahrenheit:
		return "°F"
	case UnitsDeltaDegreesFahrenheit:
		return "Δ°F"
	case UnitsDeltaDegreesKelvin:
		return "ΔK"

	case UnitsYears:
		return "yr"
	case UnitsMonths:
		return "mo"
	case UnitsWeeks:
		return "wk"
	case UnitsDays:
		return "d"
	case UnitsHours:
		return "h"
	case UnitsMinutes:
		return "min"
	case UnitsSeconds:
		return "s"
	case UnitsMilliseconds:
		return "ms"

	case UnitsMillimetersPerSecond:
		return "mm/s"
	case UnitsMillimetersPerMinute:
		return "mm/min"
	case UnitsMetersPerSecond:
		return "m/s"
	case UnitsMetersPerMinute:
		return "m/min"
	case UnitsMetersPerHour:
		return "m/h"
	case UnitsKilometersPerHour:
		return "km/h"
	case UnitsFeetPerSecond:
		return "ft/s"
	case UnitsFeetPerMinute:
		return "ft/min"
	case UnitsMilesPerHour:
		return "mi/h"

	case UnitsCubicFeet:
		return "ft³"
	case UnitsCubicMeters:
		return "m³"
	case UnitsMilliliters:
		return "mL"
	case UnitsLiters:
		return "L"
	case UnitsUSGallons:
		return "gal"

	case UnitsCubicFeetPerSecond:
		return "ft³/s"
	case UnitsCubicFeetPerMinute:
		return "ft³/min"
	case UnitsCubicFeetPerHour:
		return "ft³/h"
	case UnitsCubicMetersPerSecond:
		return "m³/s"
	case UnitsCubicMetersPerMinute:
		return "m³/min"
	case UnitsCubicMetersPerHour:
		return "m³/h"
	case UnitsMillilitersPerSecond:
		return "mL/s"
	case UnitsLitersPerSecond:
		return "L/s"
	case UnitsLitersPerMinute:
		return "L/min"
	case UnitsLitersPerHour:
		return "L/h"
	case UnitsUSGallonsPerMinute:
		return "gal/min"
	case UnitsUSGallonsPerHour:
		return "gal/h"

	case UnitsKilogramsPerCubicMeter:
		return "kg/m³"

	case UnitsNoUnits:
		return ""

	case UnitsPartsPerMillion:
		return "ppm"
	case UnitsPartsPerBillion:
		return "ppb"
	case UnitsPercent:
		return "%"
	case UnitsPercentRelativeHumidity:
		return "%RH"
	case UnitsPercentPerSecond:
		return "%/s"
	case UnitsPerMille:
		return "‰"

	case UnitsRadians:
		return "rad"
	case UnitsRadiansPerSecond:
		return "rad/s"
	case UnitsRevolutionsPerMinute:
		return "rpm"

	case UnitsGramsPerGram:
		return "g/g"
	case UnitsKilogramsPerKilogram:
		return "kg/kg"
	case UnitsGramsPerKilogram:
		return "g/kg"
	case UnitsMilligramsPerGram:
		return "mg/g"
	case UnitsMilligramsPerKilogram:
		return "mg/kg"
	case UnitsGramsPerMilliliter:
		return "g/mL"
	case UnitsGramsPerLiter:
		return "g/L"
	case UnitsMilligramsPerLiter:
		return "mg/L"
	case UnitsMicrogramsPerLiter:
		return "µg/L"
	case UnitsGramsPerCubicMeter:
		return "g/m³"
	case UnitsMilligramsPerCubicMeter:
		return "mg/m³"
	case UnitsMicrogramsPerCubicMeter:
		return "µg/m³"
	case UnitsNanogramsPerCubicMeter:
		return "ng/m³"

	case UnitsWattsPerMeterPerDegreeKelvin:
		return "W/(m·K)"
	case UnitsWattsPerSquareMeterDegreeKelvin:
		return "W/(m²·K)"

	case UnitsMetersPerSecondPerSecond:
		return "m/s²"

	default:
		return u.String()
	}
}
