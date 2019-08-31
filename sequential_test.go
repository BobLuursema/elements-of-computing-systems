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
	r.out = strToBool("00")
	r.tick(strToBool("01"), false)
	assertSlice(r.out, strToBool("00"), t)
	r.tick(strToBool("01"), true)
	assertSlice(r.out, strToBool("01"), t)
}

func TestRAM8(t *testing.T) {
	r := getRAM8(3)
	result := r.read(strToBool("000"))
	assertSlice(result, strToBool("000"), t)
	r.tick(strToBool("101"), strToBool("110"), true)
	result = r.read(strToBool("000"))
	assertSlice(result, strToBool("000"), t)
	result = r.read(strToBool("001"))
	assertSlice(result, strToBool("000"), t)
	result = r.read(strToBool("010"))
	assertSlice(result, strToBool("000"), t)
	result = r.read(strToBool("011"))
	assertSlice(result, strToBool("000"), t)
	result = r.read(strToBool("100"))
	assertSlice(result, strToBool("000"), t)
	result = r.read(strToBool("101"))
	assertSlice(result, strToBool("000"), t)
	result = r.read(strToBool("110"))
	assertSlice(result, strToBool("101"), t)
	result = r.read(strToBool("111"))
	assertSlice(result, strToBool("000"), t)
}

func TestRAM64(t *testing.T) {
	r := getRAM64(3)
	result := r.read(strToBool("000000"))
	assertSlice(result, strToBool("000"), t)
	r.tick(strToBool("110"), strToBool("011001"), true)
	result = r.read(strToBool("011001"))
	assertSlice(result, strToBool("110"), t)
}

func TestRAM512(t *testing.T) {
	r := getRAM512(3)
	result := r.read(strToBool("000000000"))
	assertSlice(result, strToBool("000"), t)
	r.tick(strToBool("110"), strToBool("001101001"), true)
	result = r.read(strToBool("001101001"))
	assertSlice(result, strToBool("110"), t)
}

func TestRAM4K(t *testing.T) {
	r := getRAM4K(3)
	result := r.read(strToBool("000000000000"))
	assertSlice(result, strToBool("000"), t)
	r.tick(strToBool("110"), strToBool("000000000000"), true)
	result = r.read(strToBool("000000000000"))
	assertSlice(result, strToBool("110"), t)
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
