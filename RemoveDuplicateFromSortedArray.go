//6. Remove duplicate from sorted array.
func removeDuplicates(nums []int) int {
    n := len(nums)
    if n ==0{
        return 0
    }

    count,prev:=1,nums[0]
    for i:=1;i< n;i++{
        if nums[i]!=prev{
            prev=nums[i]
            nums[count]= nums[i]
            count++
        }
    }
    return count
}
