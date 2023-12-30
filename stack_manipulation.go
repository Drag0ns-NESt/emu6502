package emu6502

const STACK_BOTTOM = 0x100

// pushToStack pushes a given value to stack
func (cpu *CPU6502) pushToStack(value uint8) {
	cpu.Memory[STACK_BOTTOM+uint16(cpu.SP)] = value
	cpu.SP--
}

// pullFromStack pulls a top value from stack
func (cpu *CPU6502) pullFromStack() uint8 {
	cpu.SP++
	return cpu.Memory[STACK_BOTTOM+uint16(cpu.SP)]
}

// pushToStack16 pushes two bytes to stack. Higher byte will be pushed first
func (cpu *CPU6502) pushToStack16(value uint16) {
	// Push PC High Byte onto Stack
	cpu.pushToStack(uint8((value >> 8) & 0xFF))

	// Push PC Low Byte onto Stack
	cpu.pushToStack(uint8(value & 0xFF))
}

// brk executes BRK instruction
func (cpu *CPU6502) brk() {
	// coauthored by chatGPT
	// Increment PC to point to the next instruction after BRK
	cpu.PC++

	cpu.pushToStack16(cpu.PC)
	// Set Break flag (B) and push status onto Stack
	cpu.pushToStack(cpu.cpuStatusToByte() | 0x10)

	// Set Interrupt Disable (I) flag
	cpu.I = true

	// Load Interrupt Vector (0xFFFE/F) into PC for Interrupt Service Routine (ISR)
	cpu.PC = uint16(cpu.Memory[0xFFFF])<<8 | uint16(cpu.Memory[0xFFFE])
}

// php executes PHP instruction pushing status register to stack
func (cpu *CPU6502) php() {
	cpu.PC += 1
	cpu.pushToStack(cpu.cpuStatusToByte())
}

// plp executes PLP instruction pulling new status register value from a stack
func (cpu *CPU6502) plp() {
	cpu.PC += 1
	cpu.setCPUStatus(cpu.pullFromStack())
}
