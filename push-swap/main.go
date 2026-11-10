package main

import (
	"fmt"
	"os"
	"pushSwap/InstructionDeTrie"
	"pushSwap/TritroisNombres"
	"pushSwap/TricinqNombres"
	"pushSwap/TriplusDeNombre"
	"strings"
	"strconv"
)

func main() {

	argument := os.Args[1:]

	if len(argument) == 0 {
		return
	}
	for _, i := range argument {
		nombre := strings.Fields(i)
		for k:=0 ; k < len(nombre) ; k++ {
			entier, err := strconv.Atoi(nombre[k])
			if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		    }
			InstructionDeTrie.A = append(InstructionDeTrie.A, entier)
	    }
	}
	for i := 0; i < len(InstructionDeTrie.A); i++ {
		for k := i + 1; k < len(InstructionDeTrie.A); k++ {
			if InstructionDeTrie.A[i] == InstructionDeTrie.A[k] {
				fmt.Fprintln(os.Stderr, "Error")
				return
			}
		}
	}
	
	if InstructionDeTrie.Estranger(){
		return
	}
	if len(InstructionDeTrie.A) <= 3 {
		TritroisNombres.TritroisNombres()
	} else if len(InstructionDeTrie.A) >= 4 && len(InstructionDeTrie.A) <= 5 {
		TricinqNombres.TricinqNombres()
	}else if len(InstructionDeTrie.A)>= 6 {
		TriplusDeNombre.TriplusDeNombre()
	}
	for _, gestes := range InstructionDeTrie.Gestes {
		fmt.Println(gestes)
	}
}

