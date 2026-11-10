package TriplusDeNombre

import ("pushSwap/InstructionDeTrie"

)


func TriplusDeNombre(){


n := len(InstructionDeTrie.A)

for i := 0; i < n; i++ {

	if InstructionDeTrie.A[0] >= 0 && InstructionDeTrie.A[0] <= 2 {
        InstructionDeTrie.Pb()
        InstructionDeTrie.Rb()
	}else if InstructionDeTrie.A[0] >= 0 && InstructionDeTrie.A[0] <= 4 {
        InstructionDeTrie.Pb()
    } else {
        InstructionDeTrie.Ra()
    }
}
n = len(InstructionDeTrie.A)

for i := 0; i < n; i++ {
        InstructionDeTrie.Pb()
    }

for len(InstructionDeTrie.B) != 0 {

GrandNombre:=InstructionDeTrie.B[0]
indexmax:=0

    for i:=1 ; i < len(InstructionDeTrie.B) ; i ++{
	    if  GrandNombre < InstructionDeTrie.B[i] {
		    GrandNombre = InstructionDeTrie.B[i]
		    indexmax = i
	    }
    }
	for indexmax != 0 {
		    if indexmax <= len(InstructionDeTrie.B)/2 {
			    InstructionDeTrie.Rb()
			    indexmax = (indexmax - 1 + len(InstructionDeTrie.B)) % len(InstructionDeTrie.B)
		    } else {
			    InstructionDeTrie.Rrb()
			    indexmax = (indexmax + 1) % len(InstructionDeTrie.B)
		    }
	    }
		InstructionDeTrie.Pa()
		
    }
}
	



