package TricinqNombres

import ("pushSwap/TritroisNombres"
"pushSwap/InstructionDeTrie"

)



func TricinqNombres(){

   
   

   for len(InstructionDeTrie.A) != 3 {

	PetitNombre := InstructionDeTrie.A[0]
    indexmin := 0
    
       for i:=1 ; i < len(InstructionDeTrie.A) ; i ++{
	        if  PetitNombre > InstructionDeTrie.A[i] {
		        PetitNombre = InstructionDeTrie.A[i]
		        indexmin = i
	        }
        }
	    for indexmin != 0 {
		    if indexmin <= len(InstructionDeTrie.A)/2 {
			    InstructionDeTrie.Ra()
			    indexmin = (indexmin - 1 + len(InstructionDeTrie.A)) % len(InstructionDeTrie.A)
		    } else {
			    InstructionDeTrie.Rra()
			    indexmin = (indexmin + 1) % len(InstructionDeTrie.A)
		    }
	    }
		InstructionDeTrie.Pb()
		
    }
	TritroisNombres.TritroisNombres()
	if InstructionDeTrie.Estranger(){
		for len(InstructionDeTrie.B) != 0{
			InstructionDeTrie.Pa()
		}
	}
}