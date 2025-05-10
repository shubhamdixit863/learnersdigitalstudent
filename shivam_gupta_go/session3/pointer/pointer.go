package pointer
import "fmt"

func PointerAccess(){
	c:=9
  d:=&c
	fmt.Println(d)
	fmt.Println(*d)
}