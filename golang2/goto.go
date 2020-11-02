package main	
import "fmt"
func main(){
	var age int
	election:
	fmt.Print("Enter you are age")
	fmt.Scanln(&age)
	if(age<=17){
		fmt.Print("u r not eligiable to vote.\n")
		goto election
	}else{
		fmt.Print("u r eligible to vote")
	}
}