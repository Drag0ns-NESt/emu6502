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
	// 0x3-
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
	// 0x4-
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
