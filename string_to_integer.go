package leetcode

import (
    "math"
    "strings"
)

func myAtoi(str string) int {
    str = strings.Trim(str, " ")
    size := len(str)
    if size == 0 {
        return 0
    }

    negative := false
    fix := 0
    num := 0.0
    first := str[0]
    if first == 45 {
        fix = 1
        negative = true
    }

    if first == 43 {
        fix = 1
    }

    nums := []int {}
    for i := fix; i < size; i++ {
        if str[i] < 48 || str[i] > 57 {
            break
        }

        nums = append(nums, int(str[i] - 48))
    }

    size = len(nums)
    for i := 0; i < size; i++ {
        num += float64(nums[i]) * math.Pow10(size - i - 1)
    }

    if negative {
        if num > float64(1<<31) {
            return -1 << 31
        }

        return -int(num)
    }

    if num > float64(1<<31 - 1) {
        return 1<<31 - 1
    }

    return int(num)
}