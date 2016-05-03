package leetcode

import "sort"

func topKFrequent(nums []int, k int) []int {
    frequents := map[int]int{}
    for _, num := range nums {
        if _, exists := frequents[num]; exists {
            frequents[num]++
        } else {
            frequents[num] = 1
        }
    }

    keysMap := map[int][]int{}
    var sortCount []int
    for num, count := range frequents {
        keysMap[count] = append(keysMap[count], num)
    }
    for count, _ := range keysMap {
        sortCount = append(sortCount, count)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(sortCount)))

    values := make([]int, k)
    i := 0
    for _, count := range sortCount {
        for _, num := range keysMap[count] {
            values[i] = num
            i++
            if i == k {
                return values
            }
        }
    }

    return values
}