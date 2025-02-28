package main

import "fmt"

func reverseString( str string) string {
	runes := []rune(str) 
	str2 := make([]rune, len(runes)) 

	copy(str2, runes) 
 for i:=0;i<len(str2)/2;i++{
	str2[i],str2[len(str)-1-i]=str2[len(str)-1-i],str2[i]
 }
 return string(str2)
}
func main(){
 fmt.Println(reverseString("hello"))
}