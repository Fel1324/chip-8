package main

import (
	"testing"

	"github.com/gaoliveira21/chip8/utils"
)

func TestNewCpu(t *testing.T) {
	cpu := NewCpu()

	inMemoryFonts := [len(utils.Fontdata)]byte{}

	for i := 0x050; i <= 0x09F; i++ {
		font := cpu.mmu.Fetch(uint16(i))

		inMemoryFonts[i-0x050] = byte(font >> 8)
	}

	if cpu.pc != 0x200 {
		t.Errorf("cpu.pc = %d; expected 0x200", cpu.pc)
	}

	if inMemoryFonts != utils.Fontdata {
		t.Error("Error loading fonts")
	}
}

func TestDecode(t *testing.T) {
	cpu := NewCpu()

	opcode := cpu.Decode(0xABCD)

	var expected uint16 = 0xA000

	if opcode.instruction != expected {
		t.Errorf("opcode.instruction = 0x%X; expected 0x%X", opcode.instruction, expected)
	}

	expected = 0xB
	if opcode.registerX != uint8(expected) {
		t.Errorf("opcode.registerX = 0x%X; expected 0x%X", opcode.registerX, expected)
	}

	expected = 0xC
	if opcode.registerY != uint8(expected) {
		t.Errorf("opcode.registerY = 0x%X; expected 0x%X", opcode.registerY, expected)
	}

	expected = 0xD
	if opcode.n != uint8(expected) {
		t.Errorf("opcode.n = 0x%X; expected 0x%X", opcode.n, expected)
	}

	expected = 0xCD
	if opcode.nn != uint8(expected) {
		t.Errorf("opcode.n = 0x%X; expected 0x%X", opcode.nn, expected)
	}

	expected = 0xBCD
	if opcode.nnn != expected {
		t.Errorf("opcode.n = 0x%X; expected 0x%X", opcode.nnn, expected)
	}
}
