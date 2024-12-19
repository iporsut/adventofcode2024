//go:build ignore

package main

import (
	"fmt"
	"math"
	"strings"
)

type Machine struct {
	A, B, C int // Registers
	PC      int // Program Counter
	Prog    []int

	OutBuffer strings.Builder
}

func (m *Machine) adv(operand int) {
	switch operand {
	case 4:
		operand = m.A
	case 5:
		operand = m.B
	case 6:
		operand = m.C
	}
	m.A = m.A / int(math.Pow(2, float64(operand)))
	m.PC += 2
}

func (m *Machine) bxl(operand int) {
	m.B = m.B ^ operand
	m.PC += 2
}

func (m *Machine) bst(operand int) {
	switch operand {
	case 4:
		operand = m.A
	case 5:
		operand = m.B
	case 6:
		operand = m.C
	}
	m.B = operand % 8
	m.PC += 2
}

func (m *Machine) jnz(operand int) {
	if m.A != 0 {
		m.PC = operand
	} else {
		m.PC += 2
	}
}

func (m *Machine) bxc(operand int) {
	m.B = m.B ^ m.C
	m.PC += 2
}

func (m *Machine) out(operand int) {
	switch operand {
	case 4:
		operand = m.A
	case 5:
		operand = m.B
	case 6:
		operand = m.C
	}

	operand = operand % 8
	m.OutBuffer.WriteString(fmt.Sprintf("%d", operand))
	m.PC += 2
}

func (m *Machine) bdv(operand int) {
	switch operand {
	case 4:
		operand = m.A
	case 5:
		operand = m.B
	case 6:
		operand = m.C
	}
	m.B = m.A / int(math.Pow(2, float64(operand)))
	m.PC += 2
}

func (m *Machine) cdv(operand int) {
	switch operand {
	case 4:
		operand = m.A
	case 5:
		operand = m.B
	case 6:
		operand = m.C
	}
	m.C = m.A / int(math.Pow(2, float64(operand)))
	m.PC += 2
}

func (m *Machine) flush() {
	fmt.Println(m.OutBuffer.String())
}

func (m *Machine) run() {
	for m.PC < len(m.Prog) {
		switch m.Prog[m.PC] {
		case 0:
			m.adv(m.Prog[m.PC+1])
		case 1:
			m.bxl(m.Prog[m.PC+1])
		case 2:
			m.bst(m.Prog[m.PC+1])
		case 3:
			m.jnz(m.Prog[m.PC+1])
		case 4:
			m.bxc(m.Prog[m.PC+1])
		case 5:
			m.out(m.Prog[m.PC+1])
		case 6:
			m.bdv(m.Prog[m.PC+1])
		case 7:
			m.cdv(m.Prog[m.PC+1])
		}
	}
}

func main() {
	m := Machine{
		Prog: []int{0, 1, 5, 4, 3, 0},
		A:    729,
		B:    0,
		C:    0,
		PC:   0,
	}
	m.run()
	m.flush()
}

var ex1 = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

var input = `Register A: 30886132
Register B: 0
Register C: 0

Program: 2,4,1,1,7,5,0,3,1,4,4,4,5,5,3,0`
