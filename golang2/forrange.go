/*forRange
package main
import "fmt"
func main(){
	langs:=[]string{"Go","C","C++","Java"}
	for r,k:=range langs{
		fmt.Println(r,k)
	}
}*/

/*for range with strings
package main
import "fmt"
func main(){
	for i,s:=range "Kanna"{
		fmt.Printf("%U represents %c and it is at position %d\n",s,s,i)
	}
}*/

/*Go range with map
package main
import "fmt"
func main(){
	fruits:=map[string]string{"A":"APPLE","B":"BANANA","C":"CHERRY"}
	for key,value:=range fruits{
		fmt.Printf("%s -> %s\n",key,value)
	}
	for key:=range fruits{
		fmt.Println("key:",key)
	}
}*/

/*Go range with channel
package main
import "fmt"
func main(){
	ch:=make(chan string)
	go func(){
		ch<-"k"
		ch<-"a"
		ch<-"n"
		ch<-"n"
		ch<-"a"
		close(ch)
	}()
		for  n :=range ch{
		fmt.Println(n)
	}
}*/




