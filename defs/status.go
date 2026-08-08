package defs

// Device Statuses
const (
	StatusOperational byte = 0x00
	StatusOperationalReadOnly byte = 0x01
	StatusDownloadRequired byte = 0x02
	StatusDownloadInProgress byte = 0x03
	StatusNonOperational byte = 0x04
	StatusBackupInProgress byte = 0x05
	MaxDeviceStatus byte = 0x06
)

// Status Flags
const (
	StatusFlagInAlarm byte = 0x00
	StatusFlagFault byte = 0x01
	StatusFlagOverridden byte = 0x02
	StatusFlagOutOfService byte = 0x03
)
