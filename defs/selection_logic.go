package defs

type SelectionLogic byte

// Selection logic constants.
const (
	SelectionLogicAnd SelectionLogic = 0x00
	SelectionLogicOr  SelectionLogic = 0x01
	SelectionLogicAll SelectionLogic = 0x02
)
