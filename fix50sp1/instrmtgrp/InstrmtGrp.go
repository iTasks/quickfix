package instrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
)

//NoRelatedSym is a repeating group in InstrmtGrp
type NoRelatedSym struct {
	//Instrument Component
	instrument.Instrument
}

//InstrmtGrp is a fix50sp1 Component
type InstrmtGrp struct {
	//NoRelatedSym is a non-required field for InstrmtGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *InstrmtGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }