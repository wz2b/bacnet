package defs

type NodeType byte

// Node type constants.
const (
	NodeUnknown        NodeType = 0x00
	NodeSystem         NodeType = 0x01
	NodeNetwork        NodeType = 0x02
	NodeDevice         NodeType = 0x03
	NodeOrganizational NodeType = 0x04
	NodeArea           NodeType = 0x05
	NodeEquipment      NodeType = 0x06
	NodePoint          NodeType = 0x07
	NodeCollection     NodeType = 0x08
	NodeProperty       NodeType = 0x09
	NodeFunctional     NodeType = 0x0A
	NodeOther          NodeType = 0x0B
)
