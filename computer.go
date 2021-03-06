package main

import (
	"bufio"
	"os"
	"strings"
)

type cpu struct {
	dRegister register
	aRegister register
	count     counter
}

/*
instruction
0123 4567 8901 2345
111a cccc ccdd djjj

What is the output
C				a=0  a=1
101010  0
111111  1
111010  -1
001100  D
110000  A		 M
001101  !D
110001  !A   !M
001111  -D
110011  -A   -M
011111  D+1
110111  A+1  M+1
001110  D-1
110010  A-1  M-1?
000010  D+A  D+M
010011  D-A  D-M
000111  A-D  M-D
000000  D&A  D&M
010101  D|A  D|M

Where to store the output
d1	d2	d3
A		D		M[A]

Where to jump next
j1	j2	j3
<0  =0  >0
*/
func (c *cpu) tick(instruction []bool, inM []bool, reset bool) ([]bool, bool, []bool, []bool) {
	// Trigger ALU
	output, isZero, isNegative := alu(
		c.dRegister.out,
		muxMulti(c.aRegister.out, inM, instruction[3]),
		instruction[4],
		instruction[5],
		instruction[6],
		instruction[7],
		instruction[8],
		instruction[9])
	// Store in D register
	shouldLoadD := and(instruction[11], instruction[0])
	c.dRegister.tick(output, shouldLoadD)
	// Store in A register
	shouldLoadA := or(not(instruction[0]), instruction[10])
	c.aRegister.tick(muxMulti(instruction, output, instruction[0]), shouldLoadA)
	// Check if we need to jump
	shouldJump1 := and(isNegative, instruction[13])
	shouldJump2 := and(isZero, instruction[14])
	shouldJump3 := and(and(not(isNegative), not(isZero)), instruction[15])
	shouldJump := and(or(or(shouldJump1, shouldJump2), shouldJump3), instruction[0])
	// Increment program counter, load A register, or reset
	c.count.tick(c.aRegister.out[1:], true, shouldJump, reset)
	return output, instruction[12], c.aRegister.out[1:], c.count.read()
}

type memory struct {
	ram0 ram16k
	ram1 ram16k
}

// Run next tick, takes a 15 bit address
func (m *memory) tick(in []bool, address []bool, load bool) {
	a0, a1 := dmux(load, address[0])
	m.ram0.tick(in, address[1:], a0)
	m.ram1.tick(in, address[1:], a1)
}

// Read an address in memory, takes a 15 bit address
func (m *memory) read(address []bool) []bool {
	return muxMulti(
		m.ram0.read(address[1:]),
		m.ram1.read(address[1:]),
		address[0])
}

func getMemory(bits int) memory {
	return memory{ram0: getRAM16K(bits), ram1: getRAM16K(bits)}
}

type rom32k struct {
	rom0 ram16k
	rom1 ram16k
}

// Load a .hack file into ROM
func (r *rom32k) loadProgram(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	address := strToBool("000000000000000")
	index := 0
	for scanner.Scan() {
		in := scanner.Text()
		r.write(strToBool(strings.TrimSuffix(in, "\n")), address)
		address = increment(address)
		index++
	}
}

// Write to a 15 bit address
func (r *rom32k) write(in []bool, address []bool) {
	a0, a1 := dmux(true, address[0])
	r.rom0.tick(in, address[1:], a0)
	r.rom1.tick(in, address[1:], a1)
}

// Read an address from ROM, takes a 15 bit address
func (r *rom32k) read(address []bool) []bool {
	return muxMulti(
		r.rom0.read(address[1:]),
		r.rom1.read(address[1:]),
		address[0])
}

func getROM(bits int) rom32k {
	return rom32k{rom0: getRAM16K(bits), rom1: getRAM16K(bits)}
}

type computer struct {
	processor cpu
	program   rom32k
	data      memory
}

// Execute the next tick
func (c *computer) tick(reset bool) {
	instruction := c.program.read(c.processor.count.read())
	inM := c.data.read(c.processor.aRegister.out[1:])
	outM, writeM, addressM, _ := c.processor.tick(instruction, inM, reset)
	c.data.tick(outM, addressM, writeM)
}

// Load a .hack file into ROM
func (c *computer) loadProgram(filePath string) {
	c.program.loadProgram(filePath)
}

func getComputer(bits int) computer {
	return computer{
		processor: cpu{
			dRegister: getRegister(bits),
			aRegister: getRegister(bits),
			count:     getCounter(bits)},
		program: getROM(bits),
		data:    getMemory(bits)}
}
