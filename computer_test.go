package main

import (
	"fmt"
	"testing"
)

func TestStrToBool(t *testing.T) {
	output := strToBool("010101")
	assertSlice(output, []bool{false, true, false, true, false, true}, t)
}

func TestLoadProgram(t *testing.T) {
	rom := getROM(16)
	rom.loadProgram("test.hack")
	assertSlice(
		rom.read([]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}),
		strToBool("0000000000010000"),
		t)
	assertSlice(
		rom.read([]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, true}),
		strToBool("1110111111001000"),
		t)
	assertSlice(
		rom.read([]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, true, false}),
		strToBool("0000000000000010"),
		t)
}

func TestComputer(t *testing.T) {
	comp := getComputer(16)
	comp.loadProgram("test.hack")
	comp.tick(false)
	comp.tick(false)
	comp.tick(false)
	comp.tick(false)
	comp.tick(false)
	comp.tick(false)
	comp.tick(false)
	comp.tick(false)
	comp.tick(false)
}

func TestCPU(t *testing.T) {
	parameters := []struct {
		aRegister       []bool
		dRegister       []bool
		instruction     []bool
		inM             []bool
		outM            []bool
		writeM          bool
		addressM        []bool
		pc              []bool
		aRegisterOutput []bool
		dRegisterOutput []bool
	}{
		// GENERATE OUTPUT TESTS
		{ // 0 | 0
			strToBool("1000000000000000"), // aRegister
			strToBool("0100000000000000"), // dRegister
			strToBool("1110101010000000"), // instruction
			strToBool("0010000000000000"), // inM
			strToBool("0000000000000000"), // outM
			false,                         // writeM
			strToBool("1000000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("1000000000000000"), // aRegisterOutput
			strToBool("0100000000000000"), // dRegisterOutput
		},
		{ // 1 | 1
			strToBool("1000000000000000"), // aRegister
			strToBool("0100000000000000"), // dRegister
			strToBool("1110111111000000"), // instruction
			strToBool("0010000000000000"), // inM
			strToBool("0000000000000001"), // outM
			false,                         // writeM
			strToBool("1000000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("1000000000000000"), // aRegisterOutput
			strToBool("0100000000000000"), // dRegisterOutput
		},
		{ // -1 | 2
			strToBool("1000000000000000"), // aRegister
			strToBool("0100000000000000"), // dRegister
			strToBool("1110111010000000"), // instruction
			strToBool("0010000000000000"), // inM
			strToBool("1111111111111111"), // outM
			false,                         // writeM
			strToBool("1000000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("1000000000000000"), // aRegisterOutput
			strToBool("0100000000000000"), // dRegisterOutput
		},
		{ // D | 3
			strToBool("1000000000000000"), // aRegister
			strToBool("0100000000000000"), // dRegister
			strToBool("1110001100000000"), // instruction
			strToBool("0010000000000000"), // inM
			strToBool("0100000000000000"), // outM
			false,                         // writeM
			strToBool("1000000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("1000000000000000"), // aRegisterOutput
			strToBool("0100000000000000"), // dRegisterOutput
		},
		{ // A | 4
			strToBool("1000000000000000"), // aRegister
			strToBool("0100000000000000"), // dRegister
			strToBool("1110110000000000"), // instruction
			strToBool("0010000000000000"), // inM
			strToBool("1000000000000000"), // outM
			false,                         // writeM
			strToBool("1000000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("1000000000000000"), // aRegisterOutput
			strToBool("0100000000000000"), // dRegisterOutput
		},
		{ // M | 5
			strToBool("1000000000000000"), // aRegister
			strToBool("0100000000000000"), // dRegister
			strToBool("1111110000000000"), // instruction
			strToBool("0010000000000000"), // inM
			strToBool("0010000000000000"), // outM
			false,                         // writeM
			strToBool("1000000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("1000000000000000"), // aRegisterOutput
			strToBool("0100000000000000"), // dRegisterOutput
		},
		{ // !D | 6
			strToBool("1000000000000000"), // aRegister
			strToBool("0100000000000000"), // dRegister
			strToBool("1110001101000000"), // instruction
			strToBool("0010000000000000"), // inM
			strToBool("1011111111111111"), // outM
			false,                         // writeM
			strToBool("1000000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("1000000000000000"), // aRegisterOutput
			strToBool("0100000000000000"), // dRegisterOutput
		},
		{ // !A | 7
			strToBool("1000000000000000"), // aRegister
			strToBool("0100000000000000"), // dRegister
			strToBool("1110110001000000"), // instruction
			strToBool("0010000000000000"), // inM
			strToBool("0111111111111111"), // outM
			false,                         // writeM
			strToBool("1000000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("1000000000000000"), // aRegisterOutput
			strToBool("0100000000000000"), // dRegisterOutput
		},
		{ // !M | 8
			strToBool("1000000000000000"), // aRegister
			strToBool("0100000000000000"), // dRegister
			strToBool("1111110001000000"), // instruction
			strToBool("0010000000000000"), // inM
			strToBool("1101111111111111"), // outM
			false,                         // writeM
			strToBool("1000000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("1000000000000000"), // aRegisterOutput
			strToBool("0100000000000000"), // dRegisterOutput
		},
		{ // -D | 9
			strToBool("1000000000000000"), // aRegister
			strToBool("0100000000000000"), // dRegister
			strToBool("1110001111000000"), // instruction
			strToBool("0010000000000000"), // inM
			strToBool("1100000000000000"), // outM
			false,                         // writeM
			strToBool("1000000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("1000000000000000"), // aRegisterOutput
			strToBool("0100000000000000"), // dRegisterOutput
		},
		{ // -A | 10
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1110110011000000"), // instruction
			strToBool("0010000000000000"), // inM
			strToBool("1100000000000000"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // -M | 11
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1111110011000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("1111111111111100"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // D+1 | 12
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1110011111000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("0010000000000001"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // A+1 | 13
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1110110111000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("0100000000000001"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // M+1 | 14
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1111110111000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("0000000000000101"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // D-1 | 15
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1110001110000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("0001111111111111"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // A-1 | 16
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1110110010000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("0011111111111111"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // M-1 | 17
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1111110010000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("0000000000000011"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // D+A | 18
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1110000010000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("0110000000000000"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // D+M | 19
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1111000010000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("0010000000000100"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // D-A | 20
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1110010011000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("1110000000000000"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // D-M | 21
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1111010011000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("0001111111111100"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // A-D | 22
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1110000111000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("0010000000000000"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // M-D | 23
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000000000"), // dRegister
			strToBool("1111000111000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("1110000000000100"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000000000"), // dRegisterOutput
		},
		{ // D&A | 24
			strToBool("0100000000000001"), // aRegister
			strToBool("0010000000000001"), // dRegister
			strToBool("1110000000000000"), // instruction
			strToBool("0000000000000100"), // inM
			strToBool("0000000000000001"), // outM
			false,                         // writeM
			strToBool("0100000000000001"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000001"), // aRegisterOutput
			strToBool("0010000000000001"), // dRegisterOutput
		},
		{ // D&M | 25
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000100000"), // dRegister
			strToBool("1111000000000000"), // instruction
			strToBool("0000000000100100"), // inM
			strToBool("0000000000100000"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000100000"), // dRegisterOutput
		},
		{ // D|A | 26
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000100000"), // dRegister
			strToBool("1110010101000000"), // instruction
			strToBool("0000000000100100"), // inM
			strToBool("0110000000100000"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000100000"), // dRegisterOutput
		},
		{ // D|M | 27
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000100000"), // dRegister
			strToBool("1111010101000000"), // instruction
			strToBool("0000000000100100"), // inM
			strToBool("0010000000100100"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000100000"), // dRegisterOutput
		},
		// STORE OUTPUT TESTS
		{ // A | 28
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000100000"), // dRegister
			strToBool("1110111111100000"), // instruction
			strToBool("0000000000100100"), // inM
			strToBool("0000000000000001"), // outM
			false,                         // writeM
			strToBool("0000000000000001"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0000000000000001"), // aRegisterOutput
			strToBool("0010000000100000"), // dRegisterOutput
		},
		{ // D | 29
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000100000"), // dRegister
			strToBool("1110111111010000"), // instruction
			strToBool("0000000000100100"), // inM
			strToBool("0000000000000001"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0000000000000001"), // dRegisterOutput
		},
		{ // M[A] | 30
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000100000"), // dRegister
			strToBool("1110111111001000"), // instruction
			strToBool("0000000000100100"), // inM
			strToBool("0000000000000001"), // outM
			true,                          // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000100000"), // dRegisterOutput
		},
		// A INSTRUCTION
		{ // A instruction | 31
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000100000"), // dRegister
			strToBool("0000000000000010"), // instruction
			strToBool("0000000000100100"), // inM
			strToBool("0000000000000000"), // outM
			false,                         // writeM
			strToBool("0000000000000010"), // addressM
			strToBool("0000000000000001"), // pc
			strToBool("0000000000000010"), // aRegisterOutput
			strToBool("0010000000100000"), // dRegisterOutput
		},
		// SHOULD JUMP
		{ // jump if <0 | 32
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000100000"), // dRegister
			strToBool("1110111010000100"), // instruction
			strToBool("0000000000100100"), // inM
			strToBool("1111111111111111"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0100000000000000"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000100000"), // dRegisterOutput
		},
		{ // jump if =0 | 33
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000100000"), // dRegister
			strToBool("1110101010000010"), // instruction
			strToBool("0000000000100100"), // inM
			strToBool("0000000000000000"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0100000000000000"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000100000"), // dRegisterOutput
		},
		{ /// jump if >0 | 34
			strToBool("0100000000000000"), // aRegister
			strToBool("0010000000100000"), // dRegister
			strToBool("1110111111000001"), // instruction
			strToBool("0000000000100100"), // inM
			strToBool("0000000000000001"), // outM
			false,                         // writeM
			strToBool("0100000000000000"), // addressM
			strToBool("0100000000000000"), // pc
			strToBool("0100000000000000"), // aRegisterOutput
			strToBool("0010000000100000"), // dRegisterOutput
		},
	}
	for i, p := range parameters {
		cpu := cpu{
			dRegister: getRegister(16),
			aRegister: getRegister(16),
			count:     getCounter(16)}
		cpu.aRegister.tick(p.aRegister, true)
		cpu.dRegister.tick(p.dRegister, true)
		outM, writeM, addressM, pc := cpu.tick(
			p.instruction,
			p.inM,
			false,
		)
		if err, message := assertSlice2(outM, p.outM); err {
			t.Errorf("Error in test %d for outM: %s", i, message)
		}
		if writeM != p.writeM {
			t.Error("writeM is incorrect")
		}
		if err, message := assertSlice2(addressM, p.addressM); err {
			t.Errorf("Error in test %d for addressM: %s", i, message)
		}
		if err, message := assertSlice2(pc, p.pc); err {
			t.Errorf("Error in test %d for pc: %s", i, message)
		}
		if err, message := assertSlice2(cpu.dRegister.out, p.dRegisterOutput); err {
			t.Errorf("Error in test %d for dRegister: %s", i, message)
		}
		if err, message := assertSlice2(cpu.aRegister.out, p.aRegisterOutput); err {
			t.Errorf("Error in test %d for aRegister: %s", i, message)
		}
	}
}

func assertSlice2(actual []bool, expected []bool) (bool, string) {
	if len(expected) != len(actual) {
		return true, "Length is different."
	}
	for index, value := range expected {
		if value != actual[index] {
			return true, fmt.Sprintf("Index %d is different", index)
		}
	}
	return false, ""
}
