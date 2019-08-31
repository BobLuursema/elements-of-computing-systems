package main

/*
Chapter 1
Implementations of basic logic gates
*/

// Perform NAND
func nand(input1 bool, input2 bool) bool {
	return !(input1 && input2)
}

// Perform NOT
func not(input1 bool) bool {
	return nand(input1, input1)
}

// Perform AND
func and(input1 bool, input2 bool) bool {
	return not(nand(input1, input2))
}

// Perform OR
func or(input1 bool, input2 bool) bool {
	result1 := and(input1, input2)
	result2 := xor(input1, input2)
	return xor(result1, result2)
}

// Perform XOR
func xor(input1 bool, input2 bool) bool {
	result1 := nand(input1, input2)
	result2 := nand(result1, input1)
	result3 := nand(result1, input2)
	return nand(result2, result3)
}

// Perform NOR
func nor(input1 bool, input2 bool) bool {
	return not(or(input1, input2))
}

// Takes two inputs and returns one of the two based on the 1 bit selector
func mux(input1 bool, input2 bool, sel bool) bool {
	result1 := nand(input1, sel)
	result2 := and(input1, result1) // input1 alleen door als sel = 0

	result3 := and(input2, sel) // input2 alleen door als sel = 1

	return xor(result2, result3) // combineer result2 en result3
}

// Sends the input to output 1 or output 2 based on the 1 bit selector
func dmux(input1 bool, sel bool) (bool, bool) {
	result1 := not(sel)
	output1 := and(input1, result1) // stuur naar output2 als sel false is

	output2 := and(input1, sel) // stuur naar output1 als sel true is
	return output1, output2
}

// Perform bitwise NOT
func notMulti(input []bool) []bool {
	result := make([]bool, len(input))
	for index, value := range input {
		result[index] = not(value)
	}
	return result
}

// Perform bitwise AND
func andMulti(input1 []bool, input2 []bool) []bool {
	result := make([]bool, len(input1))
	for index := 0; index < len(input1); index++ {
		result[index] = and(input1[index], input2[index])
	}
	return result
}

// Perform bitwise OR
func orMulti(input1 []bool, input2 []bool) []bool {
	result := make([]bool, len(input1))
	for index := 0; index < len(input1); index++ {
		result[index] = or(input1[index], input2[index])
	}
	return result
}

// Takes two inputs and returns one of the two based on the 1 bit selector
func muxMulti(input1 []bool, input2 []bool, sel bool) []bool {
	result := make([]bool, len(input1))
	for index := 0; index < len(input1); index++ {
		result[index] = mux(input1[index], input2[index], sel)
	}
	return result
}

// Sends the input to output 1 or output 2 based on the 1 bit selector
func dmuxMulti(input []bool, sel bool) ([]bool, []bool) {
	result1 := make([]bool, len(input))
	result2 := make([]bool, len(input))
	for index, value := range input {
		r1, r2 := dmux(value, sel)
		result1[index] = r1
		result2[index] = r2
	}
	return result1, result2
}

// Return true if there is a 1 in any position, else 0
func orMultiWay(input []bool) bool {
	result := false
	for _, value := range input {
		result = or(result, value)
	}
	return result
}

// Takes four inputs and returns one of the four based on the 2 bit selector
func mux4WayMulti(input1 []bool, input2 []bool, input3 []bool, input4 []bool, sel []bool) []bool {
	result1 := muxMulti(input1, input2, sel[1])
	result2 := muxMulti(input3, input4, sel[1])
	return muxMulti(result1, result2, sel[0])
}

// Takes eight inputs and returns one of the four based on the 3 bit selector
func mux8WayMulti(input1 []bool, input2 []bool, input3 []bool, input4 []bool, input5 []bool, input6 []bool, input7 []bool, input8 []bool, sel []bool) []bool {
	result1 := mux4WayMulti(input1, input2, input3, input4, sel[1:])
	result2 := mux4WayMulti(input5, input6, input7, input8, sel[1:])
	return muxMulti(result1, result2, sel[0])
}

// Sends the input to one of the four outputs based on the 2 bit selector
func dmux4Way(input bool, sel []bool) (bool, bool, bool, bool) {
	result1, result2 := dmux(input, sel[0])
	result3, result4 := dmux(result1, sel[1])
	result5, result6 := dmux(result2, sel[1])
	return result3, result4, result5, result6
}

// Sends the input to one of the eight outputs based on the 3 bit selector
func dmux8Way(input bool, sel []bool) (bool, bool, bool, bool, bool, bool, bool, bool) {
	r1, r2, r3, r4 := dmux4Way(input, sel)
	o1, o2 := dmux(r1, sel[2])
	o3, o4 := dmux(r2, sel[2])
	o5, o6 := dmux(r3, sel[2])
	o7, o8 := dmux(r4, sel[2])
	return o1, o2, o3, o4, o5, o6, o7, o8
}
