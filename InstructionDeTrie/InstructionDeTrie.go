package InstructionDeTrie

 
var A , B []int

func Sa() {
   if len(A) >= 2 {
	A[0], A[1] = A[1], A[0]
   }
   Noter("sa")
}



func Ra(){
    if len(A)>= 2 {
		A = append(A[1:], A[0])
	}
	Noter("ra")
}


func Rra(){
    if len(A)>=2{
	    fond := A[len(A)-1]
	    A = append([]int{fond}, A[:len(A)-1]...)
	}
	Noter("rra")
}


func Pa(){
    if len(B) > 0 {
      haut := B[0]
	  B = B[1:]
	  A= append([]int{haut}, A...)
	}
	Noter("pa")
}

func Pb(){
    if len(A) > 0 {
		haut := A[0]
		A = A[1:]
		B = append([]int{haut}, B...)
	}
	Noter("pb")
}

func Rrr(){
	if len(A) >= 2 {
        fond := A[len(A)-1] 
         A = append([]int{fond}, A[:len(A)-1]...)
	}
	if len(B) >=2 {
		fond := B[len(B)-1]
		B =append([]int{fond}, B[:len(B)-1]...)
	}
    Noter("rrr")
}

func Rb(){
    if len(B) >= 2 {
	   B = append(B[1:], B[0])
	}
	Noter("rb")
}

func Ss(){
    if len(A)>=2{
		A[0], A[1] = A[1], A[0]
	}
	if len(B) >= 2 {
	    B[0], B[1] = B[1], B[0]
   }
   Noter("ss")
}

func Sb(){
   if len(B) >= 2 {
	B[0], B[1] = B[1], B[0]
   }
   Noter("sb")
}

func Rr(){
    if len(A)>= 2 {
		A = append(A[1:], A[0])
	}
	if len(B) >= 2 {
	   B = append(B[1:], B[0])
	}
	Noter("rr")
}

func Rrb(){
    if len(B) >=2{
	    fond := B[len(B)-1]
	    B= append([]int{fond}, B[:len(B)-1]...)
	}
	Noter("rrb")
}



var Gestes []string

func Noter(instruction string)  {
	
    Gestes = append(Gestes, instruction)
}