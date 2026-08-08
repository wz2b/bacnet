package defs

type NetworkPortType byte

const (
	PortTypeEthernet  NetworkPortType = 0x00
	PortTypeArcnet    NetworkPortType = 0x01
	PortTypeMSTP      NetworkPortType = 0x02
	PortTypePTP       NetworkPortType = 0x03
	PortTypeLontalk   NetworkPortType = 0x04
	PortTypeBIP       NetworkPortType = 0x05
	PortTypeZigbee    NetworkPortType = 0x06
	PortTypeVirtual   NetworkPortType = 0x07
	PortTypeNonBACnet NetworkPortType = 0x08
)

type NetworkPortQuality byte

const (
	PortQualityUnknown           NetworkPortQuality = 0x00
	PortQualityLearned           NetworkPortQuality = 0x01
	PortQualityLearnedConfigured NetworkPortQuality = 0x02
	PortQualityConfigured        NetworkPortQuality = 0x03
)

type NetworkPortCommand byte

const (
	PortCommandIdle                   NetworkPortCommand = 0x00
	PortCommandDiscardChanges         NetworkPortCommand = 0x01
	PortCommandRenewFDRegistration    NetworkPortCommand = 0x02
	PortCommandRestartSlaveDiscovery  NetworkPortCommand = 0x03
	PortCommandRenewDHCP              NetworkPortCommand = 0x04
	PortCommandRestartAutonegotiation NetworkPortCommand = 0x05
	PortCommandDisconnect             NetworkPortCommand = 0x06
	PortCommandRestartPort            NetworkPortCommand = 0x07
)
