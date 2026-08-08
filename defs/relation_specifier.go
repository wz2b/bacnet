package defs

type RelationSpecifier byte

// Relation specifier constants.
const (
	RelationSpecifierEqual              RelationSpecifier = 0x00
	RelationSpecifierNotEqual           RelationSpecifier = 0x01
	RelationSpecifierLessThan           RelationSpecifier = 0x02
	RelationSpecifierGreaterThan        RelationSpecifier = 0x03
	RelationSpecifierLessThanOrEqual    RelationSpecifier = 0x04
	RelationSpecifierGreaterThanOrEqual RelationSpecifier = 0x05
)
