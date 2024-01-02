package emu6502

// adc performs ADC (ADd with Carry) operation for value in accumulator with a
// given argument
func (cpu *CPU6502) adc(arg uint8) {
	cpu.PC += 1
	result := uint16(cpu.A) + uint16(arg)
	if cpu.C {
		result += 1
	}

	// if result is can't fit byte set carry flag
	cpu.C = result > 0xff

	// set overflow (V) flag if result overflowed to 7-bit (there were not 7 bit
	// set for operands but it is for result)
	cpu.V = ((uint8(result) & 0x80) & ^((cpu.A & 0x80) | (arg & 0x80))) != 0
	cpu.A = uint8(result & 0xff)

	// set zero flag if result is 0
	cpu.Z = cpu.A == 0

	// set negative flag if result can be interpreted as signed negative number
	cpu.N = (cpu.A & 0x80) != 0

	// TODO: decimal mode. mb someday...
}

// dey performs DEY (DEcrement Y) operation
func (cpu *CPU6502) dey() {
	cpu.PC += 1
	cpu.Y -= 1

	cpu.Z = cpu.Y == 0
	cpu.N = (cpu.Y & 0x80) != 0
}
