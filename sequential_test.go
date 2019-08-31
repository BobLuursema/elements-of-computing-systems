package main

import "testing"

func TestDataFlipFlop(t *testing.T) {
	r := dataFlipFlop{currentValue: false}
	r.tick(true)
	if !(r.currentValue == true) {
		t.Error("Error 1")
	}
	r.tick(true)
	if !(r.currentValue == true) {
		t.Error("Error 2")
	}
	r.tick(false)
	if !(r.currentValue == false) {
		t.Error("Error 3")
	}
}

func TestBit(t *testing.T) {
	b := bit{dff: dataFlipFlop{}}
	if !(b.out == false) {
		t.Error("Error 1")
	}
	b.tick(true, false)
	if !(b.out == false) {
		t.Error("Error 2")
	}
	b.tick(true, true)
	if !(b.out == true) {
		t.Error("Error 3")
	}
}

func TestRegister(t *testing.T) {
	r := register{bits: []bit{bit{dff: dataFlipFlop{}}, bit{dff: dataFlipFlop{}}}}
	r.out = []bool{false, false}
	r.tick([]bool{false, true}, false)
	assertSlice(r.out, []bool{false, false}, t)
	r.tick([]bool{false, true}, true)
	assertSlice(r.out, []bool{false, true}, t)
}

func TestRAM8(t *testing.T) {
	r := getRAM8(3)
	result := r.read([]bool{false, false, false})
	assertSlice(result, []bool{false, false, false}, t)
	r.tick([]bool{true, false, true}, []bool{true, true, false}, true)
	result = r.read([]bool{false, false, false})
	assertSlice(result, []bool{false, false, false}, t)
	result = r.read([]bool{false, false, true})
	assertSlice(result, []bool{false, false, false}, t)
	result = r.read([]bool{false, true, false})
	assertSlice(result, []bool{false, false, false}, t)
	result = r.read([]bool{false, true, true})
	assertSlice(result, []bool{false, false, false}, t)
	result = r.read([]bool{true, false, false})
	assertSlice(result, []bool{false, false, false}, t)
	result = r.read([]bool{true, false, true})
	assertSlice(result, []bool{false, false, false}, t)
	result = r.read([]bool{true, true, false})
	assertSlice(result, []bool{true, false, true}, t)
	result = r.read([]bool{true, true, true})
	assertSlice(result, []bool{false, false, false}, t)
}

func TestRAM64(t *testing.T) {
	r := getRAM64(3)
	result := r.read([]bool{false, false, false, false, false, false})
	assertSlice(result, []bool{false, false, false}, t)
	r.tick([]bool{true, true, false}, []bool{false, true, true, false, false, true}, true)
	result = r.read([]bool{false, true, true, false, false, true})
	assertSlice(result, []bool{true, true, false}, t)
}

func TestRAM512(t *testing.T) {
	r := getRAM512(3)
	result := r.read([]bool{false, false, false, false, false, false, false, false, false})
	assertSlice(result, []bool{false, false, false}, t)
	r.tick([]bool{true, true, false}, []bool{false, false, true, true, false, true, false, false, true}, true)
	result = r.read([]bool{false, false, true, true, false, true, false, false, true})
	assertSlice(result, []bool{true, true, false}, t)
}

func TestRAM4K(t *testing.T) {
	r := getRAM4K(3)
	result := r.read([]bool{false, false, false, false, false, false, false, false, false, false, false, false})
	assertSlice(result, []bool{false, false, false}, t)
	r.tick([]bool{true, true, false}, []bool{false, false, false, false, false, false, false, false, false, false, false, false}, true)
	result = r.read([]bool{false, false, false, false, false, false, false, false, false, false, false, false})
	assertSlice(result, []bool{true, true, false}, t)
}

func TestRAM16K(t *testing.T) {
	r := getRAM16K(3)
	result := r.read(strToBool("00 0000 0100 0001"))
	assertSlice(result, strToBool("000"), t)
	r.tick(strToBool("001"), strToBool("00 0000 0100 0001"), true)
	result = r.read(strToBool("00 0000 0100 0001"))
	assertSlice(result, strToBool("001"), t)
}

func TestCounter(t *testing.T) {
	c := getCounter(4)
	// Check default
	assertSlice(c.count.out, strToBool("000"), t)
	// Check inc
	c.tick(strToBool("010"), true, false, false)
	assertSlice(c.count.out, strToBool("001"), t)
	// check load
	c.tick(strToBool("100"), false, true, false)
	assertSlice(c.count.out, strToBool("100"), t)
	// check reset
	c.tick(strToBool("010"), false, false, true)
	assertSlice(c.count.out, strToBool("000"), t)
	// check priority
	c.tick(strToBool("010"), true, true, true)
	assertSlice(c.count.out, strToBool("000"), t)
	c.tick(strToBool("010"), true, true, false)
	assertSlice(c.count.out, strToBool("010"), t)
	c.tick(strToBool("010"), true, false, false)
	assertSlice(c.count.out, strToBool("011"), t)
	c.tick(strToBool("010"), false, false, false)
	assertSlice(c.count.out, strToBool("011"), t)
}
