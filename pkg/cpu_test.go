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

func TestMovingSequence2(t *testing.T) {
	cpu := NewCPU6502()

	cpu.PC = 0x600
	bytecode := []byte{
		0xa2, 0x01, 0xa9, 0xaa,
		0x95, 0xa0, 0xe8, 0x95,
		0xa0,
	}

	LoadAndExecute(cpu, bytecode)

	cpu.assertAEquals(t, 0xaa)
	cpu.assertXEquals(t, 0x02)
	cpu.assertPCEquals(t, 0x609)
	cpu.assertMemoryEquals(t, 0x00a0, 0x00, 0xaa, 0xaa)
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

func TestRelativeAddressingSequence(t *testing.T) {
	cpu := setupCPU()

	LoadAndExecute(cpu, []byte{
		0xa9, 0x01, 0xc9, 0x02,
		0xd0, 0x02, 0x85, 0x22,
		0x00,
	})

	cpu.assertAEquals(t, 0x01)
	cpu.assertNegativeEquals(t, true)
	cpu.assertPCEquals(t, 0x0609)

	cpu.assertMemoryEquals(t, 0x0022, 0x00)
}

func TestIndirectAddressingSequence(t *testing.T) {
	cpu := setupCPU()

	LoadAndExecute(cpu, []byte{
		0xa9, 0x01, 0x85, 0xf0,
		0xa9, 0xcc, 0x85, 0xf1,
		0x6c, 0xf0, 0x00,
	})

	cpu.assertAEquals(t, 0xcc)
	cpu.assertPCEquals(t, 0xcc01)

	cpu.assertMemoryEquals(t, 0x00f0, 0x01, 0xcc)
}

func TestIndexedIndirectAddressingSequence(t *testing.T) {
	cpu := setupCPU()

	LoadAndExecute(cpu, []byte{
		0xa2, 0x01, 0xa9, 0x05,
		0x85, 0x01, 0xa9, 0x07,
		0x85, 0x02, 0xa0, 0x0a,
		0x8c, 0x05, 0x07, 0xa1,
		0x00,
	})

	cpu.assertAEquals(t, 0x0a)
	cpu.assertXEquals(t, 0x01)
	cpu.assertYEquals(t, 0x0a)
	cpu.assertPCEquals(t, 0x0611)

	cpu.assertMemoryEquals(t, 0x00, 0x00, 0x05, 0x07)
	cpu.assertMemoryEquals(t, 0x0a)
}
