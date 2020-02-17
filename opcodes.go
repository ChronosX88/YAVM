package main

type Opcode int

const (
  // LOAD is opcode for loading number into register
  // Args: 
  // 1. register for loading number to
  // 2. two bytes for number itself
  LOAD Opcode = 0

  // ADD is opcode for adding two number from register and storing to register which specified in bytecode
  // Args:
  // 1. first number from register
  // 2. second number from register
  // 3. to which register need to store result of adding
  ADD Opcode = 1
  SUB Opcode = 2
  MULTIPLE Opcode = 3
  DIVIDE Opcode = 4
  HALT Opcode = 45
  ILLEGAL Opcode = 50
)

func OpcodeFrom(b uint8) Opcode {
  switch b {
    case 0: {
      return LOAD
    }
    case 1: {
      return ADD
    }
    case 2: {
      return SUB
    }
    case 3: {
      return MULTIPLE
    }
    case 4: {
      return DIVIDE
    }
    case 45: {
      return HALT
    }
    default: {
      return ILLEGAL
    }
  }
} 
