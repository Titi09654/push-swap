package TritroisNombres

import ("testing"
"pushSwap/InstructionDeTrie")

// Pour ce test vous pouvez mettre des nombres allant de deux à trois. 
// Au delà le test ne fonctionnera pas car cette fonction est fait pour ranger au plus trois nombres


func TestEstranger(t *testing.T){

	entrer := []int{2, 4, 1}
	attendu := []int{1, 2, 4}

	InstructionDeTrie.A = entrer
	TritroisNombres()

	for i := 0 ; i < len(entrer); i++{
		if InstructionDeTrie.A[i] != attendu[i] {
			t.Errorf("attendu %d, obtenu %d", attendu[i], InstructionDeTrie.A[i])
		}
	}
}
