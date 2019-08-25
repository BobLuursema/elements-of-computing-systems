package main

/*
Chapter 2
Constructing an ALU.
*/

func halfAdder(input1 bool, input2 bool) (bool, bool) {
	sum := xor(input1, input2)
	carry := and(input1, input2)
	return sum, carry
}

func fullAdder(input1 bool, input2 bool, input3 bool) (bool, bool) {
	sum1, carry1 := halfAdder(input1, input2)
	sum2, carry2 := halfAdder(sum1, input3)
	return sum2, or(carry1, carry2)
}

func add(input1 []bool, input2 []bool) []bool {
	result := make([]bool, len(input1))
	carry := false
	sum := false
	for index := len(input1) - 1; index >= 0; index-- {
		sum, carry = fullAdder(input1[index], input2[index], carry)
		result[index] = sum
	}
	return result
}

func increment(input []bool) []bool {
	one := make([]bool, len(input))
	one[len(one)-1] = true
	return add(input, one)
}

func alu(inputX []bool, inputY []bool, zeroX bool, negateX bool, zeroY bool, negateY bool, addFunction bool, negateOutput bool) ([]bool, bool, bool) {
	zero := make([]bool, len(inputX))
	inputX = muxMulti(inputX, zero, zeroX)
	inputX = muxMulti(inputX, notMulti(inputX), negateX)
	inputY = muxMulti(inputY, zero, zeroY)
	inputY = muxMulti(inputY, notMulti(inputY), negateY)
	output := muxMulti(andMulti(inputX, inputY), add(inputX, inputY), addFunction)
	output = muxMulti(output, notMulti(output), negateOutput)
	return output, not(orMultiWay(output)), output[0]
}
