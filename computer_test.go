package main

import "testing"

func TestStrToBool(t *testing.T) {
	output := strToBool("010101")
	assertSlice(output, []bool{false, true, false, true, false, true}, t)
}

func TestLoadProgram(t *testing.T) {
	rom := getROM()
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
