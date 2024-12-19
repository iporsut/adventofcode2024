//go:build ignore

package main

import (
	"fmt"
	"math/big"
	"slices"
	"strings"
)

type Machine struct {
	A, B, C *big.Int // Registers
	PC      int64    // Program Counter
	Prog    []int64

	OutBuffer []int64
}

func (m *Machine) adv(operand int64) {
	bOp := big.NewInt(int64(operand))

	switch operand {
	case 4:
		bOp = m.A
	case 5:
		bOp = m.B
	case 6:
		bOp = m.C
	}

	m.A = new(big.Int).Div(m.A, new(big.Int).Exp(big.NewInt(2), bOp, nil))
	m.PC += 2
}

func (m *Machine) bxl(operand int64) {
	bOp := big.NewInt(int64(operand))

	m.B = new(big.Int).Xor(m.B, bOp)
	m.PC += 2
}

func (m *Machine) bst(operand int64) {
	bOp := big.NewInt(int64(operand))
	switch operand {
	case 4:
		bOp = m.A
	case 5:
		bOp = m.B
	case 6:
		bOp = m.C
	}
	m.B = new(big.Int).Mod(bOp, big.NewInt(8))
	m.PC += 2
}

func (m *Machine) jnz(operand int64) {
	if m.A.Cmp(big.NewInt(0)) != 0 {
		m.PC = operand
	} else {
		m.PC += 2
	}
}

func (m *Machine) bxc(operand int64) {
	m.B = new(big.Int).Xor(m.B, m.C)
	m.PC += 2
}

func (m *Machine) out(operand int64) {
	bOp := big.NewInt(int64(operand))
	switch operand {
	case 4:
		bOp = m.A
	case 5:
		bOp = m.B
	case 6:
		bOp = m.C
	}

	bOp = new(big.Int).Mod(bOp, big.NewInt(8))
	m.OutBuffer = append(m.OutBuffer, bOp.Int64())
	m.PC += 2
}

func (m *Machine) bdv(operand int64) {
	bOp := big.NewInt(int64(operand))

	switch operand {
	case 4:
		bOp = m.A
	case 5:
		bOp = m.B
	case 6:
		bOp = m.C
	}

	m.B = new(big.Int).Div(m.A, new(big.Int).Exp(big.NewInt(2), bOp, nil))
	m.PC += 2
}

func (m *Machine) cdv(operand int64) {
	bOp := big.NewInt(int64(operand))

	switch operand {
	case 4:
		bOp = m.A
	case 5:
		bOp = m.B
	case 6:
		bOp = m.C
	}

	m.C = new(big.Int).Div(m.A, new(big.Int).Exp(big.NewInt(2), bOp, nil))
	m.PC += 2
}

func (m *Machine) flush() {
	outs := []string{}
	for _, v := range m.OutBuffer {
		outs = append(outs, fmt.Sprintf("%d", v))
	}
	fmt.Println(strings.Join(outs, ","))
}

func (m *Machine) printProg() {
	for i := 0; i < len(m.Prog); i += 2 {
		var opCode, operand int64
		opCode = m.Prog[i]
		operand = m.Prog[i+1]
		switch opCode {
		case 0:
			fmt.Printf("%d: adv %d\n", i, operand)
		case 1:
			fmt.Printf("%d: bxl %d\n", i, operand)
		case 2:
			fmt.Printf("%d: bst %d\n", i, operand)
		case 3:
			fmt.Printf("%d: jnz %d\n", i, operand)
		case 4:
			fmt.Printf("%d: bxc %d\n", i, operand)
		case 5:
			fmt.Printf("%d: out %d\n", i, operand)
		case 6:
			fmt.Printf("%d: bdv %d\n", i, operand)
		case 7:
			fmt.Printf("%d: cdv %d\n", i, operand)
		}
	}
}

func (m *Machine) run() bool {
	for m.PC < int64(len(m.Prog)) {
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

		// m.printRegisters()
	}

	return false
}

func (m *Machine) printRegisters() {
	fmt.Printf("Register A: %d\n", m.A)
	fmt.Printf("Register B: %d\n", m.B)
	fmt.Printf("Register C: %d\n", m.C)
	fmt.Printf("\n")
}

func (m *Machine) reset() {
	m.A = big.NewInt(0)
	m.B = big.NewInt(0)
	m.C = big.NewInt(0)
	m.PC = 0
	m.OutBuffer = []int64{}
}

func main() {
	m := Machine{
		Prog: []int64{2, 4, 1, 1, 7, 5, 0, 3, 1, 4, 4, 4, 5, 5, 3, 0},
		A:    big.NewInt(11),
		B:    big.NewInt(0),
		C:    big.NewInt(0),
		PC:   0,
	}
	m.printProg()

	vs := []*big.Int{big.NewInt(0)}

	for c := len(m.Prog) - 1; c >= 0; {
		found := false
		newVS := []*big.Int{}
		for _, v := range vs {
			for i := 0; i < 8; i++ {
				m.reset()
				a := new(big.Int).Add(v, big.NewInt(int64(i)))
				m.A = a
				m.run()

				if slices.Compare(m.OutBuffer, m.Prog[c:]) == 0 {
					found = true
					fmt.Println("a ", a, m.OutBuffer)
					newVS = append(newVS, a.Mul(a, big.NewInt(8)))
				}
			}
		}
		if found {
			c--
		}
		fmt.Println(newVS)
		vs = newVS
	}

}

var ex1 = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

var input = `Register A: 30886132
Register B: 0
Register C: 0

Program: 2,4,1,1,7,5,0,3,1,4,4,4,5,5,3,0`
