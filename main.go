package main


import ("os"
    "strconv"
	"fmt"
	"pushSwap/TritroisNombres"
	"pushSwap/InstructionDeTrie"
    )



func main(){

	
    
argument := os.Args[1:]

if len(argument)== 0 {
	return
	}
for _, i := range argument {
    nombre, err := strconv.Atoi(i)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
		}
	InstructionDeTrie.A = append(InstructionDeTrie.A, nombre)
    }
for i := 0 ; i < len(InstructionDeTrie.A); i ++{
	for k := i+1 ; k< len(InstructionDeTrie.A) ; k++{
		if InstructionDeTrie.A[i] == InstructionDeTrie.A[k] {
            fmt.Println("doublons de nombre :", InstructionDeTrie.A[i] )
		    fmt.Fprintln(os.Stderr, "Error")
		    return
		    }	
	    }
    }
	if len(InstructionDeTrie.A) == 3 {
		TritroisNombres.TritroisNombres()
	}
	for _, gestes := range InstructionDeTrie.Gestes{
		fmt.Println(gestes)
	}
}


