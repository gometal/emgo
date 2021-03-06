package spi

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f030x6/mmap"
)

type SPI_Periph struct {
	CR1     RCR1
	CR2     RCR2
	SR      RSR
	DR      RDR
	CRCPR   RCRCPR
	RXCRCR  RRXCRCR
	TXCRCR  RTXCRCR
	I2SCFGR RI2SCFGR
}

func (p *SPI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var SPI1 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI1_BASE)))

type CR1 uint32

func (b CR1) Field(mask CR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1) J(v int) CR1 {
	return CR1(bits.MakeField32(v, uint32(mask)))
}

type RCR1 struct{ mmio.U32 }

func (r *RCR1) Bits(mask CR1) CR1     { return CR1(r.U32.Bits(uint32(mask))) }
func (r *RCR1) StoreBits(mask, b CR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR1) SetBits(mask CR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR1) ClearBits(mask CR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR1) Load() CR1             { return CR1(r.U32.Load()) }
func (r *RCR1) Store(b CR1)           { r.U32.Store(uint32(b)) }

func (r *RCR1) AtomicStoreBits(mask, b CR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCR1) AtomicSetBits(mask CR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCR1) AtomicClearBits(mask CR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCR1 struct{ mmio.UM32 }

func (rm RMCR1) Load() CR1   { return CR1(rm.UM32.Load()) }
func (rm RMCR1) Store(b CR1) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) CPHA() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CPHA)}}
}

func (p *SPI_Periph) CPOL() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CPOL)}}
}

func (p *SPI_Periph) MSTR() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(MSTR)}}
}

func (p *SPI_Periph) BR() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(BR)}}
}

func (p *SPI_Periph) SPE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SPE)}}
}

func (p *SPI_Periph) LSBFIRST() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(LSBFIRST)}}
}

func (p *SPI_Periph) SSI() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SSI)}}
}

func (p *SPI_Periph) SSM() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SSM)}}
}

func (p *SPI_Periph) RXONLY() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(RXONLY)}}
}

func (p *SPI_Periph) CRCL() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CRCL)}}
}

func (p *SPI_Periph) CRCNEXT() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CRCNEXT)}}
}

func (p *SPI_Periph) CRCEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(CRCEN)}}
}

func (p *SPI_Periph) BIDIOE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(BIDIOE)}}
}

func (p *SPI_Periph) BIDIMODE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(BIDIMODE)}}
}

type CR2 uint32

func (b CR2) Field(mask CR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2) J(v int) CR2 {
	return CR2(bits.MakeField32(v, uint32(mask)))
}

type RCR2 struct{ mmio.U32 }

func (r *RCR2) Bits(mask CR2) CR2     { return CR2(r.U32.Bits(uint32(mask))) }
func (r *RCR2) StoreBits(mask, b CR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR2) SetBits(mask CR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR2) ClearBits(mask CR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR2) Load() CR2             { return CR2(r.U32.Load()) }
func (r *RCR2) Store(b CR2)           { r.U32.Store(uint32(b)) }

func (r *RCR2) AtomicStoreBits(mask, b CR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCR2) AtomicSetBits(mask CR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCR2) AtomicClearBits(mask CR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCR2 struct{ mmio.UM32 }

func (rm RMCR2) Load() CR2   { return CR2(rm.UM32.Load()) }
func (rm RMCR2) Store(b CR2) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) RXDMAEN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(RXDMAEN)}}
}

func (p *SPI_Periph) TXDMAEN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(TXDMAEN)}}
}

func (p *SPI_Periph) SSOE() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(SSOE)}}
}

func (p *SPI_Periph) NSSP() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(NSSP)}}
}

func (p *SPI_Periph) FRF() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(FRF)}}
}

func (p *SPI_Periph) ERRIE() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(ERRIE)}}
}

func (p *SPI_Periph) RXNEIE() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(RXNEIE)}}
}

func (p *SPI_Periph) TXEIE() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(TXEIE)}}
}

func (p *SPI_Periph) DS() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(DS)}}
}

func (p *SPI_Periph) FRXTH() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(FRXTH)}}
}

