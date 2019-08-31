package main

/*
Chapter 3
Create storage
*/

type dataFlipFlop struct {
	currentValue bool
}

func (d *dataFlipFlop) tick(input bool) {
	d.currentValue = input
}

type bit struct {
	dff dataFlipFlop
	out bool
}

func (b *bit) tick(in bool, load bool) {
	b.dff.tick(mux(b.dff.currentValue, in, load))
	b.out = b.dff.currentValue
}

type register struct {
	bits []bit
	out  []bool
}

func (r *register) tick(in []bool, load bool) {
	for index := 0; index < len(r.bits); index++ {
		r.bits[index].tick(in[index], load)
		r.out[index] = r.bits[index].out
	}
}

func getRegister(registerSize int) register {
	bits := make([]bit, registerSize)
	for j := 0; j < registerSize; j++ {
		bits[j] = bit{dff: dataFlipFlop{}}
	}
	return register{bits: bits, out: make([]bool, registerSize)}
}

type ram interface {
	tick(in []bool, address []bool, load bool)
	read(address []bool) []bool
}

type ram8 struct {
	registers []register
}

func (r *ram8) tick(in []bool, address []bool, load bool) {
	a0, a1, a2, a3, a4, a5, a6, a7 := dmux8Way(load, address)
	r.registers[0].tick(in, a0)
	r.registers[1].tick(in, a1)
	r.registers[2].tick(in, a2)
	r.registers[3].tick(in, a3)
	r.registers[4].tick(in, a4)
	r.registers[5].tick(in, a5)
	r.registers[6].tick(in, a6)
	r.registers[7].tick(in, a7)
}

func (r *ram8) read(address []bool) []bool {
	return mux8WayMulti(
		r.registers[0].out,
		r.registers[1].out,
		r.registers[2].out,
		r.registers[3].out,
		r.registers[4].out,
		r.registers[5].out,
		r.registers[6].out,
		r.registers[7].out,
		address)
}

func getRAM8(registerSize int) ram8 {
	registers := make([]register, 8)
	for i := 0; i < 8; i++ {
		registers[i] = getRegister(registerSize)
	}
	return ram8{registers: registers}
}

type ram64 struct {
	ram8s []ram8
}

func (r *ram64) tick(in []bool, address []bool, load bool) {
	a0, a1, a2, a3, a4, a5, a6, a7 := dmux8Way(load, address[:3])
	r.ram8s[0].tick(in, address[3:], a0)
	r.ram8s[1].tick(in, address[3:], a1)
	r.ram8s[2].tick(in, address[3:], a2)
	r.ram8s[3].tick(in, address[3:], a3)
	r.ram8s[4].tick(in, address[3:], a4)
	r.ram8s[5].tick(in, address[3:], a5)
	r.ram8s[6].tick(in, address[3:], a6)
	r.ram8s[7].tick(in, address[3:], a7)
}

func (r *ram64) read(address []bool) []bool {
	return mux8WayMulti(
		r.ram8s[0].read(address[3:]),
		r.ram8s[1].read(address[3:]),
		r.ram8s[2].read(address[3:]),
		r.ram8s[3].read(address[3:]),
		r.ram8s[4].read(address[3:]),
		r.ram8s[5].read(address[3:]),
		r.ram8s[6].read(address[3:]),
		r.ram8s[7].read(address[3:]),
		address[:3])
}

func getRAM64(registerSize int) ram64 {
	ram8s := make([]ram8, 8)
	for i := 0; i < 8; i++ {
		ram8s[i] = getRAM8(registerSize)
	}
	return ram64{ram8s: ram8s}
}

type ram512 struct {
	ram64s []ram64
}

func (r *ram512) tick(in []bool, address []bool, load bool) {
	a0, a1, a2, a3, a4, a5, a6, a7 := dmux8Way(load, address[:3])
	r.ram64s[0].tick(in, address[3:], a0)
	r.ram64s[1].tick(in, address[3:], a1)
	r.ram64s[2].tick(in, address[3:], a2)
	r.ram64s[3].tick(in, address[3:], a3)
	r.ram64s[4].tick(in, address[3:], a4)
	r.ram64s[5].tick(in, address[3:], a5)
	r.ram64s[6].tick(in, address[3:], a6)
	r.ram64s[7].tick(in, address[3:], a7)
}

