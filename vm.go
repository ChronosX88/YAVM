package main

import "fmt"

type VM struct {
  Registers []int32
  CurrentProgramBytecode []uint8
  CurrentProgramPos int
  Remainder uint32
}

func NewVM() *VM {
  return &VM {
    Registers: make([]int32, 32),
    CurrentProgramBytecode: make([]uint8, 0),
    CurrentProgramPos: 0,
    Remainder: 0,
  }
}

func (vm *VM) Run() {
  for {
    if vm.CurrentProgramPos >= len(vm.CurrentProgramBytecode) {
      break
    }

    switch vm.decodeOpcode() {
      case HALT: {
        fmt.Println("HALT encountered")
        return
      }
      case LOAD: {
        register := uint(vm.next8Bits())
        number := uint32(vm.next16Bits())
        fmt.Println(register, number)
        vm.Registers[register] = int32(number)
      }
      case ADD: {
        number1 := vm.Registers[uint(vm.next8Bits())]
        number2 := vm.Registers[uint(vm.next8Bits())]
        targetRegister := vm.next8Bits()
        vm.Registers[targetRegister] = number1 + number2
      }
      case SUB: {
        number1 := vm.Registers[uint(vm.next8Bits())]
        number2 := vm.Registers[uint(vm.next8Bits())]
        targetRegister := vm.next8Bits()
        vm.Registers[targetRegister] = number1 - number2
      }
      case MULTIPLE: {
        number1 := vm.Registers[uint(vm.next8Bits())]
        number2 := vm.Registers[uint(vm.next8Bits())]
        targetRegister := vm.next8Bits()
        vm.Registers[targetRegister] = number1 * number2
      }
      case DIVIDE: {
        number1 := vm.Registers[uint(vm.next8Bits())]
        number2 := vm.Registers[uint(vm.next8Bits())]
        targetRegister := vm.next8Bits()
        vm.Registers[targetRegister] = number1 / number2
        vm.Remainder = uint32(number1 % number2)
      }
      default: {
        fmt.Println("Unrecognized opcode, terminating VM...")
        return
      }
    }
  }
}

func (vm *VM) decodeOpcode() Opcode {
  opcode := OpcodeFrom(vm.CurrentProgramBytecode[vm.CurrentProgramPos])
  vm.CurrentProgramPos++
  return opcode
}

func (vm *VM) next8Bits() uint8 {
  result := vm.CurrentProgramBytecode[vm.CurrentProgramPos]
  vm.CurrentProgramPos++
  return result
}

func (vm *VM) next16Bits() uint16 {
  result := ((uint16(vm.CurrentProgramBytecode[vm.CurrentProgramPos])) << 8) | uint16(vm.CurrentProgramBytecode[vm.CurrentProgramPos + 1])
  vm.CurrentProgramPos += 2
  return result
} 
