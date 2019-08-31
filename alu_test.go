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
	result := add(strToBool("000"), strToBool("010"))
	assertSlice(result, strToBool("010"), t)
	result = add(strToBool("001"), strToBool("010"))
	assertSlice(result, strToBool("011"), t)
	result = add(strToBool("001"), strToBool("011"))
	assertSlice(result, strToBool("100"), t)
	result = add(strToBool("101"), strToBool("011"))
	assertSlice(result, strToBool("000"), t)
}

func TestIncrement(t *testing.T) {
	result := increment(strToBool("00"))
	assertSlice(result, strToBool("01"), t)
	result = increment(strToBool("01"))
	assertSlice(result, strToBool("10"), t)
	result = increment(strToBool("11"))
	assertSlice(result, strToBool("00"), t)
}

func TestAlu(t *testing.T) {
	output, zero, negative := alu(strToBool("00"), strToBool("01"), false, false, false, false, true, false)
	assertSlice(output, strToBool("01"), t)
	if !(zero == false && negative == false) {
		t.Error("Error 1")
	}
	output, zero, negative = alu(strToBool("00"), strToBool("01"), false, false, false, false, false, false)
	assertSlice(output, strToBool("00"), t)
	if !(zero == true && negative == false) {
		t.Error("Error 2")
	}
	output, zero, negative = alu(strToBool("01"), strToBool("01"), false, false, false, false, true, false)
	assertSlice(output, strToBool("10"), t)
	if !(zero == false && negative == true) {
		t.Error("Error 3")
	}
	output, zero, negative = alu(strToBool("10"), strToBool("01"), true, false, false, false, true, false)
	assertSlice(output, strToBool("01"), t)
	if !(zero == false && negative == false) {
		t.Error("Error 4")
	}
	output, zero, negative = alu(strToBool("10"), strToBool("01"), false, false, true, false, true, false)
	assertSlice(output, strToBool("10"), t)
	if !(zero == false && negative == true) {
		t.Error("Error 5")
	}
	output, zero, negative = alu(strToBool("10"), strToBool("01"), false, true, false, false, false, false)
	assertSlice(output, strToBool("01"), t)
	if !(zero == false && negative == false) {
		t.Error("Error 6")
	}
	output, zero, negative = alu(strToBool("10"), strToBool("01"), false, false, false, true, false, false)
	assertSlice(output, strToBool("10"), t)
	if !(zero == false && negative == true) {
		t.Error("Error 7")
	}
	output, zero, negative = alu(strToBool("10"), strToBool("01"), false, false, false, true, false, true)
	assertSlice(output, strToBool("01"), t)
	if !(zero == false && negative == false) {
		t.Error("Error 8")
	}
}
