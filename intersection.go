package leetcode

func intersection(nums1 []int, nums2 []int) []int {
    nums := map[int]int {}
    for _, num := range nums1 {
        nums[num] = 1
    }

    v := []int {}
    for _, num := range nums2 {
        if _, exists := nums[num]; exists {
            v = append(v, num)
            delete(nums, num)
        }
    }

    return v
}