package main
import "fmt"

func main(){
	nums:= []int{1,2,4,5,6,7,8}
	n:= nums[len(nums)-1]
	Sum:=(n*(n+1))/2
    Total_Sum := 0;

	for i:=0;i<len(nums);i++{
		Total_Sum =Total_Sum+nums[i]
	}
    Miss_num:= Sum - Total_Sum
	fmt.Println(Miss_num)
}