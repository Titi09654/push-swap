package InstructionDeTrie

import ("testing")



func TestSa(t *testing.T){


	entrer:= []int{2, 1}
	attendu := []int{1, 2}

	A = entrer
	Sa()

	for i:=0 ; i<len(entrer); i++{
		if A[i] != attendu[i]{
			t.Errorf("attendu %d, obtenu %d", attendu[i], A[i])
		} 
	}
}

func TestRa(t *testing.T){


	entrer:= []int{3, 1, 2}
	attendu := []int{1, 2, 3}

	A = entrer
	Ra()

	for i:=0 ; i<len(entrer); i++{
		if A[i] != attendu[i]{
			t.Errorf("attendu %d, obtenu %d", attendu[i], A[i])
		} 
	}
}

func TestRra(t *testing.T){


	entrer := []int{2, 3, 1}
	
	attendu := []int{1, 2, 3}
	

	A = entrer
	Rra()

	for i:=0 ; i<len(entrer); i++{
		if A[i] != attendu[i]{
			t.Errorf("attendu %d, obtenu %d", attendu[i], A[i])
		} 
	}
}

func TestPa(t *testing.T){


	entrer_A:= []int{2, 3}
	entrer_B := []int{1}

	attendu_A := []int{1, 2, 3}
	attendu_B := []int{}

	A, B = entrer_A, entrer_B
	Pa()


	if len(attendu_A) != len(A){
		t.Fatalf("Longueur de A incorrecte : attendu %d, obtenu %d", len(attendu_A), len(A))
	} 
	for i:=0 ; i<len(attendu_A); i++{
		if A[i] != attendu_A[i]{
            t.Errorf(" attendu %d, obtenu %d", attendu_A[i], A[i])
		}
			
		 
		if len(B) != len(attendu_B) {
			t.Errorf("La pile B devrait être vide.")
		} 
	}
}

func TestPb(t *testing.T){


	entrer_A:= []int{1, 2, 3}
	entrer_B := []int{}

	attendu_A := []int{2, 3}
	attendu_B := []int{1}

	A, B = entrer_A, entrer_B
	Pb()


	if len(attendu_A) != len(A){
		t.Fatalf("Longueur de A incorrecte : attendu %d, obtenu %d", len(attendu_A), len(A))
	} 
	for i:=0 ; i<len(attendu_A); i++{
		if A[i] != attendu_A[i]{
            t.Errorf(" attendu %d, obtenu %d", attendu_A[i], A[i])
		}
			
		 
		if len(B) != len(attendu_B) {
			t.Errorf("La pile B devrait être vide.")
		} 
	}
}

func TestRrr(t *testing.T){


	entrer_A:= []int{2, 3, 1}
	entrer_B := []int{2, 3, 1}

	attendu_A := []int{1, 2, 3}
	attendu_B := []int{1, 2, 3}

	A, B = entrer_A, entrer_B
	Rrr()

	for i:=0 ; i<len(entrer_A); i++{
		if A[i] != attendu_A[i]{
			t.Errorf("attendu_A %d, obtenu %d", attendu_A[i], A[i])
		} 
		if B[i] != attendu_B[i]{
			t.Errorf("attendu_B %d, obtenu %d", attendu_B[i], B[i])
		}
	}
}

func TestRb(t *testing.T){


	entrer:= []int{3, 1, 2}
	attendu := []int{1, 2, 3}

	B = entrer
	Rb()

	for i:=0 ; i<len(entrer); i++{
		if B[i] != attendu[i]{
			t.Errorf("attendu %d, obtenu %d", attendu[i], B[i])
		} 
	}
}

func TestSs(t *testing.T){

    entrer_A := []int{2, 1, 3}
	entrer_B := []int{2, 1, 3}

	attendu_A :=[]int{1, 2, 3}
	attendu_B := []int{1, 2, 3}

	A, B = entrer_A, entrer_B
	Ss()

	for i:=0; i<len(entrer_A); i++{
		if A[i] != attendu_A[i]{
			t.Errorf("attendu %d, obtenu %d", attendu_A[i], A[i])
		} 
	}
	for l:=0 ; l<len(entrer_B);l++{
		if B[l] != attendu_B[l]{
			t.Errorf("attendu %d, obtenu %d", attendu_B[l], B[l])
		}
	}
}

func TestSb(t *testing.T){

   entrer := []int{2, 1, 3}
   attendu := []int{1, 2, 3}

   B = entrer 
   Sb()

   for i:=0; i<len(entrer); i++{
	    if B[i] != attendu[i]{
			t.Errorf("attendu %d, obtenu %d", attendu[i], B[i])
		}
   }
}

func TestRr(t *testing.T){

	entrer_A:= []int{3, 1, 2}
	entrer_B := []int{3, 1, 2}

	attendu_A := []int{1, 2, 3}
	attendu_B := []int{1, 2, 3}

	A, B = entrer_A, entrer_B
	Rr()

	for i:=0 ; i<len(entrer_A); i++{
		if A[i] != attendu_A[i]{
			t.Errorf("attendu %d, obtenu %d", attendu_A[i], A[i])
		}
	}
	for k:=0; k<len(entrer_B); k++{
		if B[k] != attendu_B[k]{
			t.Errorf("attendu %d, obtenu %d", attendu_B[k], B[k])
		}
	}
}

func TestRrb(t *testing.T){


	entrer := []int{2, 3, 1}
	
	attendu := []int{1, 2, 3}
	

	B = entrer
	Rrb()

	for i:=0 ; i<len(entrer); i++{
		if B[i] != attendu[i]{
			t.Errorf("attendu %d, obtenu %d", attendu[i], B[i])
		} 
	}
}




