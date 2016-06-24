package leetcode

func intersect(nums1 []int, nums2 []int) []int {
    nums := map[int]int {}
    for _, num := range nums1 {
        if _, exists := nums[num]; exists {
            nums[num]++
        } else {
            nums[num] = 1
        }
    }

    v := []int {}
    for _, num := range nums2 {
        if count, exists := nums[num]; exists {
            if count > 0 {
                v = append(v, num)
            }
            
            nums[num]--
        }
    }

    return v
}