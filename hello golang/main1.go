package main

import "fmt"

func main() {
	var smallfloat float32
	fmt.Println(smallfloat)
	smallfloat = 23.56565
	fmt.Println(smallfloat)

	var bigfloat float64
	fmt.Println(bigfloat)
	bigfloat = 23456789.123456789
	fmt.Println(bigfloat)

	var mycomplex complex128
	mycomplex = 1 + 2i
	fmt.Println(mycomplex)

	var myRealPart , myImaginaryPart float64
	myRealPart = real(mycomplex)
	myImaginaryPart = imag(mycomplex)
	fmt.Println(myRealPart)
	fmt.Println(myImaginaryPart)
}
