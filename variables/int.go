package variables

import "fmt"

func MostrarInt (){
	var commonInt int 
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("int comun = ", commonInt)
	fmt.Println("int de 32 bits = ", intde32)
	fmt.Println("int de 64 bits = ", intde64)
}