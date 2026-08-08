package objects

import "github.com/wz2b/bacnet/bactypes"

type Device struct {
	bactypes.DefaultObject
	SystemStatus                     string                      `json:"systemStatus"`
	VendorName                       string                      `json:"vendorName"`
	VendorIdentifier                 string                      `json:"vendorIdentifier"`
	ModelName                        string                      `json:"modelName"`
	FirmwareRevision                 string                      `json:"firmwareRevision"`
	ApplicationSoftwareVersion       string                      `json:"applicationSoftwareVersion"`
	ProtocolVersion                  string                      `json:"protocolVersion"`
	ProtocolRevision                 string                      `json:"protocolRevision"`
	ProtocolServicesSupported        []string                    `json:"protocolServicesSupported"`
	ProtocolObjectTypesSupported     string                      `json:"protocolObjectTypesSupported"`
	ObjectList                       map[string]*bactypes.Object `json:"objectList"`
	MaxAPDULengthAccepted            string                      `json:"maxAPDULengthAccepted"`
	SegmentationSupported            bool                        `json:"segmentationSupported"`
	APDUTimeout                      int32                       `json:"APDUTimeout"`
	NumberOfAPDURetries              int                         `json:"numberOfAPDURetries"`
	DeviceAddressBinding             string                      `json:"deviceAddressBinding"`
	DatabaseRevision                 string                      `json:"databaseRevision"`
	Location                         string                      `json:"location"`
	Description                      string                      `json:"description"`
	StructuredObjectList             string                      `json:"structuredObjectList"`
	MaxSegmentsAccepted              string                      `json:"maxSegmentsAccepted"`
	VTClassesSupported               string                      `json:"VTClassesSupported"`
	ActiveVTSessions                 string                      `json:"activeVTSessions"`
	LocalTime                        string                      `json:"localTime"`
	LocalDate                        string                      `json:"localDate"`
	UTCOffset                        int                         `json:"UTCOffset"`
	DaylightSavingsStatus            bool                        `json:"daylightSavingsStatus"`
	APDUSegmentTimeout               int32                       `json:"APDUSegmentTimeout"`
	TimeSynchronizationRecipients    string                      `json:"timeSynchronizationRecipients"`
	MaxMaster                        string                      `json:"maxMaster"`
	MaxInfoFrames                    string                      `json:"maxInfoFrames"`
	ConfigurationFiles               string                      `json:"configurationFiles"`
	LastRestoreTime                  string                      `json:"lastRestoreTime"`
	BackupFailureTimeout             string                      `json:"backupFailureTimeout"`
	BackupPreparationTime            string                      `json:"backupPreparationTime"`
	RestorePreparationTime           string                      `json:"restorePreparationTime"`
	RestoreCompletionTime            string                      `json:"restoreCompletionTime"`
	BackupAndRestoreState            string                      `json:"backupAndRestoreState"`
	ActiveCovSubscriptions           string                      `json:"activeCovSubscriptions"`
	SlaveProxyEnable                 string                      `json:"slaveProxyEnable"`
	ManualSlaveAddressBinding        string                      `json:"manualSlaveAddressBinding"`
	AutoSlaveDiscovery               string                      `json:"autoSlaveDiscovery"`
	SlaveAddressBinding              string                      `json:"slaveAddressBinding"`
	LastRestartReason                string                      `json:"lastRestartReason"`
	TimeOfDeviceRestart              string                      `json:"timeOfDeviceRestart"`
	RestartNotificationRecipients    string                      `json:"restartNotificationRecipients"`
	UTCTimeSynchronizationRecipients string                      `json:"UTCTimeSynchronizationRecipients"`
	TimeSynchronizationInterval      string                      `json:"timeSynchronizationInterval"`
	AlignIntervals                   string                      `json:"alignIntervals"`
	IntervalOffset                   string                      `json:"intervalOffset"`
	ProfileName                      string                      `json:"profileName"`
}
