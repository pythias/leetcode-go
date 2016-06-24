package leetcode

import "testing"

func TestIntersection(t *testing.T) {
    cases := []struct {
        nums1, nums2, want []int
    }{
        {[]int {1, 2, 2, 1}, []int {2, 2}, []int {2}},
        {[]int {1, 2, 2, 3}, []int {2, 3, 4}, []int {2, 3}},
    }
    for _, c := range cases {
        got := intersection(c.nums1, c.nums2)
        if intArrayEquals(got, c.want) == false {
            t.Errorf("intersection(%s, %s) == %s, want %s", c.nums1, c.nums2, got, c.want)
        }
    }
}