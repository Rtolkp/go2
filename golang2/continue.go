/*continue statement
package main
import "fmt"
func main(){
	var ctr=0
	for ctr<10{
		ctr++
		if(ctr==5){
			println("5 is skipped")
			continue
			println("This won't be printed too")
		}
		fmt.Printf("Number is %d\n",ctr)
	}
	fmt.Println("out of loop")
}*/


package main
import "fmt"
func main(){
	var i=0 
	var j=0
	outerfor:
	for i<3{
		j=0
		i++
		for j<3{
			j++
			fmt.Printf("i is %d and j is %d\n",i,j)
			if(i==2){
				fmt.Println("I am skipped")
				continue outerfor
			}
			fmt.Println("I am printed")
		}
	}

}