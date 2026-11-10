package main

import ("bufio"
"os"
"fmt"
"strconv"
"pushSwap/InstructionDeTrie"
"pushSwap/TritroisNombres"
"strings" )

func main(){

	var instructions []string

    if len(os.Args) == 1 {
        return
	}

    scanner := bufio.NewScanner(os.Stdin)

	argument := os.Args[1:]

	for _, i := range argument {
		nombre := strings.Fields(i)
		for k:=0 ; k<len(nombre) ; k++{
			entier, err := strconv.Atoi(nombre[k])
            if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		    }
			InstructionDeTrie.A = append(InstructionDeTrie.A, entier)
	    }	
	}

	for scanner.Scan() {
	
		instructions = append(instructions, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}
	for l := 0; l < len(InstructionDeTrie.A); l++ {
		for m := l + 1; m < len(InstructionDeTrie.A); m++ {
			if InstructionDeTrie.A[l] == InstructionDeTrie.A[m] {
				fmt.Fprintln(os.Stderr, "Error")
				return
			}
		}
	}
	for n := 0 ; n < len (instructions) ; n++ {
		switch (instructions[n]){
		case "sa":
			InstructionDeTrie.Sa()
		case "ra":
			InstructionDeTrie.Ra()
		case "rra":
			InstructionDeTrie.Rra()
		case "pa":
			InstructionDeTrie.Pa()
		case "pb":
			InstructionDeTrie.Pb()
		case "rrr":
			InstructionDeTrie.Rrr()
		case "sb":
			InstructionDeTrie.Sb()
		case "rb":
			InstructionDeTrie.Rb()
		case "ss":
			InstructionDeTrie.Ss()
		case "rrb":
			InstructionDeTrie.Rrb()
		case "rr":
			InstructionDeTrie.Rr()
		default:
			fmt.Fprintln(os.Stderr, "Error")
		    return

		}
		
	}
	if TritroisNombres.Estranger() && len(InstructionDeTrie.B) == 0{
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}