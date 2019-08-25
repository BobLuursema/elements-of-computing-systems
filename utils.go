package main

import (
	"os"
	"testing"
)

func strToBool(in string) []bool {
	output := make([]bool, len(in))
	for i, c := range in {
		output[i] = c == '1'
	}
	return output
}

func boolToStr(in []bool) string {
	output := ""
	for _, v := range in {
		binary := ""
		if v {
			binary = "1"
		} else {
			binary = "0"
		}
		output = output + binary
	}
	return output
}

func assertSlice(output []bool, input []bool, t *testing.T) {
	if len(input) != len(output) {
		t.Error("Length is different.")
	}
	for index, value := range input {
		if value != output[index] {
			t.Errorf("Index %d is different", index)
		}
	}
}

func dump(rom *rom32k, filename string) {
	f, _ := os.Create(filename)
	defer f.Close()
	for _, r4k := range rom.rom0.ram4ks {
		for _, r512 := range r4k.ram512s {
			for _, r64 := range r512.ram64s {
				for _, r8 := range r64.ram8s {
					for _, r := range r8.registers {
						f.WriteString(boolToStr(r.out) + "\n")
					}
				}
			}
		}
	}
}
