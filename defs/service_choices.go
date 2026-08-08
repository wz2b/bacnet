package defs

// Service choices constants.
const (
	ServiceConfirmedAcknowledgeAlarm byte = 0x00
	ServiceConfirmedCOVNotification byte = 0x01
	ServiceConfirmedEventNotification byte = 0x02
	ServiceConfirmedGetAlarmSummary byte = 0x03
	ServiceConfirmedGetEnrollmentSummary byte = 0x04
	ServiceConfirmedGetEventInformation byte = 0x1D
	ServiceConfirmedSubscribeCOV byte = 0x05
	ServiceConfirmedSubscribeCOVProperty byte = 0x1C
	ServiceConfirmedLifeSafetyOperation byte = 0x1B
	ServiceConfirmedAtomicReadFile byte = 0x06
	ServiceConfirmedAtomicWriteFile byte = 0x07
	ServiceConfirmedAddListElement byte = 0x08
	ServiceConfirmedRemoveListElement byte = 0x09
	ServiceConfirmedCreateObject byte = 0x0A
	ServiceConfirmedDeleteObject byte = 0x0B
	ServiceConfirmedReadProperty byte = 0x0C
	ServiceConfirmedReadPropertyConditional byte = 0x0D
	ServiceConfirmedReadPropertyMultiple byte = 0x0E
	ServiceConfirmedReadRange byte = 0x1A
	ServiceConfirmedWriteProperty byte = 0x0F
	ServiceConfirmedWritePropertyMultiple byte = 0x10
	ServiceConfirmedDeviceCommunicationControl byte = 0x11
	ServiceConfirmedPrivateTransfer byte = 0x12
	ServiceConfirmedTextMessage byte = 0x13
	ServiceConfirmedReinitializeDevice byte = 0x14
	ServiceConfirmedVTOpen byte = 0x15
	ServiceConfirmedVTClose byte = 0x16
	ServiceConfirmedVTData byte = 0x17
	ServiceConfirmedAuthenticate byte = 0x18
	ServiceConfirmedRequestKey byte = 0x19
	MaxConfirmedService byte = 0x1E
)

// Service choices constants.
const (
	ServiceUnconfirmedIAm byte = 0x00
	ServiceUnconfirmedIHave byte = 0x01
	ServiceUnconfirmedCOVNotification byte = 0x02
	ServiceUnconfirmedEventNotification byte = 0x03
	ServiceUnconfirmedPrivateTransfer byte = 0x04
	ServiceUnconfirmedTextMessage byte = 0x05
	ServiceUnconfirmedTimeSynchronization byte = 0x06
	ServiceUnconfirmedWhoHas byte = 0x07
	ServiceUnconfirmedWhoIs byte = 0x08
	ServiceUnconfirmedUTCTimeSynchronization byte = 0x09
	ServiceUnconfirmedWriteGroup byte = 0x0A
	MaxUnconfirmedService byte = 0x0B
)

// Service choices constants.
const (
	ServiceSupportedAcknowledgeAlarm byte = 0x00
	ServiceSupportedConfirmedCOVNotification byte = 0x01
	ServiceSupportedConfirmedEventNotification byte = 0x02
	ServiceSupportedGetAlarmSummary byte = 0x03
	ServiceSupportedGetEnrollmentSummary byte = 0x04
	ServiceSupportedGetEventInformation byte = 0x27
	ServiceSupportedSubscribeCOV byte = 0x05
	ServiceSupportedSubscribeCOVProperty byte = 0x26
	ServiceSupportedLifeSafetyOperation byte = 0x25
	ServiceSupportedAtomicReadFile byte = 0x06
	ServiceSupportedAtomicWriteFile byte = 0x07
	ServiceSupportedAddListElement byte = 0x08
	ServiceSupportedRemoveListElement byte = 0x09
	ServiceSupportedCreateObject byte = 0x0A
	ServiceSupportedDeleteObject byte = 0x0B
	ServiceSupportedReadProperty byte = 0x0C
	ServiceSupportedReadPropertyConditional byte = 0x0D
	ServiceSupportedReadPropertyMultiple byte = 0x0E
	ServiceSupportedReadRange byte = 0x23
	ServiceSupportedWriteProperty byte = 0x0F
	ServiceSupportedWritePropertyMultiple byte = 0x10
	ServiceSupportedWriteGroup byte = 0x28
	ServiceSupportedDeviceCommunicationControl byte = 0x11
	ServiceSupportedPrivateTransfer byte = 0x12
	ServiceSupportedTextMessage byte = 0x13
	ServiceSupportedReinitializeDevice byte = 0x14
	ServiceSupportedVTOpen byte = 0x15
	ServiceSupportedVTClose byte = 0x16
	ServiceSupportedVTData byte = 0x17
	ServiceSupportedAuthenticate byte = 0x18
	ServiceSupportedRequestKey byte = 0x19
	ServiceSupportedIAm byte = 0x1A
	ServiceSupportedIHave byte = 0x1B
	ServiceSupportedUnconfirmedCOVNotification byte = 0x1C
	ServiceSupportedUnconfirmedEventNotification byte = 0x1D
	ServiceSupportedUnconfirmedPrivateTransfer byte = 0x1E
	ServiceSupportedUnconfirmedTextMessage byte = 0x1F
	ServiceSupportedTimeSynchronization byte = 0x20
	ServiceSupportedUTCTimeSynchronization byte = 0x24
	ServiceSupportedWhoHas byte = 0x21
	ServiceSupportedWhoIs byte = 0x22
)
