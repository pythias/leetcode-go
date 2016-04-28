package leetcode

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if _, prs := m[nums[i]]; prs {
            return []int {m[nums[i]], i}
        } else {
            m[target - nums[i]] = i
        }
    }

    return []int {}
}
