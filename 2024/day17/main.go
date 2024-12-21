package main

import (
	"adventofcode-2024/utils"
	"errors"
	"math"
)

type Computer struct {
	registers   map[string]int
	instruction []int
	value       []int
}

func main() {
	lines := utils.ReadFileLineByLine("./test.txt")

	computer := Computer{registers: map[string]int{}}

	for _, line := range lines {
		if line != "" {
			x := utils.SplitLineMultipleDelimiters(line, []string{"Register ", ": ", ","}, true)
			if x[0] != "Program" {
				computer.registers[x[0]] = utils.Atoi(x[1])
			} else {
				for index, val := range x {
					if index == 0 {
						continue
					}
					if index%2 == 1 {
						computer.instruction = append(computer.instruction, utils.Atoi(val))
					} else {
						computer.value = append(computer.value, utils.Atoi(val))
					}
				}
			}
		}
	}
	println(computer.instruction)
	part1(computer)
}

func part1(computer Computer) {
	for index, instruction := range computer.instruction {
		instructionValue := computer.value[index]
		switch instruction {
		case 0:
			computer.registerDv("A", instructionValue)
		case 1:
			computer.bxl(instructionValue)
		case 2:
			computer.bst(instructionValue)
		case 3:
			computer.jnz(instructionValue)
		case 4:
			computer.bxc(instructionValue)
		case 5:
			computer.out(instructionValue)
		case 6:
			computer.registerDv("B", instructionValue)
		case 7:
			computer.registerDv("C", instructionValue)
		default:
			panic("unexpected instruction")
		}
	}
}

func (computer *Computer) getComboFromCode(code int) (int, error) {
	if code < 4 {
		return code, nil
	}
	if code == 4 {
		return computer.registers["A"], nil
	}
	if code == 5 {
		return computer.registers["B"], nil
	}
	if code == 6 {
		return computer.registers["C"], nil
	}
	if code == 7 {
		return 0, errors.New("code 7 invalid")
	}
	return 0, errors.New("wtf?")
}

// done divide and save in a register
func (computer *Computer) registerDv(registerToSave string, code int) {
	num := computer.registers["A"]
	val, _ := computer.getComboFromCode(code)
	den := math.Pow(2, float64(val))
	computer.registers[registerToSave] = num / int(den)
}

// done bitwise Xor with litteral value
func (computer *Computer) bxl(code int) int {
	num := computer.registers["B"]
	binNum := toBin(num)
	operand := code
	binOperand := toBin(operand)
	println(binNum, binOperand)

	longest := binNum
	shortest := binOperand
	diff := len(longest) - len(shortest)
	offset := diff
	bxValue := ""
	if diff < 0 {
		tmp := longest
		longest = shortest
		shortest = tmp
		diff = diff * -1
		offset = offset * -1
	}
	for index, char := range longest {
		if diff > 0 {
			if char != '0' {
				bxValue = bxValue + "1"
			} else {
				bxValue = bxValue + "0"
			}
			diff--
		} else {
			if char != rune(shortest[index-offset]) {
				bxValue = bxValue + "1"
			} else {
				bxValue = bxValue + "0"
			}
		}
	}
	dec := toDec(bxValue)

	computer.registers["B"] = dec
	return dec
}

// done bitwise Xor with value from register C
func (computer *Computer) bxc(code int) {
	res := computer.bxl(computer.registers["C"])
	computer.registers["B"] = res
}

// done mod 8 of the value
func (computer *Computer) bst(code int) {
	val, _ := computer.getComboFromCode(code)
	res := val % 8
	computer.registers["B"] = res
}

// todo jumps
func (computer *Computer) jnz(code int) {
	if computer.registers["A"] != 0 {
		//find a way to jump by code
	}
}

// bst with compo operator
func (computer *Computer) out(code int) {
	val, _ := computer.getComboFromCode(code)
	res := val % 8
	println(res)
}

func toBin(n int) string {
	bin := ""
	for {
		if n == 0 {
			if bin == "" {
				bin = "0"
			}
			break
		}
		if n%2 == 0 {
			bin = bin + "0"
		} else {
			bin = bin + "1"
		}
		n = n / 2
	}
	reversedBin := ""
	for i := len(bin) - 1; i >= 0; i-- {
		reversedBin += string(bin[i])
	}
	bin = reversedBin
	return bin
}

func toDec(n string) int {
	dec := 0
	base := 1
	length := len(n)
	for i := length - 1; i >= 0; i-- {
		if n[i] == '1' {
			dec += base
		}
		base *= 2
	}
	return dec
}
