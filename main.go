package main

import "fmt"

func main() {
	vm := NewVM()
	vm.CurrentProgramBytecode = []uint8{
		0, 0, 1, 244, // load $0 #500
		0, 1, 1, 244, // load $1 #500
		1, 0, 1, 2,   // add $0 $1 $2 -> $2 = 1000
        2, 0, 1, 2,   // sub $0 $1 $2 -> $2 = 0
        3, 0, 1, 2,   // multiple $0 $1 $2 -> $2 = 250000
        4, 0, 1, 2,   // divide $0 $1 $2 -> $2 = 1
	}
	vm.Run()
	fmt.Println("Current register state:", vm.Registers)
}
 
