package main

import (
	"testing"
)

func TestNand(t *testing.T) {
	if nand(true, true) != false {
		t.Error("true, true != false")
	}
	if nand(true, false) != true {
		t.Error("true, false != true")
	}
	if nand(false, true) != true {
		t.Error("false, true != true")
	}
	if nand(false, false) != true {
		t.Error("false, false != true")
	}
}

func TestNot(t *testing.T) {
	if not(false) != true {
		t.Error("false != true")
	}
	if not(true) != false {
		t.Error("true != false")
	}
}

func TestAnd(t *testing.T) {
	if and(true, true) != true {
		t.Error("true, true != true")
	}
	if and(true, false) != false {
		t.Error("true, false != false")
	}
	if and(false, true) != false {
		t.Error("false, true != false")
	}
	if and(false, false) != false {
		t.Error("false, false != false")
	}
}

func TestOr(t *testing.T) {
	if or(true, true) != true {
		t.Error("true, true != true")
	}
	if or(true, false) != true {
		t.Error("true, false != true")
	}
	if or(false, true) != true {
		t.Error("false, true != true")
	}
	if or(false, false) != false {
		t.Error("false, false != false")
	}
}

func TestXor(t *testing.T) {
	if xor(true, true) != false {
		t.Error("true, true != false")
	}
	if xor(true, false) != true {
		t.Error("true, false != true")
	}
	if xor(false, true) != true {
		t.Error("false, true != true")
	}
	if xor(false, false) != false {
		t.Error("false, false != false")
	}
}

func TestNor(t *testing.T) {
	if nor(true, true) != false {
		t.Error("true, true != false")
	}
	if nor(true, false) != false {
		t.Error("true, false != false")
	}
	if nor(false, true) != false {
		t.Error("false, true != false")
	}
	if nor(false, false) != true {
		t.Error("false, false != true")
	}
}

func TestMux(t *testing.T) {
	if !(mux(false, false, false) == false) {
		t.Error("false, false, false != false")
	}
	if !(mux(false, true, false) == false) {
		t.Error("false, true, false != false")
	}
	if !(mux(true, false, false) == true) {
		t.Error("true, false, false != true")
	}
	if !(mux(true, true, false) == true) {
		t.Error("true, true, false != true")
	}
	if !(mux(false, false, true) == false) {
		t.Error("false, false, true != false")
	}
	if !(mux(false, true, true) == true) {
		t.Error("false, true, true != true")
	}
	if !(mux(true, false, true) == false) {
		t.Error("true, false, true != false")
	}
	if !(mux(true, true, true) == true) {
		t.Error("true, true, true != true")
	}
}

func TestDmux(t *testing.T) {
	o1, o2 := dmux(true, true)
	if o1 != false || o2 != true {
		t.Error("true, true != false, true")
	}
	o1, o2 = dmux(true, false)
	if o1 != true || o2 != false {
		t.Error("true, false != true, false")
	}
	o1, o2 = dmux(false, true)
	if o1 != false || o2 != false {
		t.Error("false, true != false, false")
	}
	o1, o2 = dmux(false, false)
	if o1 != false || o2 != false {
		t.Error("false, false != false, false")
	}
}

func TestNotMulti(t *testing.T) {
	result := notMulti(strToBool("1011"))
	assertSlice(result, strToBool("0100"), t)
}

func TestAndMulti(t *testing.T) {
	result := andMulti(strToBool("1100"), strToBool("1010"))
	assertSlice(result, strToBool("1000"), t)
}

func TestOrMulti(t *testing.T) {
	result := orMulti(strToBool("1100"), strToBool("1010"))
	assertSlice(result, strToBool("1110"), t)
}

func TestMuxMulti(t *testing.T) {
	result := muxMulti(strToBool("1100"), strToBool("1010"), false)
	assertSlice(result, strToBool("1100"), t)
	result = muxMulti(strToBool("1100"), strToBool("1010"), true)
	assertSlice(result, strToBool("1010"), t)
}

func TestDmuxMulti(t *testing.T) {
	result1, result2 := dmuxMulti(strToBool("11"), false)
	assertSlice(result1, strToBool("11"), t)
	assertSlice(result2, strToBool("00"), t)
	result1, result2 = dmuxMulti(strToBool("11"), true)
	assertSlice(result1, strToBool("00"), t)
	assertSlice(result2, strToBool("11"), t)
}

func TestOrMultiWay(t *testing.T) {
	if orMultiWay(strToBool("000")) != false {
		t.Error("All false != false")
	}
	if orMultiWay(strToBool("010")) != true {
		t.Error("One true != true")
	}
}

func TestMux4WayMulti(t *testing.T) {
	result := mux4WayMulti(strToBool("11"), strToBool("10"), strToBool("01"), strToBool("00"), strToBool("00"))
	assertSlice(result, strToBool("11"), t)
	result = mux4WayMulti(strToBool("11"), strToBool("10"), strToBool("01"), strToBool("00"), strToBool("01"))
	assertSlice(result, strToBool("10"), t)
	result = mux4WayMulti(strToBool("11"), strToBool("10"), strToBool("01"), strToBool("00"), strToBool("10"))
	assertSlice(result, strToBool("01"), t)
	result = mux4WayMulti(strToBool("11"), strToBool("10"), strToBool("01"), strToBool("00"), strToBool("11"))
	assertSlice(result, strToBool("00"), t)
}

func TestMux8WayMulti(t *testing.T) {
	result := mux8WayMulti(strToBool("111"), strToBool("101"), strToBool("011"), strToBool("001"), strToBool("110"), strToBool("100"), strToBool("010"), strToBool("000"), strToBool("000"))
	assertSlice(result, strToBool("111"), t)
	result = mux8WayMulti(strToBool("111"), strToBool("101"), strToBool("011"), strToBool("001"), strToBool("110"), strToBool("100"), strToBool("010"), strToBool("000"), strToBool("101"))
	assertSlice(result, strToBool("100"), t)
}

func TestDmux4Way(t *testing.T) {
	result1, result2, result3, result4 := dmux4Way(true, strToBool("00"))
	if result1 != true || result2 != false || result3 != false || result4 != false {
		t.Error("Select 1 failed")
	}
	result1, result2, result3, result4 = dmux4Way(true, strToBool("01"))
	if result1 != false || result2 != true || result3 != false || result4 != false {
		t.Error("Select 2 failed")
	}
	result1, result2, result3, result4 = dmux4Way(true, strToBool("10"))
	if result1 != false || result2 != false || result3 != true || result4 != false {
		t.Error("Select 3 failed")
	}
	result1, result2, result3, result4 = dmux4Way(true, strToBool("11"))
	if result1 != false || result2 != false || result3 != false || result4 != true {
		t.Error("Select 4 failed")
	}
}

func TestDmux8Way(t *testing.T) {
	o1, _, _, _, _, _, _, _ := dmux8Way(true, strToBool("000"))
	if o1 != true {
		t.Error("Select 1 failed")
	}
	_, o2, _, _, _, _, _, _ := dmux8Way(true, strToBool("001"))
	if o2 != true {
		t.Error("Select 2 failed")
	}
	_, _, _, _, _, o6, _, _ := dmux8Way(true, strToBool("101"))
	if o6 != true {
		t.Error("Select 5 failed")
	}
}
