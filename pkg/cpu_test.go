package emu6502

import (
	"testing"
)

// TestNOPInstruction tests the NOP instruction
func TestNOPInstruction(t *testing.T) {
	cpu := NewCPU6502()
	initialPC := cpu.PC

	// Writing NOP opcode (0xEA) to the current PC location
	cpu.Memory.Write(cpu.PC, 0x28)

	cpu.ExecuteNext()

	// The PC should increment by 1 after NOP
	if cpu.PC != initialPC+1 {
		t.Errorf("NOP did not increment PC correctly. Expected %v, got %v", initialPC+1, cpu.PC)
	}
}

func TestMovingSequence(t *testing.T) {
	cpu := NewCPU6502()

	cpu.PC = 0x600
	bytecode := []byte{
		0xa9, 0x01, 0x8d, 0x00,
		0x02, 0xa9, 0x05, 0x8d,
		0x01, 0x02, 0xa9, 0x08,
		0x8d, 0x02, 0x02}

	LoadAndExecute(cpu, bytecode)

	cpu.assertAEquals(t, 0x08)
	cpu.assertPCEquals(t, 0x60f)
	cpu.assertMemoryEquals(t, 0x200, 0x01, 0x05, 0x08)
}

func TestArithmetics(t *testing.T) {
	cpu := setupCPU()

	LoadAndExecute(cpu, []byte{0xa9, 0xc0, 0xaa, 0xe8, 0x69, 0xc4, 0x00})

	cpu.assertAEquals(t, 0x84)
	cpu.assertXEquals(t, 0xc1)
	cpu.assertPCEquals(t, 0x0607)
	cpu.assertCarryEquals(t, true)
}

func TestBranching(t *testing.T) {
	cpu := setupCPU()

	LoadAndExecute(cpu, []byte{
		0xa2, 0x08, 0xca, 0x8e,
		0x00, 0x02, 0xe0, 0x03,
		0xd0, 0xf8, 0x8e, 0x01,
		0x02, 0x00,
	})

	cpu.assertXEquals(t, 0x03)
	cpu.assertPCEquals(t, 0x060e)
	cpu.assertCarryEquals(t, true)
	cpu.assertZeroEquals(t, true)

	cpu.assertMemoryEquals(t, 0x0200, 0x03, 0x03)
}
