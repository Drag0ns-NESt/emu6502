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

func setupCPU() *CPU6502 {
	cpu := NewCPU6502()
	cpu.PC = 0x600
	return cpu
}

func (cpu *CPU6502) assertAEquals(t *testing.T, value uint8) {
	if cpu.A != value {
		t.Errorf("Value in accumulator is expected to be equal to 0x%x. Actual: 0x%x", value, cpu.A)
	}
}

func (cpu *CPU6502) assertXEquals(t *testing.T, value uint8) {
	if cpu.X != value {
		t.Errorf("Value in X register is expected to be equal to 0x%x. Actual: 0x%x", value, cpu.X)
	}
}

func (cpu *CPU6502) assertPCEquals(t *testing.T, value uint16) {
	if cpu.PC != value {
		t.Errorf("Value in program counter is expected to be equal to 0x%x. Actual: 0x%x", value, cpu.PC)
	}
}

func (cpu *CPU6502) assertCarryEquals(t *testing.T, value bool) {
	if cpu.C != value {
		t.Errorf("Carry flag is expected to be equal to %t. Actual: %t", value, cpu.C)
	}
}

func (cpu *CPU6502) assertMemoryEquals(t *testing.T, start uint16, bytes ...byte) {
	for i, expected := range bytes {
		actual := cpu.Memory.Read(start + uint16(i))
		if actual != expected {
			t.Errorf("Value at 0x%X is expected to be equal to 0x%x. Actual: 0x%x", start+uint16(i), expected, actual)
		}
	}
}
