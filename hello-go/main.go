package main



import "fmt"



func main() {

    s1 := []int{1, 2, 3, 4}

    s2 := make([]int, len(s1))

    copy(s2, s1)



    s1[0] = 100

    fmt.Println(s1, s2)



    s1 = append(s1, 200)

    fmt.Println(s1, s2)

}