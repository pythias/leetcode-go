package leetcode

func countBits(num int) []int {
    if num < 0 {
        return []int {}
    }
// 0000
// 0001
// 0010
// 0011
// 0100
// 0101
// 0110
// 0111
// 1000

// 0 1 1 2
// 1 2 2 3
// 1 2 2 3
// 2 3 3 4
// 1 2 2 3
// 2 3 3 4
// 2 3 3 4
// 3 4 4 5
    result := make([]int, num + 1)
    result[0] = 0
    for i := 1; i < num + 1; i++ {
        result[i] = result[i >> 1] + (i & 1) 
    }

    return result
}
