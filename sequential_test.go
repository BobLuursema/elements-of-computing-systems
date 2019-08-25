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
	result := r.read([]bool{false, false, false, false, false, false, false, true, false, false, false, false, false, true})
	assertSlice(result, []bool{false, false, false}, t)
	r.tick([]bool{false, false, true}, []bool{false, false, false, false, false, false, false, true, false, false, false, false, false, true}, true)
	result = r.read([]bool{false, false, false, false, false, false, false, true, false, false, false, false, false, true})
	assertSlice(result, []bool{false, false, true}, t)
}

func TestCounter(t *testing.T) {
	c := getCounter(3)
	// Check default
	assertSlice(c.count.out, []bool{false, false, false}, t)
	// Check inc
	c.tick([]bool{false, true, false}, true, false, false)
	assertSlice(c.count.out, []bool{false, false, true}, t)
	// check load
	c.tick([]bool{true, false, false}, false, true, false)
	assertSlice(c.count.out, []bool{true, false, false}, t)
	// check reset
	c.tick([]bool{false, true, false}, false, false, true)
	assertSlice(c.count.out, []bool{false, false, false}, t)
	// check priority
	c.tick([]bool{false, true, false}, true, true, true)
	assertSlice(c.count.out, []bool{false, false, false}, t)
	c.tick([]bool{false, true, false}, true, true, false)
	assertSlice(c.count.out, []bool{false, true, false}, t)
	c.tick([]bool{false, true, false}, true, false, false)
	assertSlice(c.count.out, []bool{false, true, true}, t)
	c.tick([]bool{false, true, false}, false, false, false)
	assertSlice(c.count.out, []bool{false, true, true}, t)
}
