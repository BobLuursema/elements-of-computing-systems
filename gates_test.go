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
	result := notMulti([]bool{true, false, true, true})
	assertSlice(result, []bool{false, true, false, false}, t)
}

func TestAndMulti(t *testing.T) {
	result := andMulti([]bool{true, true, false, false}, []bool{true, false, true, false})
	assertSlice(result, []bool{true, false, false, false}, t)
}

func TestOrMulti(t *testing.T) {
	result := orMulti([]bool{true, true, false, false}, []bool{true, false, true, false})
	assertSlice(result, []bool{true, true, true, false}, t)
}

func TestMuxMulti(t *testing.T) {
	result := muxMulti([]bool{true, true, false, false}, []bool{true, false, true, false}, false)
	assertSlice(result, []bool{true, true, false, false}, t)
	result = muxMulti([]bool{true, true, false, false}, []bool{true, false, true, false}, true)
	assertSlice(result, []bool{true, false, true, false}, t)
}

func TestDmuxMulti(t *testing.T) {
	result1, result2 := dmuxMulti([]bool{true, true}, false)
	assertSlice(result1, []bool{true, true}, t)
	assertSlice(result2, []bool{false, false}, t)
	result1, result2 = dmuxMulti([]bool{true, true}, true)
	assertSlice(result1, []bool{false, false}, t)
	assertSlice(result2, []bool{true, true}, t)
}

func TestOrMultiWay(t *testing.T) {
	if orMultiWay([]bool{false, false, false}) != false {
		t.Error("All false != false")
	}
	if orMultiWay([]bool{false, true, false}) != true {
		t.Error("One true != true")
	}
}

func TestMux4WayMulti(t *testing.T) {
	result := mux4WayMulti([]bool{true, true}, []bool{true, false}, []bool{false, true}, []bool{false, false}, []bool{false, false})
	assertSlice(result, []bool{true, true}, t)
	result = mux4WayMulti([]bool{true, true}, []bool{true, false}, []bool{false, true}, []bool{false, false}, []bool{false, true})
	assertSlice(result, []bool{true, false}, t)
	result = mux4WayMulti([]bool{true, true}, []bool{true, false}, []bool{false, true}, []bool{false, false}, []bool{true, false})
	assertSlice(result, []bool{false, true}, t)
	result = mux4WayMulti([]bool{true, true}, []bool{true, false}, []bool{false, true}, []bool{false, false}, []bool{true, true})
	assertSlice(result, []bool{false, false}, t)
}

func TestMux8WayMulti(t *testing.T) {
	result := mux8WayMulti([]bool{true, true, true}, []bool{true, false, true}, []bool{false, true, true}, []bool{false, false, true}, []bool{true, true, false}, []bool{true, false, false}, []bool{false, true, false}, []bool{false, false, false}, []bool{false, false, false})
	assertSlice(result, []bool{true, true, true}, t)
	result = mux8WayMulti([]bool{true, true, true}, []bool{true, false, true}, []bool{false, true, true}, []bool{false, false, true}, []bool{true, true, false}, []bool{true, false, false}, []bool{false, true, false}, []bool{false, false, false}, []bool{true, false, true})
	assertSlice(result, []bool{true, false, false}, t)
}

func TestDmux4Way(t *testing.T) {
	result1, result2, result3, result4 := dmux4Way(true, []bool{false, false})
	if result1 != true || result2 != false || result3 != false || result4 != false {
		t.Error("Select 1 failed")
	}
	result1, result2, result3, result4 = dmux4Way(true, []bool{false, true})
	if result1 != false || result2 != true || result3 != false || result4 != false {
		t.Error("Select 2 failed")
	}
	result1, result2, result3, result4 = dmux4Way(true, []bool{true, false})
	if result1 != false || result2 != false || result3 != true || result4 != false {
		t.Error("Select 3 failed")
	}
	result1, result2, result3, result4 = dmux4Way(true, []bool{true, true})
	if result1 != false || result2 != false || result3 != false || result4 != true {
		t.Error("Select 4 failed")
	}
}

func TestDmux8Way(t *testing.T) {
	o1, _, _, _, _, _, _, _ := dmux8Way(true, []bool{false, false, false})
	if o1 != true {
		t.Error("Select 1 failed")
	}
	_, o2, _, _, _, _, _, _ := dmux8Way(true, []bool{false, false, true})
	if o2 != true {
		t.Error("Select 2 failed")
	}
	_, _, _, _, _, o6, _, _ := dmux8Way(true, []bool{true, false, true})
	if o6 != true {
		t.Error("Select 5 failed")
	}
}
