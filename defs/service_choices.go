package defs

type ConfirmedServiceChoice byte

const (
	ServiceConfirmedAcknowledgeAlarm           ConfirmedServiceChoice = 0x00
	ServiceConfirmedCOVNotification            ConfirmedServiceChoice = 0x01
	ServiceConfirmedEventNotification          ConfirmedServiceChoice = 0x02
	ServiceConfirmedGetAlarmSummary            ConfirmedServiceChoice = 0x03
	ServiceConfirmedGetEnrollmentSummary       ConfirmedServiceChoice = 0x04
	ServiceConfirmedSubscribeCOV               ConfirmedServiceChoice = 0x05
	ServiceConfirmedAtomicReadFile             ConfirmedServiceChoice = 0x06
	ServiceConfirmedAtomicWriteFile            ConfirmedServiceChoice = 0x07
	ServiceConfirmedAddListElement             ConfirmedServiceChoice = 0x08
	ServiceConfirmedRemoveListElement          ConfirmedServiceChoice = 0x09
	ServiceConfirmedCreateObject               ConfirmedServiceChoice = 0x0A
	ServiceConfirmedDeleteObject               ConfirmedServiceChoice = 0x0B
	ServiceConfirmedReadProperty               ConfirmedServiceChoice = 0x0C
	ServiceConfirmedReadPropertyConditional    ConfirmedServiceChoice = 0x0D
	ServiceConfirmedReadPropertyMultiple       ConfirmedServiceChoice = 0x0E
	ServiceConfirmedWriteProperty              ConfirmedServiceChoice = 0x0F
	ServiceConfirmedWritePropertyMultiple      ConfirmedServiceChoice = 0x10
	ServiceConfirmedDeviceCommunicationControl ConfirmedServiceChoice = 0x11
	ServiceConfirmedPrivateTransfer            ConfirmedServiceChoice = 0x12
	ServiceConfirmedTextMessage                ConfirmedServiceChoice = 0x13
	ServiceConfirmedReinitializeDevice         ConfirmedServiceChoice = 0x14
	ServiceConfirmedVTOpen                     ConfirmedServiceChoice = 0x15
	ServiceConfirmedVTClose                    ConfirmedServiceChoice = 0x16
	ServiceConfirmedVTData                     ConfirmedServiceChoice = 0x17
	ServiceConfirmedAuthenticate               ConfirmedServiceChoice = 0x18
	ServiceConfirmedRequestKey                 ConfirmedServiceChoice = 0x19
	ServiceConfirmedReadRange                  ConfirmedServiceChoice = 0x1A
	ServiceConfirmedLifeSafetyOperation        ConfirmedServiceChoice = 0x1B
	ServiceConfirmedSubscribeCOVProperty       ConfirmedServiceChoice = 0x1C
	ServiceConfirmedGetEventInformation        ConfirmedServiceChoice = 0x1D
)

type UnconfirmedServiceChoice byte

const (
	ServiceUnconfirmedIAm                    UnconfirmedServiceChoice = 0x00
	ServiceUnconfirmedIHave                  UnconfirmedServiceChoice = 0x01
	ServiceUnconfirmedCOVNotification        UnconfirmedServiceChoice = 0x02
	ServiceUnconfirmedEventNotification      UnconfirmedServiceChoice = 0x03
	ServiceUnconfirmedPrivateTransfer        UnconfirmedServiceChoice = 0x04
	ServiceUnconfirmedTextMessage            UnconfirmedServiceChoice = 0x05
	ServiceUnconfirmedTimeSynchronization    UnconfirmedServiceChoice = 0x06
	ServiceUnconfirmedWhoHas                 UnconfirmedServiceChoice = 0x07
	ServiceUnconfirmedWhoIs                  UnconfirmedServiceChoice = 0x08
	ServiceUnconfirmedUTCTimeSynchronization UnconfirmedServiceChoice = 0x09
	ServiceUnconfirmedWriteGroup             UnconfirmedServiceChoice = 0x0A
)

