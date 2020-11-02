/*switch statement
package main
import "fmt"
func main(){
		var dayofWeek=1
		switch dayofWeek{
			case 1:
			fmt.Println("Today is Monday")
			case 2:
			fmt.Println("Today is Tuesday")
			case 3:
			fmt.Println("Today is Wednesday")
			case 4:
			fmt.Println("Today is Thursday")
			case 5:
			fmt.Println("Today is Friday")
			case 6:
			fmt.Println("Today is Saturday")
			case 7:
			fmt.Println("Today is Sunday")
		    default:
		    fmt.Println("Invalid WeekDay")	
			}
}*/


/*
package main
import "fmt"
func main(){
	var dayofWeek=6
		switch dayofWeek{
		case 1,2,3,4,5:
			fmt.Println("Its is a  WeekDay")
		case 6,7:
			fmt.Println("Its is a WeekEnd")
		default:
			fmt.Println("Invalid Day")
		}
	}*/



/*switch fallthrough
package main
import "fmt"
func main(){
	var x=2
	switch x{
	case 1:
	        fmt.Println("1")
			fallthrough
	case 2:
	        fmt.Println("2")
	        fallthrough
	case 3:
	        fmt.Println("3")
	        fallthrough
	case 4:
	        fmt.Println("4")   

	    }
	}*/

/*switch short term
package main
import "fmt"
func main(){
	switch dayofWeek:=5; dayofWeek{
	case 1,2,3,4,5:
		fmt.Println("It's a WeekDay")
	case 6,7:
		fmt.Println("It's WeekEnd")
	default:
		fmt.Println("Invalid Day")

	}
}*/

/*switch with no expression
package main
import "fmt"
func main(){
	var marks=55
	switch{
	case marks>60:
	fmt.Println("Passed with 'A' Grade")
    case marks<60 && marks>=50:
    fmt.Println("Passed with 'B' Grade")
    case marks<50 && marks>35:
    fmt.Println("Passed with 'c' Grade")
    default:
    	fmt.Println("You r fail")


	}

}*/
		
