package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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


d1	d2	d3
A		D		M[A]

j1	j2	j3
<0  =0  >0
*/
func (c *cpu) tick(instruction []bool, inM []bool, reset bool) {
	// Trigger ALU
	output, isZero, isNegative := alu(c.dRegister.out, muxMulti(c.aRegister.out, inM, instruction[3]), instruction[4], instruction[5], instruction[6], instruction[7], instruction[8], instruction[9])
	// Store in D register
	c.dRegister.tick(output, instruction[11])
	// Store in A register
	c.aRegister.tick(muxMulti(output, instruction, instruction[0]), instruction[10])
	// Check if we need to jump
	shouldJump1 := and(isNegative, instruction[13])
	shouldJump2 := and(isZero, instruction[14])
	shouldJump3 := and(and(not(isNegative), not(isZero)), instruction[12])
	shouldJump := or(or(shouldJump1, shouldJump2), shouldJump3)
	// Increment program counter, load A register, or reset
	c.count.tick(c.aRegister.out, true, shouldJump, reset)
}

type memory struct {
	ram ram16k
}

func (m *memory) tick(in []bool, address []bool, load bool) {
	m.ram.tick(in, address, load)
}

func (m *memory) read(address []bool) []bool {
	return m.ram.read(address)
}

func getMemory() memory {
	return memory{ram: getRAM16K(16)}
}

type rom32k struct {
	rom0 ram16k
	rom1 ram16k
}

func (r *rom32k) loadProgram(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	address := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	index := 0
	for scanner.Scan() {
		in := scanner.Text()
		fmt.Println(boolToStr(address))
		r.write(strToBool(strings.TrimSuffix(in, "\n")), address)
		dump(r, "z_dump-"+strconv.Itoa(index))
		address = increment(address)
		index++
	}
}

func (r *rom32k) write(in []bool, address []bool) {
	a0, a1 := dmux(true, address[14])
	r.rom0.tick(in, address[:14], a0)
	r.rom1.tick(in, address[:14], a1)
}

func (r *rom32k) read(address []bool) []bool {
	return muxMulti(r.rom0.read(address[:14]), r.rom1.read(address[:14]), address[14])
}

func getROM() rom32k {
	return rom32k{rom0: getRAM16K(16), rom1: getRAM16K(16)}
}
