package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
    cases := []struct {
        nums []int;
        target int; 
        want []int
    }{
        {[]int {2, 7, 11, 15}, 9, []int {0, 1}},
        {[]int {4, 10, 1, 7}, 11, []int {1, 2}},
        {[]int {4, 10, 1, 7}, 12, []int {}},
    }
    for _, c := range cases {
        got := twoSum(c.nums, c.target)
        if intArrayEquals(got, c.want) == false {
            t.Errorf("twoSum('%s', '%d') == '%s', want '%s'", c.nums, c.target, got, c.want)
        }
    }
}