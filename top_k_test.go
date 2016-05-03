package leetcode

import "testing"

func TestTopKFrequent(t *testing.T) {
    cases := []struct {
        nums []int;
        k int; 
        want []int
    }{
        {[]int {1, 2, 3, 1, 2, 1}, 2, []int {1, 2}},
        {[]int {1, 2, 1, 3}, 2, []int {1, 2}},
        {[]int {1, 2, 3, 1}, 3, []int {1, 2, 3}},
    }
    for _, c := range cases {
        got := topKFrequent(c.nums, c.k)
        if intArrayEquals(got, c.want) == false {
            t.Errorf("topKFrequent('%s', '%d') == '%s', want '%s'", c.nums, c.k, got, c.want)
        }
    }
}