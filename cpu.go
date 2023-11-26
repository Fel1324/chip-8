package main

import (
	"fmt"

	"github.com/gaoliveira21/chip8/utils"
)

const (
	RAM_SIZE  = 4096 // 4 KB
	FREQUENCY = 700  // Instructions per second

	INSTRUCTION_BITMASK = 0xF000
	X_BITMASK           = 0x0F00
	Y_BITMASK           = 0x00F0
	N_BITMASK           = 0x000F
	NN_BITMASK          = 0x00FF
	NNN_BITMASK         = 0x0FFF
)

// Display
const (
	WIDTH  = 0x40
	HEIGHT = 0x20
)

type CPU struct {
	memory     [RAM_SIZE]uint8
	pc         uint16   // Program Counter
	i          uint16   // I Register
	v          [16]byte // Variable registers
	stack      utils.Stack
	display    [HEIGHT][WIDTH]byte
	delayTimer uint8
	soundTimer uint8
}

type opcode struct {
	instruction uint16
	registerX   uint8
	registerY   uint8
	n           uint8
	nn          uint8
	nnn         uint16
}

func (cpu *CPU) loadFont() {
	for i := 0x050; i <= 0x09F; i++ {
		cpu.memory[i] = fontdata[i-0x050]
	}
}

func NewCpu() CPU {
	cpu := CPU{
		pc: 0x200,
	}

	cpu.loadFont()

	return cpu
}

func (cpu *CPU) Fetch() uint16 {
	hb := uint16(cpu.memory[cpu.pc])
	cpu.pc++

	lb := uint16(cpu.memory[cpu.pc])
	cpu.pc++

	return (hb << 8) | lb
}

func (cpu *CPU) Decode(data uint16) (oc *opcode) {

	return &opcode{
		instruction: data & INSTRUCTION_BITMASK,
		registerX:   uint8((data & X_BITMASK) >> 8),
		registerY:   uint8((data & Y_BITMASK) >> 4),
		n:           uint8(data & N_BITMASK),
		nn:          uint8(data & NN_BITMASK),
		nnn:         (data & NNN_BITMASK),
	}
}

func (cpu *CPU) Clock() {
	data := cpu.Fetch()

	opcode := cpu.Decode(data)

	fmt.Printf("%X", opcode)

	switch opcode.instruction {
	case 0x000:
		switch opcode.n {
		case 0x00:
			fmt.Print("CLS")

		case 0x0E:
			fmt.Print("RET")

		default:
			fmt.Print("sys")
		}
	}
}