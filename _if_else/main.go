package main 
import "fmt"
func main(){
	i:=24
	if i%2==0{
		fmt.Println("even")
	} else if i%2!=0 {
		fmt.Println("odd")
	}else{
		fmt.Println("not a number")
	}
}