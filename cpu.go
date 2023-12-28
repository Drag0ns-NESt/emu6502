package emu6502

type CPU6502 struct {
	// A is accumulator register
	A uint8
	// X register
	X uint8
	// Y register
	Y uint8

	// PC is program counter. It stores an address of the next instruction to be
	// executed
	PC uint16

	// SP is stack pointer. It stores an address of current top of the stack
	SP uint8

	// C is carry flag. It is set whenever there are overflow happening in
	// arithmetic operations
	C bool

	// Z is zero flag. It is set after result of any operation is zero and unset
	// when it is not
	Z bool

	// D is decimal flag. Is is used to set processor to BCD (Binary-Coder Decimal)
	// mode for arithmetic operations. It is not used in current implementation and
	// added just FYI
	D bool

	// I is interrupt disable flag. It enables and disables CPU's ability to respond
	// to maskable interrupts.
	I bool

	// B is break flag. Indicates that a software interrupt was executed
	B bool

	// U is an extra unused bit in 6502 status register. Can be used for some things,
	// but added here just FYI
	U bool

	// V is signed overflow flag. Set whenever a signed arithmetic overlow happens.
	V bool

	// N is negative flag. Set whenever result of operation can be interpreted as
	// negative in context of signed arithmetic operations (most significant bit is set)
	N bool

	// Memory represents a memory used for storing
	Memory [0x10000]byte
}

func nop(cpu *CPU6502) {
	// do nothing, just increment program counter
	cpu.PC++
}

// asl performs Arithmetic Shift Left operation on the provided value
func (cpu *CPU6502) asl(value uint8) uint8 {
	// coauthored with chatGPT
	// Set the Carry flag if the most significant bit (bit 7) is 1, otherwise clear it
	cpu.C = (value & 0x80) != 0

	// Perform the shift left operation
	shiftedValue := value << 1

	// Update the Zero flag if the result is zero
	cpu.Z = shiftedValue == 0

	// Update the Negative flag based on bit 7 of the shifted value
	cpu.N = (shiftedValue & 0x80) != 0

	// Update program counter
	cpu.PC += 1
	return shiftedValue
}

// ora executes ORA instruction performing bitwise OR for A register and a
// given argument
func (cpu *CPU6502) ora(arg uint8) {
	// Performing the bitwise OR
	cpu.A |= arg

	// Setting zero flag if result is zero
	cpu.Z = cpu.A == 0

	// Setting negative flag if result can be interpreted as bitewise negative value
	// maximum bit is set to one
	cpu.N = (cpu.A & 0x80) != 0

	cpu.PC += 1
}

const STACK_BOTTOM = 0x100

// brk executes BRK instruction
func (cpu *CPU6502) brk() {
	// authored by chatGPT
	// Increment PC to point to the next instruction after BRK
	cpu.PC++

	// Push PC High Byte onto Stack
	cpu.Memory[STACK_BOTTOM + uint16(cpu.SP)] = uint8((cpu.PC >> 8) & 0xFF)
	cpu.SP--

	// Push PC Low Byte onto Stack
	cpu.Memory[STACK_BOTTOM + 0x0100+uint16(cpu.SP)] = uint8(cpu.PC & 0xFF)
	cpu.SP--

	// Set Break flag (B) and push status onto Stack
	cpu.Memory[STACK_BOTTOM + 0x0100+uint16(cpu.SP)] = cpuStatusToByte(cpu) | 0x10 // Set B flag
	cpu.SP--

	// Set Interrupt Disable (I) flag
	cpu.I = true

	// Load Interrupt Vector (0xFFFE/F) into PC for Interrupt Service Routine (ISR)
	cpu.PC = uint16(cpu.Memory[0xFFFF])<<8 | uint16(cpu.Memory[0xFFFE])
}

// php executes PHP instruction pushing status register to stack
func (cpu *CPU6502) php() {
	cpu.PC += 1
	
	cpu.Memory[]
}

// NewCPU6502 creates and initializes new 6502 CPU emulator instance
func NewCPU6502() *CPU6502 {
	return &CPU6502{
		A:      0,
		X:      0,
		Y:      0,
		PC:     0,
		SP:     0xFF, // Initialize stack pointer to 0xFF (top of stack)
		C:      false,
		Z:      false,
		I:      false,
		D:      false,
		V:      false,
		N:      false,
		Memory: [0x10000]byte{},
	}
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
