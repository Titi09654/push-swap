package TricinqNombres

import ("testing"
"pushSwap/InstructionDeTrie")

// Pour ce test vous pouvez mettre des nombres allant de trois à cinq. 
// Au delà le test ne fonctionnera pas car cette fonction est fait pour ranger au plus cinq nombres


func TestEstranger(t *testing.T){


    entrer := []int{8, 6, 1, 9}
    attendu := []int{1, 6, 8, 9}

    InstructionDeTrie.A = entrer
    TricinqNombres()

    for i := 0 ; i <len(entrer) ; i++{
	    if InstructionDeTrie.A[i] != attendu[i]{
		    t.Errorf("attendu %d, obtenu %d", InstructionDeTrie.A[i], attendu[i])
	    }
    }

}