func (p *SPI_Periph) LDMARX() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(LDMARX)}}
}

func (p *SPI_Periph) LDMATX() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(LDMATX)}}
}

type SR uint32

func (b SR) Field(mask SR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR) J(v int) SR {
	return SR(bits.MakeField32(v, uint32(mask)))
}

type RSR struct{ mmio.U32 }

func (r *RSR) Bits(mask SR) SR      { return SR(r.U32.Bits(uint32(mask))) }
func (r *RSR) StoreBits(mask, b SR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR) SetBits(mask SR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR) ClearBits(mask SR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR) Load() SR             { return SR(r.U32.Load()) }
func (r *RSR) Store(b SR)           { r.U32.Store(uint32(b)) }

func (r *RSR) AtomicStoreBits(mask, b SR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSR) AtomicSetBits(mask SR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSR) AtomicClearBits(mask SR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSR struct{ mmio.UM32 }

func (rm RMSR) Load() SR   { return SR(rm.UM32.Load()) }
func (rm RMSR) Store(b SR) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) RXNE() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(RXNE)}}
}

func (p *SPI_Periph) TXE() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(TXE)}}
}

func (p *SPI_Periph) CRCERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(CRCERR)}}
}

func (p *SPI_Periph) MODF() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(MODF)}}
}

func (p *SPI_Periph) OVR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(OVR)}}
}

func (p *SPI_Periph) BSY() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(BSY)}}
}

func (p *SPI_Periph) FRE() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(FRE)}}
}

func (p *SPI_Periph) FRLVL() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(FRLVL)}}
}

func (p *SPI_Periph) FTLVL() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(FTLVL)}}
}

type DR uint32

func (b DR) Field(mask DR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR) J(v int) DR {
	return DR(bits.MakeField32(v, uint32(mask)))
}

type RDR struct{ mmio.U32 }

func (r *RDR) Bits(mask DR) DR      { return DR(r.U32.Bits(uint32(mask))) }
func (r *RDR) StoreBits(mask, b DR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDR) SetBits(mask DR)      { r.U32.SetBits(uint32(mask)) }
func (r *RDR) ClearBits(mask DR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDR) Load() DR             { return DR(r.U32.Load()) }
func (r *RDR) Store(b DR)           { r.U32.Store(uint32(b)) }

