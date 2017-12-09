package day8

import (
	"strings"
	"regexp"
	"strconv"
	"fmt"
)

func Day8Part1(input string) (int, int) {
	instructions := parseInstructions(input)

	registers := map[string]int{}
	var allMax int

	for _, instr := range instructions {
		if !instr.condition.evaluate(registers) {
			continue
		}

		switch instr.instructionType {
		case INCREMENT:
			registers[instr.register] = registers[instr.register] + instr.value
		case DECREMENT:
			registers[instr.register] = registers[instr.register] - instr.value
		}

		if registers[instr.register] > allMax {
			allMax = registers[instr.register]
		}
	}

	var max int

	for _, v := range registers {
		if v > max {
			max = v
		}
	}

	return max, allMax
}

func parseInstructions(input string) []*instruction {
	r := regexp.MustCompile("^(\\w+) (inc|dec) (-?\\d+) if (\\w+) (>|>=|<|<=|==|!=) (-?\\d+)$")

	lines := strings.Split(input, "\n")
	instructions := make([]*instruction, len(lines))

	for i, line := range lines {
		result := r.FindStringSubmatch(line)

		if result == nil {
			panic("no regex match for " + line)
		}

		var instructionType instructionType
		switch result[2] {
		case "inc":
			instructionType = INCREMENT
		case "dec":
			instructionType = DECREMENT
		default:
			panic("unknown instruction type " + result[2])
		}

		instructionValue, err := strconv.Atoi(result[3])
		if err != nil {
			panic(err)
		}

		var conditionType conditionType
		switch result[5] {
		case ">":
			conditionType = LARGER
		case "<":
			conditionType = LESS
		case ">=":
			conditionType = LARGER_EQUAL
		case "<=":
			conditionType = LESS_EQUAL
		case "==":
			conditionType = EQUAL
		case "!=":
			conditionType = NOT_EQUAL
		default:
			panic("unknown condition type " + result[5])
		}

		conditionValue, err := strconv.Atoi(result[6])
		if err != nil {
			panic(err)
		}

		instructions[i] = &instruction{
			register:        result[1],
			instructionType: instructionType,
			value:           instructionValue,
			condition: &condition{
				register:      result[4],
				conditionType: conditionType,
				value:         conditionValue,
			},
		}
	}

	return instructions
}

type instruction struct {
	register        string
	instructionType instructionType
	value           int
	condition       *condition
}

func (i *instruction) String() string {
	return fmt.Sprintf("%s %s %d if %s %s %d", i.register, i.instructionType, i.value, i.condition.register, i.condition.conditionType, i.condition.value)
}

type condition struct {
	register      string
	conditionType conditionType
	value         int
}

func (c *condition) evaluate(registers map[string]int) bool {
	v, ok := registers[c.register]
	if !ok {
		v = 0
	}

	switch c.conditionType {
	case LARGER:
		return v > c.value
	case LESS:
		return v < c.value
	case LARGER_EQUAL:
		return v >= c.value
	case LESS_EQUAL:
		return v <= c.value
	case EQUAL:
		return v == c.value
	case NOT_EQUAL:
		return v != c.value
	}

	panic("unknown condition type " + strconv.Itoa(int(c.conditionType)))
}

type instructionType int

const (
	INCREMENT instructionType = iota
	DECREMENT
)

func (i instructionType) String() string {
	switch i {
	case INCREMENT:
		return "inc"
	case DECREMENT:
		return "dec"
	}
	panic("unknown instruction type " + strconv.Itoa(int(i)))
}

type conditionType int

const (
	LARGER       conditionType = iota
	LESS
	LARGER_EQUAL
	LESS_EQUAL
	EQUAL
	NOT_EQUAL
)

func (c conditionType) String() string {
	switch c {
	case LARGER:
		return ">"
	case LESS:
		return "<"
	case LARGER_EQUAL:
		return ">="
	case LESS_EQUAL:
		return "<="
	case EQUAL:
		return "=="
	case NOT_EQUAL:
		return "!="
	}
	panic("unknown condition type " + strconv.Itoa(int(c)))
}
