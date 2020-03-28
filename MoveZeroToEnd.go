//4/ Move Zeroes to the end maininting relative orders of other elements

func moveZeroes(nums []int){
	count:=0
	n:=len(nums)
	for i:=0;i< n;i++{
		if nums[i]!=0{
			nums[count]=nums[i]
			count++
		}
	}
	for count < n{
		nums[count]=0
		count++
