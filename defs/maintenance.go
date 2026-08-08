package defs

type MaintenanceState uint16

const (
	MaintenanceNone                   MaintenanceState = 0x00
	MaintenancePeriodicTest           MaintenanceState = 0x01
	MaintenanceNeedServiceOperational MaintenanceState = 0x02
	MaintenanceNeedServiceInoperative MaintenanceState = 0x03

	MaintenanceProprietaryMin MaintenanceState = 0x100
	MaintenanceProprietaryMax MaintenanceState = 0xFFFF
)
