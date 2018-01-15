package main 
import (
	"fmt";
	"math/rand";
	"time"
)
type IRoll func( int) int 
func fixedRoll(size int) int {
	return size - 1
}
func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}
func rollTwo(size int ) (int ,int ){
	return dieRoll(size),dieRoll(size)
}
func main(){
	fmt.Println(dieRoll(6))
	fmt.Println("Hello ,Go ")
	funcs :=[]IRoll{
		fixedRoll,
		dieRoll,
	}
	for i,runfunc := range funcs{
		fmt.Printf("Haha ,the middleware is OK ,%i fun runs %d\r\n",i,runfunc(6))
	}
	r1,r2 := rollTwo(6)
	fmt.Printf("the rool resut is %d ,%d,%d",r1,r2,funcs[0](6))
}
