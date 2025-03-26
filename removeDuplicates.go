func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    count := 1
    prevElement := nums[0]

    for i := 1; i < len(nums); i++ {
        if nums[i] != prevElement {
            nums[count] = nums[i]
            count++
            prevElement = nums[i]
        }
    }

    return count
}
