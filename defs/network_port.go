package defs

// Network port constants.
const (
	PortTypeEthernet  byte = 0x00
	PortTypeArcnet    byte = 0x01
	PortTypeMSTP      byte = 0x02
	PortTypePTP       byte = 0x03
	PortTypeLontalk   byte = 0x04
	PortTypeBIP       byte = 0x05
	PortTypeZigbee    byte = 0x06
	PortTypeVirtual   byte = 0x07
	PortTypeNonBACnet byte = 0x08
)

// Network port constants.
const (
	PortQualityUnknown           byte = 0x00
	PortQualityLearned           byte = 0x01
	PortQualityLearnedConfigured byte = 0x02
	PortQualityConfigured        byte = 0x03
)

// Network port constants.
const (
	PortCommandIdle                   byte = 0x00
	PortCommandDiscardChanges         byte = 0x01
	PortCommandRenewFDRegistration    byte = 0x02
	PortCommandRestartSlaveDiscovery  byte = 0x03
	PortCommandRenewDHCP              byte = 0x04
	PortCommandRestartAutonegotiation byte = 0x05
	PortCommandDisconnect             byte = 0x06
	PortCommandRestartPort            byte = 0x07
)
