package defs

type BVLCFunction byte

const (
	BVLCFunctionResult                            BVLCFunction = 0x00
	BVLCFunctionWriteBroadcastDistributionTable   BVLCFunction = 0x01
	BVLCFunctionReadBroadcastDistributionTable    BVLCFunction = 0x02
	BVLCFunctionReadBroadcastDistributionTableACK BVLCFunction = 0x03
	BVLCFunctionForwardedNPDU                     BVLCFunction = 0x04
	BVLCFunctionRegisterForeignDevice             BVLCFunction = 0x05
	BVLCFunctionReadForeignDeviceTable            BVLCFunction = 0x06
	BVLCFunctionReadForeignDeviceTableACK         BVLCFunction = 0x07
	BVLCFunctionDeleteForeignDeviceTableEntry     BVLCFunction = 0x08
	BVLCFunctionDistributeBroadcastToNetwork      BVLCFunction = 0x09
	BVLCFunctionOriginalUnicastNPDU               BVLCFunction = 0x0A
	BVLCFunctionOriginalBroadcastNPDU             BVLCFunction = 0x0B
)

// Bvlc constants.
const (
	BVLCResultSuccessfulCompletion               byte = 0x00
	BVLCResultWriteBroadcastDistributionTableNAK byte = 0x10
	BVLCResultReadBroadcastDistributionTableNAK  byte = 0x20
	BVLCResultReadForeignDeviceTableNAK          byte = 0x40
	BVLCResultDeleteForeignDeviceTableEntryNAK   byte = 0x50
	BVLCResultDistributeBroadcastToNetworkNAK    byte = 0x60
)
