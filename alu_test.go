package main

import "testing"

func TestHalfAdder(t *testing.T) {
	sum, carry := halfAdder(false, false)
	if sum != false || carry != false {
		t.Error("Error 1")
	}
	sum, carry = halfAdder(false, true)
	if sum != true || carry != false {
		t.Error("Error 1")
	}
	sum, carry = halfAdder(true, false)
	if sum != true || carry != false {
		t.Error("Error 1")
	}
	sum, carry = halfAdder(true, true)
	if sum != false || carry != true {
		t.Error("Error 1")
	}
}

func TestFullAdder(t *testing.T) {
	sum, carry := fullAdder(false, false, false)
	if !(sum == false && carry == false) {
		t.Error("Error 1")
	}
	sum, carry = fullAdder(true, false, false)
	if !(sum == true && carry == false) {
		t.Error("Error 2")
	}
	sum, carry = fullAdder(false, true, false)
	if !(sum == true && carry == false) {
		t.Error("Error 3")
	}
	sum, carry = fullAdder(true, true, false)
	if !(sum == false && carry == true) {
		t.Error("Error 4")
	}
	sum, carry = fullAdder(false, false, true)
	if !(sum == true && carry == false) {
		t.Error("Error 5")
	}
	sum, carry = fullAdder(true, false, true)
	if !(sum == false && carry == true) {
		t.Error("Error 6")
	}
	sum, carry = fullAdder(false, true, true)
	if !(sum == false && carry == true) {
		t.Error("Error 7")
	}
	sum, carry = fullAdder(true, true, true)
	if !(sum == true && carry == true) {
		t.Error("Error 8")
	}
}

func TestAdd(t *testing.T) {
	result := add([]bool{false, false, false}, []bool{false, true, false})
	assertSlice(result, []bool{false, true, false}, t)
	result = add([]bool{false, false, true}, []bool{false, true, false})
	assertSlice(result, []bool{false, true, true}, t)
	result = add([]bool{false, false, true}, []bool{false, true, true})
	assertSlice(result, []bool{true, false, false}, t)
	result = add([]bool{true, false, true}, []bool{false, true, true})
	assertSlice(result, []bool{false, false, false}, t)
}

func TestIncrement(t *testing.T) {
	result := increment([]bool{false, false})
	assertSlice(result, []bool{false, true}, t)
	result = increment([]bool{false, true})
	assertSlice(result, []bool{true, false}, t)
	result = increment([]bool{true, true})
	assertSlice(result, []bool{false, false}, t)
}

func TestAlu(t *testing.T) {
	output, zero, negative := alu([]bool{false, false}, []bool{false, true}, false, false, false, false, true, false)
	assertSlice(output, []bool{false, true}, t)
	if !(zero == false && negative == false) {
		t.Error("Error 1")
	}
	output, zero, negative = alu([]bool{false, false}, []bool{false, true}, false, false, false, false, false, false)
	assertSlice(output, []bool{false, false}, t)
	if !(zero == true && negative == false) {
		t.Error("Error 2")
	}
	output, zero, negative = alu([]bool{false, true}, []bool{false, true}, false, false, false, false, true, false)
	assertSlice(output, []bool{true, false}, t)
	if !(zero == false && negative == true) {
		t.Error("Error 3")
	}
	output, zero, negative = alu([]bool{true, false}, []bool{false, true}, true, false, false, false, true, false)
	assertSlice(output, []bool{false, true}, t)
	if !(zero == false && negative == false) {
		t.Error("Error 4")
	}
	output, zero, negative = alu([]bool{true, false}, []bool{false, true}, false, false, true, false, true, false)
	assertSlice(output, []bool{true, false}, t)
	if !(zero == false && negative == true) {
		t.Error("Error 5")
	}
	output, zero, negative = alu([]bool{true, false}, []bool{false, true}, false, true, false, false, false, false)
	assertSlice(output, []bool{false, true}, t)
	if !(zero == false && negative == false) {
		t.Error("Error 6")
	}
	output, zero, negative = alu([]bool{true, false}, []bool{false, true}, false, false, false, true, false, false)
	assertSlice(output, []bool{true, false}, t)
	if !(zero == false && negative == true) {
		t.Error("Error 7")
	}
	output, zero, negative = alu([]bool{true, false}, []bool{false, true}, false, false, false, true, false, true)
	assertSlice(output, []bool{false, true}, t)
	if !(zero == false && negative == false) {
		t.Error("Error 8")
	}
}
