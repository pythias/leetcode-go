package leetcode

import (
    "math"
    "strconv"
)

func IsPalindrome(x int) bool {
    s := strconv.Itoa(x)
    l := len(s)
    h := int(math.Floor(float64(l / 2)))

    for i := 0; i < h; i++ {
        if s[i] != s[l - i - 1] {
            return false
        }
    }

    return true
}