func (r *ram512) read(address []bool) []bool {
	return mux8WayMulti(
		r.ram64s[0].read(address[3:]),
		r.ram64s[1].read(address[3:]),
		r.ram64s[2].read(address[3:]),
		r.ram64s[3].read(address[3:]),
		r.ram64s[4].read(address[3:]),
		r.ram64s[5].read(address[3:]),
		r.ram64s[6].read(address[3:]),
		r.ram64s[7].read(address[3:]),
		address[:3])
}

func getRAM512(registerSize int) ram512 {
	ram64s := make([]ram64, 8)
	for i := 0; i < 8; i++ {
		ram64s[i] = getRAM64(registerSize)
	}
	return ram512{ram64s: ram64s}
}

type ram4k struct {
	ram512s []ram512
}

func (r *ram4k) tick(in []bool, address []bool, load bool) {
	a0, a1, a2, a3, a4, a5, a6, a7 := dmux8Way(load, address[:3])
	r.ram512s[0].tick(in, address[3:], a0)
	r.ram512s[1].tick(in, address[3:], a1)
	r.ram512s[2].tick(in, address[3:], a2)
	r.ram512s[3].tick(in, address[3:], a3)
	r.ram512s[4].tick(in, address[3:], a4)
	r.ram512s[5].tick(in, address[3:], a5)
	r.ram512s[6].tick(in, address[3:], a6)
	r.ram512s[7].tick(in, address[3:], a7)
}

func (r *ram4k) read(address []bool) []bool {
	return mux8WayMulti(
		r.ram512s[0].read(address[3:]),
		r.ram512s[1].read(address[3:]),
		r.ram512s[2].read(address[3:]),
		r.ram512s[3].read(address[3:]),
		r.ram512s[4].read(address[3:]),
		r.ram512s[5].read(address[3:]),
		r.ram512s[6].read(address[3:]),
		r.ram512s[7].read(address[3:]),
		address[:3])
}

func getRAM4K(registerSize int) ram4k {
	ram512s := make([]ram512, 8)
	for i := 0; i < 8; i++ {
		ram512s[i] = getRAM512(registerSize)
	}
	return ram4k{ram512s: ram512s}
}

type ram16k struct {
	ram4ks []ram4k
}

func (r *ram16k) tick(in []bool, address []bool, load bool) {
	a0, a1, a2, a3 := dmux4Way(load, address[:2])
	r.ram4ks[0].tick(in, address[2:], a0)
	r.ram4ks[1].tick(in, address[2:], a1)
	r.ram4ks[2].tick(in, address[2:], a2)
	r.ram4ks[3].tick(in, address[2:], a3)
}

func (r *ram16k) read(address []bool) []bool {
	return mux4WayMulti(
		r.ram4ks[0].read(address[2:]),
		r.ram4ks[1].read(address[2:]),
		r.ram4ks[2].read(address[2:]),
		r.ram4ks[3].read(address[2:]),
		address[:2])
}

func getRAM16K(registerSize int) ram16k {
	ram4ks := make([]ram4k, 4)
	for i := 0; i < 4; i++ {
		ram4ks[i] = getRAM4K(registerSize)
	}
	return ram16k{ram4ks: ram4ks}
}

type counter struct {
	zero  []bool
	count register
}

func (c *counter) tick(in []bool, inc bool, load bool, reset bool) {
	result2 := muxMulti(c.count.out, increment(c.count.out), inc)
	result3 := muxMulti(result2, in, load)
	result4 := muxMulti(result3, c.zero, reset)
	c.count.tick(result4, true)
}

func (c *counter) read() []bool {
	return c.count.out
}

func getCounter(registerSize int) counter {
	return counter{zero: make([]bool, registerSize-1), count: getRegister(registerSize - 1)}
}
