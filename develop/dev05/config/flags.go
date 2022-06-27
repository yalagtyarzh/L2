package config

// Flags - структура, хранящая в себе всевозможные флаги, использующиеся в программе
type Flags struct {
	After      uint
	Before     uint
	Context    uint
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}
