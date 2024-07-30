package play

import "github.com/zeppelinmc/zeppelin/net/io"

// clientbound
const PacketIdPlayerAbilitiesServerbound = 0x23

type PlayerAbilitiesServerbound struct {
	Flags int8
}

func (PlayerAbilitiesServerbound) ID() int32 {
	return PacketIdPlayerAbilitiesServerbound
}

func (a *PlayerAbilitiesServerbound) Encode(w io.Writer) error {
	return w.Byte(a.Flags)
}

func (a *PlayerAbilitiesServerbound) Decode(r io.Reader) error {
	return r.Byte(&a.Flags)
}