func (r *RDR) AtomicStoreBits(mask, b DR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDR) AtomicSetBits(mask DR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDR) AtomicClearBits(mask DR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDR struct{ mmio.UM32 }

func (rm RMDR) Load() DR   { return DR(rm.UM32.Load()) }
func (rm RMDR) Store(b DR) { rm.UM32.Store(uint32(b)) }

type CRCPR uint32

func (b CRCPR) Field(mask CRCPR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CRCPR) J(v int) CRCPR {
	return CRCPR(bits.MakeField32(v, uint32(mask)))
}

type RCRCPR struct{ mmio.U32 }

func (r *RCRCPR) Bits(mask CRCPR) CRCPR   { return CRCPR(r.U32.Bits(uint32(mask))) }
func (r *RCRCPR) StoreBits(mask, b CRCPR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCRCPR) SetBits(mask CRCPR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCRCPR) ClearBits(mask CRCPR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCRCPR) Load() CRCPR             { return CRCPR(r.U32.Load()) }
func (r *RCRCPR) Store(b CRCPR)           { r.U32.Store(uint32(b)) }

func (r *RCRCPR) AtomicStoreBits(mask, b CRCPR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCRCPR) AtomicSetBits(mask CRCPR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCRCPR) AtomicClearBits(mask CRCPR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCRCPR struct{ mmio.UM32 }

func (rm RMCRCPR) Load() CRCPR   { return CRCPR(rm.UM32.Load()) }
func (rm RMCRCPR) Store(b CRCPR) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) CRCPOLY() RMCRCPR {
	return RMCRCPR{mmio.UM32{&p.CRCPR.U32, uint32(CRCPOLY)}}
}

type RXCRCR uint32

func (b RXCRCR) Field(mask RXCRCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RXCRCR) J(v int) RXCRCR {
	return RXCRCR(bits.MakeField32(v, uint32(mask)))
}

type RRXCRCR struct{ mmio.U32 }

func (r *RRXCRCR) Bits(mask RXCRCR) RXCRCR  { return RXCRCR(r.U32.Bits(uint32(mask))) }
func (r *RRXCRCR) StoreBits(mask, b RXCRCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRXCRCR) SetBits(mask RXCRCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRXCRCR) ClearBits(mask RXCRCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRXCRCR) Load() RXCRCR             { return RXCRCR(r.U32.Load()) }
func (r *RRXCRCR) Store(b RXCRCR)           { r.U32.Store(uint32(b)) }

func (r *RRXCRCR) AtomicStoreBits(mask, b RXCRCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRXCRCR) AtomicSetBits(mask RXCRCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRXCRCR) AtomicClearBits(mask RXCRCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRXCRCR struct{ mmio.UM32 }

func (rm RMRXCRCR) Load() RXCRCR   { return RXCRCR(rm.UM32.Load()) }
func (rm RMRXCRCR) Store(b RXCRCR) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) RXCRC() RMRXCRCR {
	return RMRXCRCR{mmio.UM32{&p.RXCRCR.U32, uint32(RXCRC)}}
}

type TXCRCR uint32

func (b TXCRCR) Field(mask TXCRCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TXCRCR) J(v int) TXCRCR {
	return TXCRCR(bits.MakeField32(v, uint32(mask)))
}

type RTXCRCR struct{ mmio.U32 }

func (r *RTXCRCR) Bits(mask TXCRCR) TXCRCR  { return TXCRCR(r.U32.Bits(uint32(mask))) }
func (r *RTXCRCR) StoreBits(mask, b TXCRCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTXCRCR) SetBits(mask TXCRCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTXCRCR) ClearBits(mask TXCRCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTXCRCR) Load() TXCRCR             { return TXCRCR(r.U32.Load()) }
func (r *RTXCRCR) Store(b TXCRCR)           { r.U32.Store(uint32(b)) }

func (r *RTXCRCR) AtomicStoreBits(mask, b TXCRCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTXCRCR) AtomicSetBits(mask TXCRCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTXCRCR) AtomicClearBits(mask TXCRCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTXCRCR struct{ mmio.UM32 }

func (rm RMTXCRCR) Load() TXCRCR   { return TXCRCR(rm.UM32.Load()) }
func (rm RMTXCRCR) Store(b TXCRCR) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) TXCRC() RMTXCRCR {
	return RMTXCRCR{mmio.UM32{&p.TXCRCR.U32, uint32(TXCRC)}}
}

type I2SCFGR uint32

func (b I2SCFGR) Field(mask I2SCFGR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask I2SCFGR) J(v int) I2SCFGR {
	return I2SCFGR(bits.MakeField32(v, uint32(mask)))
}

type RI2SCFGR struct{ mmio.U32 }

func (r *RI2SCFGR) Bits(mask I2SCFGR) I2SCFGR { return I2SCFGR(r.U32.Bits(uint32(mask))) }
func (r *RI2SCFGR) StoreBits(mask, b I2SCFGR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RI2SCFGR) SetBits(mask I2SCFGR)      { r.U32.SetBits(uint32(mask)) }
func (r *RI2SCFGR) ClearBits(mask I2SCFGR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RI2SCFGR) Load() I2SCFGR             { return I2SCFGR(r.U32.Load()) }
func (r *RI2SCFGR) Store(b I2SCFGR)           { r.U32.Store(uint32(b)) }

func (r *RI2SCFGR) AtomicStoreBits(mask, b I2SCFGR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RI2SCFGR) AtomicSetBits(mask I2SCFGR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RI2SCFGR) AtomicClearBits(mask I2SCFGR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMI2SCFGR struct{ mmio.UM32 }

func (rm RMI2SCFGR) Load() I2SCFGR   { return I2SCFGR(rm.UM32.Load()) }
func (rm RMI2SCFGR) Store(b I2SCFGR) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) I2SMOD() RMI2SCFGR {
	return RMI2SCFGR{mmio.UM32{&p.I2SCFGR.U32, uint32(I2SMOD)}}
}
