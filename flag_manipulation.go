package emu6502

// clc executes CLC instruction clearing carry flag
func (cpu *CPU6502) clc() {
	cpu.PC += 1
	cpu.C = false
}

// bit executes BIT instruction
func (cpu *CPU6502) bit(value uint8) {
	cpu.PC += 1
	cpu.N = value&(0x01<<7) == 1
	cpu.V = value&(0x01<<6) == 1

	cpu.Z = (value & cpu.A) == 0
}

// cpuStatusToByte converts CPU status to byte format
func cpuStatusToByte(cpu *CPU6502) uint8 {
	status := uint8(0x20) // Bit 5 is always set for the 6502
	if cpu.N {
		status |= 0x80
	}
	if cpu.V {
		status |= 0x40
	}
	if cpu.U {
		status |= 0x20
	}
	if cpu.B {
		status |= 0x10
	}
	if cpu.D {
		status |= 0x08
	}
	if cpu.I {
		status |= 0x04
	}
	if cpu.Z {
		status |= 0x02
	}
	if cpu.C {
		status |= 0x01
	}
	return status
}
