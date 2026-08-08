package defs

// Maintenance constants.
const (
	MaintenanceNone uint16 = 0x00
	MaintenancePeriodicTest uint16 = 0x01
	MaintenanceNeedServiceOperational uint16 = 0x02
	MaintenanceNeedServiceInoperative uint16 = 0x03
	MaintenanceProprietaryMin uint16 = 0x100
	MaintenanceProprietaryMax uint16 = 0xFFFF
)
