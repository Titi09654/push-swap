package TriplusDeNombre

import ("testing"
"pushSwap/InstructionDeTrie")

func TestEsranger(t *testing.T){

   entrer :=[]int{4, 8, 9, 2, 6, 11, 76, 0, 1, 3, 5, 99}
   attendu := []int{0, 1, 2, 3, 4, 5, 6, 8, 9, 11, 76, 99}

   InstructionDeTrie.A = entrer
   TriplusDeNombre()

   for i := 0 ; i<len(entrer); i++{
	   if InstructionDeTrie.A[i] != attendu[i]{
		   t.Errorf("attendu %d, obtenu %d", attendu[i], InstructionDeTrie.A[i])
	   }
   }



}