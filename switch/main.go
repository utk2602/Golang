package main 
import "fmt"
func main() {
	i:=10
	switch {
	case i%2==0:
		fmt.Println("even")		
	case i%2!=0:
		fmt.Println("odd")
	default:
		fmt.Println("not a number")
	}	
}