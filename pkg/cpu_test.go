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
