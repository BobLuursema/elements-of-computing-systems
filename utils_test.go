package main

import (
	"testing"
)

func TestIntToBools(t *testing.T) {
	output := intToBools(16, 9)
	assertSlice(output, strToBool("000010000"), t)
}
