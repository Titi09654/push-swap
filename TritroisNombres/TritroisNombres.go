package TritroisNombres

import ("pushSwap/InstructionDeTrie")

func TritroisNombres(){
	if len(InstructionDeTrie.A) == 3{
        if Estranger(){
		    
			return
	    } 
	    copie := make([]int,len(InstructionDeTrie.A))
	    copy(copie, InstructionDeTrie.A)
	    InstructionDeTrie.Sa()
	    InstructionDeTrie.Rra()

	    if Estranger(){
		    
			return
	    }
	    copy(InstructionDeTrie.A, copie)
		InstructionDeTrie.Gestes = nil
	    InstructionDeTrie.Sa()

	    if Estranger(){
		   
			return
	    }
	    copy(InstructionDeTrie.A, copie)
		InstructionDeTrie.Gestes = nil
	    InstructionDeTrie.Ra()
     
	    if Estranger(){
		   
			return
	    }
	    copy(InstructionDeTrie.A, copie)
		InstructionDeTrie.Gestes = nil
	    InstructionDeTrie.Sa()
	    InstructionDeTrie.Ra()

	    if Estranger(){
		   
			return
	    }
	    copy(InstructionDeTrie.A, copie)
		InstructionDeTrie.Gestes = nil
	    InstructionDeTrie.Rra()

		if Estranger(){
			
			return
		}

		copy(InstructionDeTrie.A, copie)
		InstructionDeTrie.Gestes = nil
	} else{
		return
		}
}


func Estranger()bool {
 for i:=0 ; i< len(InstructionDeTrie.A)-1; i++{
	if InstructionDeTrie.A[i] > InstructionDeTrie.A[i+1] {
		return false
	    }
    } 
	return true
}


