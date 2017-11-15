package spi

import (
	"reflect"
	"rtos"
	"unsafe"
)

func (d *Driver) isr16() {
	p := d.P
	swp := int(d.swp)
	if d.txn == 0 {
		// New transaction.
		p.StoreTXD(d.txbuf[0^swp])
		p.StoreTXD(d.txbuf[1^swp])
		d.txn = 2
	}
	// SPI can generate events fast (1M event/s for max. speed) so check READY
	// event in loop before return.
	ev := p.Event(READY)
	for ev.IsSet() {
		ev.Clear()
		b := p.LoadRXD()
		if n := len(d.rxbuf); n < cap(d.rxbuf) {
			d.rxbuf = d.rxbuf[:n+1]
			if &d.rxbuf[0] != nil {
				n ^= swp
				d.rxbuf[:n+1][n] = b
			}
		}
		if d.txn < len(d.txbuf) {
			p.StoreTXD(d.txbuf[d.txn^swp])
			d.txn++
		} else {
			switch len(d.rxbuf) {
			case cap(d.rxbuf) - 1:
				// There is still one byte to receive.
				continue
			case cap(d.rxbuf):
				p.NVIC().ClearPending() // Can be edge triggered during ISR.
				d.done.Signal(1)
				return
			}
			p.StoreTXD(d.txbuf[(d.txn-1)^swp^len(d.rxbuf)&1])
		}
	}
}

// AsyncWriteRead16 starts SPI transaction: sending words from out slice
// and receiving words into in slice. It returns immediately without waiting
// for end of transaction. See Wait for more infomation.
func (d *Driver) AsyncWriteRead16(out, in []uint16) {
	txbuf := *(*reflect.StringHeader)(unsafe.Pointer(&out))
	txbuf.Len *= 2
	d.txbuf = *(*string)(unsafe.Pointer(&txbuf))
	d.txn = 0
	rxbuf := *(*reflect.SliceHeader)(unsafe.Pointer(&in))
	rxbuf.Cap = rxbuf.Len * 2
	rxbuf.Len = 0
	d.rxbuf = *(*[]byte)(unsafe.Pointer(&rxbuf))
	if len(out) == 0 {
		if len(in) == 0 {
			d.done.Reset(1)
			return
		}
		d.txbuf = "\xFF\xFF" // Rx-only mode: send 0xFF bytes.
	}
	d.done.Reset(0)
	rtos.IRQ(d.P.NVIC()).Trigger()
}

// WriteRead16 calls AsyncWriteRead16 followed by Wait.
func (d *Driver) WriteRead16(out, in []uint16) int {
	d.AsyncWriteRead16(out, in)
	return d.Wait()
}

// WriteReadWord16 writes and reads 16-bit word.
func (d *Driver) WriteReadWord16(w uint16) uint16 {
	txbuf := reflect.StringHeader{uintptr(unsafe.Pointer(&w)), 2}
	d.txbuf = *(*string)(unsafe.Pointer(&txbuf))
	d.txn = 0
	rxbuf := reflect.SliceHeader{uintptr(unsafe.Pointer(&w)), 0, 2}
	d.rxbuf = *(*[]byte)(unsafe.Pointer(&rxbuf))
	d.done.Reset(0)
	rtos.IRQ(d.P.NVIC()).Trigger()
	d.done.Wait(1, 0)
	return w
}

// AsyncRepeatWord16 starts SPI transaction that sends word n times. It returns
// immediately without waiting for end of transaction. See Wait for more
// infomation.
func (d *Driver) _AsyncRepeatWord16(w uint16, n int) {
	if n == 0 {
		d.done.Reset(1)
		return
	}
	d.rep = w
	txbuf := reflect.StringHeader{uintptr(unsafe.Pointer(&d.rep)), 2}
	d.txbuf = *(*string)(unsafe.Pointer(&txbuf))
	d.txn = 0
	rxbuf := reflect.SliceHeader{0, 0, n * 2} // Discard the received bytes.
	d.rxbuf = *(*[]byte)(unsafe.Pointer(&rxbuf))
	d.done.Reset(0)
	rtos.IRQ(d.P.NVIC()).Trigger()
}

// RepeatByte calls AsyncRepeatWord16 followed by Wait.
func (d *Driver) RepeatWord16(w uint16, n int) {
	d.AsyncRepeatWord16(w, n)
	d.Wait()
	d.isr = (*Driver).isr16
}

func (d *Driver) repeat16() {
	p := d.P
	if d.txn == 0 {
		// New transaction.
		p.StoreTXD(d.txbuf[0])
		p.StoreTXD(d.txbuf[1])
		d.txn = 2
	}
	// SPI can generate events fast (1M event/s for max. speed) so check READY
	// event in loop before return.
	ev := p.Event(READY)
	for ev.IsSet() {
		ev.Clear()
		_ = p.LoadRXD()
		switch {
		case d.txn < len(d.txbuf):
			p.StoreTXD(d.txbuf[d.txn&1])
		case d.txn > len(d.txbuf):
			p.NVIC().ClearPending() // Can be edge triggered during ISR.
			d.done.Signal(1)
			return
		}
		d.txn++
	}
}

func (d *Driver) AsyncRepeatWord16(w uint16, n int) {
	if n == 0 {
		d.done.Reset(1)
		return
	}
	d.rep = w>>8 | w<<8
	txbuf := reflect.StringHeader{uintptr(unsafe.Pointer(&d.rep)), n * 2}
	d.txbuf = *(*string)(unsafe.Pointer(&txbuf))
	d.txn = 0
	d.isr = (*Driver).repeat16
	d.done.Reset(0)
	rtos.IRQ(d.P.NVIC()).Trigger()
}