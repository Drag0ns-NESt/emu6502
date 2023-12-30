package emu6502

var opcodeFunctions [0x100]func(cpu *CPU6502) = [0x100]func(cpu *CPU6502){
	// 0x00 BRK
	func(cpu *CPU6502) {
		cpu.brk()
	},
	// 0x01 ORA, (indirect, X)
	func(cpu *CPU6502) {
		cpu.ora(cpu.indexedIndirect())
	},
	// 0x02 is not defined, assign NOP function
	nop,
	// 0x03 is not defined, assign NOP function
	nop,
	// 0x04 is not defined, assign NOP function
	nop,
	// 0x05 ORA, zero page
	func(cpu *CPU6502) {
		cpu.ora(cpu.zeroPage())
	},
	// 0x06 ASL, zero page
	func(cpu *CPU6502) {
		cpu.executeWithZeroPage(cpu.asl)
	},
	// 0x07 is not defined, assign NOP function
	nop,
	// 0x08 PHP
	func(cpu *CPU6502) {
		cpu.php()
	},
	// 0x09 ORA, immediate
	func(cpu *CPU6502) {
		cpu.ora(cpu.immediate())
	},
	// 0x0a ASL, accumulator
	func(cpu *CPU6502) {
		cpu.executeWithAccumulator(cpu.asl)
	},
	// 0x0b is not defined, assign NOP function
	nop,
	// 0x0c is not defined, assign NOP function
	nop,
	// 0x0d ORA, absolute
	func(cpu *CPU6502) {
		cpu.ora(cpu.absolute())
	},
	// 0x0e ASL, absolute
	func(cpu *CPU6502) {
		cpu.executeWithAbsolute(cpu.asl)
	},
	// 0x0f is not defined, assign NOP function
	nop,

	// 0x10 BPL, relative
	func(cpu *CPU6502) {
		cpu.bpl(cpu.relative())
	},
	// 0x11 ORA, inderect indexed: (address), Y
	func(cpu *CPU6502) {
		cpu.ora(cpu.indirectIndexed())
	},
	// 0x12 is not defined, assign NOP function
	nop,
	// 0x13 is not defined, assign NOP function
	nop,
	// 0x14 is not defined, assign NOP function
	nop,
	// 0x15 ORA, zero-page, X
	func(cpu *CPU6502) {
		cpu.ora(cpu.zeroPageX())
	},
	// 0x16 ASL, zero-page, X
	func(cpu *CPU6502) {
		cpu.executeWithZeroPageX(cpu.asl)
	},
	// 0x17 is not defined, assign NOP function
	nop,
	// 0x18 CLC
	func(cpu *CPU6502) {
		cpu.clc()
	},
	// 0x19 ORA, absolute, Y
	func(cpu *CPU6502) {
		cpu.ora(cpu.absoluteY())
	},
	// 0x1a is not defined, assign NOP function
	nop,
	// 0x1b is not defined, assign NOP function
	nop,
	// 0x1c is not defined, assign NOP function
	nop,
	// 0x1d ORA, absolute, X
	func(cpu *CPU6502) {
		cpu.ora(cpu.absoluteX())
	},
	// 0x1e ASL, absolute, X
	func(cpu *CPU6502) {
		cpu.executeWithAbsoluteY(cpu.asl)
	},
	// 0x1f is not defined, assign NOP function
	nop,
	// 0x20 JSR, absolute
	func(cpu *CPU6502) {
		cpu.jsr(cpu.absoluteAddress())
	},
	// 0x21 AND, (indirect, X)
	func(cpu *CPU6502) {
		cpu.and(cpu.indexedIndirect())
	},
	// 0x22 is not defined, assign NOP function
	nop,
	// 0x23 is not defined, assign NOP function
	nop,
	// 0x24 BIT, zero page
	func(cpu *CPU6502) {
		cpu.bit(cpu.zeroPage())
	},
	// 0x25 AND, zero page
	func(cpu *CPU6502) {
		cpu.and(cpu.zeroPage())
	},
	// 0x26 ROL, zero page
	func(cpu *CPU6502) {
		cpu.executeWithZeroPage(cpu.rol)
	},
	// 0x27 is not defined, assign NOP function
	nop,
	// 0x28 PLP
	func(cpu *CPU6502) {
		cpu.plp()
	},
	// 0x29 AND, immediate
	func(cpu *CPU6502) {
		cpu.and(cpu.immediate())
	},
	// 0x2a ROL, accumulator
	func(cpu *CPU6502) {
		cpu.executeWithAccumulator(cpu.rol)
	},
	// 0x2b is not defined, assign NOP function
	nop,
	// 0x2c BIT, absolute
	func(cpu *CPU6502) {
		cpu.bit(cpu.absolute())
	},
	// 0x2d AND, absolute
	func(cpu *CPU6502) {
		cpu.and(cpu.absolute())
	},
	// 0x2e ROL, absolute
	func(cpu *CPU6502) {
		cpu.executeWithAbsolute(cpu.rol)
	},
	// 0x30 BMI, relative
	func(cpu *CPU6502) {
		cpu.bmi(cpu.relative())
	},
	// 0x31 AND, indirect indexed, (address), Y
	func(cpu *CPU6502) {
		cpu.and(cpu.indirectIndexed())
	},
	// 0x32 is not defined, assign NOP function
	nop,
	// 0x33 is not defined, assign NOP function
	nop,
	// 0x34 is not defined, assign NOP function
	nop,
	// 0x35 AND, zero page, X
	func(cpu *CPU6502) {
		cpu.and(cpu.zeroPageX())
	},
	// 0x36 ROL, zero page, X
	func(cpu *CPU6502) {
		cpu.executeWithZeroPageX(cpu.rol)
	},
	// 0x37 is not defined, assign NOP function
	nop,
	// 0x38 SEC
	func(cpu *CPU6502) {
		cpu.sec()
	},
	// 0x39 AND, absolute, Y
	func(cpu *CPU6502) {
		cpu.and(cpu.absoluteY())
	},
	// 0x3a is not defined, assign NOP function
	nop,
	// 0x3b is not defined, assign NOP function
	nop,
	// 0x3c is not defined, assign NOP function
	nop,
	// 0x3d AND, absolute, X
	func(cpu *CPU6502) {
		cpu.and(cpu.absoluteX())
	},
	// 0x3e ROL, absolute
	func(cpu *CPU6502) {
		cpu.executeWithAbsoluteX(cpu.rol)
	},
	// 0x3f is not defined, assign NOP function
	nop,

	// 0x40 RTI
	func(cpu *CPU6502) {
		cpu.rti()
	},
	// 0x41 EOR indexed indirect, (address, X)
	func(cpu *CPU6502) {
		cpu.eor(cpu.indexedIndirect())
	},
	// 0x42 is not defined, assign NOP function
	nop,
	// 0x43 is not defined, assign NOP function
	nop,
	// 0x44 is not defined, assign NOP function
	nop,
	// 0x45 EOR, zero page
	func(cpu *CPU6502) {
		cpu.eor(cpu.zeroPage())
	},
	// 0x46 LSR, zero page
	func(cpu *CPU6502) {
		cpu.executeWithZeroPage(cpu.lsr)
	},
	// 0x47 is not defined, assign NOP function
	nop,
	// 0x48 PHA
	func(cpu *CPU6502) {
		cpu.pha()
	},
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	// 0x5-
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	// 0x6-
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	// 0x7-
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	// 0x8-
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	// 0x9-
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	// 0xa-
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	// 0xb-
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	// 0xc-
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	// 0xd-
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	// 0xe-
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	// 0xf-
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
	nop,
}