// ServicesSupportedBit identifies a bit position in the BACnet
// Protocol_Services_Supported bit string.
//
// These are not APDU service-choice values.
type ServicesSupportedBit byte

const (
	ServiceSupportedAcknowledgeAlarm             ServicesSupportedBit = 0x00
	ServiceSupportedConfirmedCOVNotification     ServicesSupportedBit = 0x01
	ServiceSupportedConfirmedEventNotification   ServicesSupportedBit = 0x02
	ServiceSupportedGetAlarmSummary              ServicesSupportedBit = 0x03
	ServiceSupportedGetEnrollmentSummary         ServicesSupportedBit = 0x04
	ServiceSupportedSubscribeCOV                 ServicesSupportedBit = 0x05
	ServiceSupportedAtomicReadFile               ServicesSupportedBit = 0x06
	ServiceSupportedAtomicWriteFile              ServicesSupportedBit = 0x07
	ServiceSupportedAddListElement               ServicesSupportedBit = 0x08
	ServiceSupportedRemoveListElement            ServicesSupportedBit = 0x09
	ServiceSupportedCreateObject                 ServicesSupportedBit = 0x0A
	ServiceSupportedDeleteObject                 ServicesSupportedBit = 0x0B
	ServiceSupportedReadProperty                 ServicesSupportedBit = 0x0C
	ServiceSupportedReadPropertyConditional      ServicesSupportedBit = 0x0D
	ServiceSupportedReadPropertyMultiple         ServicesSupportedBit = 0x0E
	ServiceSupportedWriteProperty                ServicesSupportedBit = 0x0F
	ServiceSupportedWritePropertyMultiple        ServicesSupportedBit = 0x10
	ServiceSupportedDeviceCommunicationControl   ServicesSupportedBit = 0x11
	ServiceSupportedPrivateTransfer              ServicesSupportedBit = 0x12
	ServiceSupportedTextMessage                  ServicesSupportedBit = 0x13
	ServiceSupportedReinitializeDevice           ServicesSupportedBit = 0x14
	ServiceSupportedVTOpen                       ServicesSupportedBit = 0x15
	ServiceSupportedVTClose                      ServicesSupportedBit = 0x16
	ServiceSupportedVTData                       ServicesSupportedBit = 0x17
	ServiceSupportedAuthenticate                 ServicesSupportedBit = 0x18
	ServiceSupportedRequestKey                   ServicesSupportedBit = 0x19
	ServiceSupportedIAm                          ServicesSupportedBit = 0x1A
	ServiceSupportedIHave                        ServicesSupportedBit = 0x1B
	ServiceSupportedUnconfirmedCOVNotification   ServicesSupportedBit = 0x1C
	ServiceSupportedUnconfirmedEventNotification ServicesSupportedBit = 0x1D
	ServiceSupportedUnconfirmedPrivateTransfer   ServicesSupportedBit = 0x1E
	ServiceSupportedUnconfirmedTextMessage       ServicesSupportedBit = 0x1F
	ServiceSupportedTimeSynchronization          ServicesSupportedBit = 0x20
	ServiceSupportedWhoHas                       ServicesSupportedBit = 0x21
	ServiceSupportedWhoIs                        ServicesSupportedBit = 0x22
	ServiceSupportedReadRange                    ServicesSupportedBit = 0x23
	ServiceSupportedUTCTimeSynchronization       ServicesSupportedBit = 0x24
	ServiceSupportedLifeSafetyOperation          ServicesSupportedBit = 0x25
	ServiceSupportedSubscribeCOVProperty         ServicesSupportedBit = 0x26
	ServiceSupportedGetEventInformation          ServicesSupportedBit = 0x27
	ServiceSupportedWriteGroup                   ServicesSupportedBit = 0x28
)
