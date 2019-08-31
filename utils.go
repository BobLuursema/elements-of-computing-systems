package main

import (
	"math"
	"os"
	"strings"
	"testing"
)

func strToBool(in string) []bool {
	in = strings.ReplaceAll(in, " ", "")
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

func assertSlice(actual []bool, expected []bool, t *testing.T) {
	if len(expected) != len(actual) {
		t.Error("Length is different.")
	}
	for index, value := range expected {
		if value != actual[index] {
			t.Errorf("Actual %s is different from expected %s", boolToStr(actual), boolToStr(expected))
		}
	}
}

func calculateSomething(input int, size int) (int, bool) {
	output := input / size
	if output == 0 {
		return input, false
	}
	return input - size, true
}

func intToBools(input int, bits int) []bool {
	output := make([]bool, 0)
	for i := int(math.Pow(2.0, float64(bits-1))); i > 0; i = i / 2 {
		input2, outputBool := calculateSomething(input, i)
		input = input2
		output = append(output, outputBool)
	}
	return output
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

func dumpROM(rom *rom32k) []string {
	data := make([]string, 0)
	for _, r4k := range rom.rom0.ram4ks {
		for _, r512 := range r4k.ram512s {
			for _, r64 := range r512.ram64s {
				for _, r8 := range r64.ram8s {
					for _, r := range r8.registers {
						data = append(data, boolToStr(r.out))
					}
				}
			}
		}
	}
	for _, r4k := range rom.rom1.ram4ks {
		for _, r512 := range r4k.ram512s {
			for _, r64 := range r512.ram64s {
				for _, r8 := range r64.ram8s {
					for _, r := range r8.registers {
						data = append(data, boolToStr(r.out))
					}
				}
			}
		}
	}
	return data
}

func dumpRAM(mem *memory) []string {
	data := make([]string, 0)
	for _, r4k := range mem.ram.ram4ks {
		for _, r512 := range r4k.ram512s {
			for _, r64 := range r512.ram64s {
				for _, r8 := range r64.ram8s {
					for _, r := range r8.registers {
						data = append(data, boolToStr(r.out))
					}
				}
			}
		}
	}
	return data
}
