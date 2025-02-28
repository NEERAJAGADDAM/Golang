package main
import "fmt"

func trailingzeros(n int) int {
	count:=0
	for n>=5{
	n/=5
	count+=n
	}
	return count
}
func main(){
   n:=10
   fmt.Println("NO.Of Trailing zeros are in", n,  ":", trailingzeros(n))
}