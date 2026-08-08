package defs

type DeviceStatus byte

// Device Statuses
const (
	StatusOperational         DeviceStatus = 0x00
	StatusOperationalReadOnly DeviceStatus = 0x01
	StatusDownloadRequired    DeviceStatus = 0x02
	StatusDownloadInProgress  DeviceStatus = 0x03
	StatusNonOperational      DeviceStatus = 0x04
	StatusBackupInProgress    DeviceStatus = 0x05
	MaxDeviceStatus           DeviceStatus = 0x06
)

type StatusFlag byte

// Status Flags
const (
	StatusFlagInAlarm      StatusFlag = 0x00
	StatusFlagFault        StatusFlag = 0x01
	StatusFlagOverridden   StatusFlag = 0x02
	StatusFlagOutOfService StatusFlag = 0x03
)
