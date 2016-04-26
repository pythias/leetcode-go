package leetcode

func countBits(num int) []int {
    if num < 0 {
        return []int {}
    }

// 0 1 1 2
// 1 2 2 3
// 1 2 2 3
// 2 3 3 4
// 1 2 2 3
// 2 3 3 4
// 2 3 3 4
// 3 4 4 5
    table := []int {
        0, 1, 1, 2, 
        1, 2, 2, 3, 
        1, 2, 2, 3, 
        2, 3, 3, 4,
    }

    index := 0
    round := 0
    result := make([]int, num + 1)
    for i := 0; i < num + 1; i++ {
        index = i & 15
        result[i] = table[index] + round
        if index == 15 {
            round++
        }
    }

    return result
}
