package TritroisNombres

import ("pushSwap/InstructionDeTrie")

func TritroisNombres(){

	if len(InstructionDeTrie.A) == 2 && !InstructionDeTrie.Estranger() {
    InstructionDeTrie.Sa()
        }
	if len(InstructionDeTrie.A) == 3{
        if InstructionDeTrie.Estranger(){
		    
			return
	    } 
	    copie := make([]int,len(InstructionDeTrie.A))
	    copy(copie, InstructionDeTrie.A)
		debutGestes := len(InstructionDeTrie.Gestes)
	    InstructionDeTrie.Sa()
	    InstructionDeTrie.Rra()

	    if InstructionDeTrie.Estranger(){
		    
			return
	    }
	    copy(InstructionDeTrie.A, copie)
		InstructionDeTrie.Gestes = InstructionDeTrie.Gestes[:debutGestes]
	    InstructionDeTrie.Sa()

	    if InstructionDeTrie.Estranger(){
		   
			return
	    }
	    copy(InstructionDeTrie.A, copie)
		InstructionDeTrie.Gestes = InstructionDeTrie.Gestes[:debutGestes]
	    InstructionDeTrie.Ra()
     
	    if InstructionDeTrie.Estranger(){
		   
			return
	    }
	    copy(InstructionDeTrie.A, copie)
		InstructionDeTrie.Gestes = InstructionDeTrie.Gestes[:debutGestes]
	    InstructionDeTrie.Sa()
	    InstructionDeTrie.Ra()

	    if InstructionDeTrie.Estranger(){
		   
			return
	    }
	    copy(InstructionDeTrie.A, copie)
		InstructionDeTrie.Gestes = InstructionDeTrie.Gestes[:debutGestes]
	    InstructionDeTrie.Rra()

		if InstructionDeTrie.Estranger(){
			
			return
		}

		copy(InstructionDeTrie.A, copie)
		InstructionDeTrie.Gestes = InstructionDeTrie.Gestes[:debutGestes]
	} else{
		return
		}
}





