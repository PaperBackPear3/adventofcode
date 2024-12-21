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
	lines := utils.ReadFileLineByLine("./input.txt")

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
	part1(computer)
}

func part1(computer Computer) {
	stringOut := ""
	for internalPointer := 0; internalPointer < len(computer.instruction); internalPointer++ {
		instruction := computer.instruction[internalPointer]
		instructionValue := computer.value[internalPointer]
		println("\033[32m", "internalPointer:", internalPointer, "instruction:", instruction, "value:", instructionValue, "\033[0m")
		switch instruction {
		case 0:
			computer.registerDv("A", instructionValue)
		case 1:
			computer.bxl(instructionValue)
		case 2:
			computer.bst(instructionValue)
		case 3:
			jumpTo := computer.jnz(instructionValue)
			if jumpTo >= -1 {
				internalPointer = jumpTo
			}
		case 4:
			computer.bxc(instructionValue)
		case 5:
			toPrint := computer.out(instructionValue)
			stringOut = stringOut + utils.Itoa(toPrint)
		case 6:
			computer.registerDv("B", instructionValue)
		case 7:
			computer.registerDv("C", instructionValue)
		default:
			panic("unexpected instruction")
		}
	}
	println("stringOut: ", stringOut)
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
	println("\033[31m" + "bins:" + binNum + " " + binOperand + "\033[0m")

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

// done jumps
func (computer *Computer) jnz(code int) int {
	if computer.registers["A"] != 0 {
		return code - 1
	} else {
		return -2
	}
}

// bst with compo operator
func (computer *Computer) out(code int) int {
	val, _ := computer.getComboFromCode(code)
	res := val % 8
	println("out: ", res, " code: ", code)
	return res
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
