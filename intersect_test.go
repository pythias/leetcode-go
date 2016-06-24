package leetcode

import "testing"

func TestIntersect(t *testing.T) {
    cases := []struct {
        nums1, nums2, want []int
    }{
        {[]int {1}, []int {1, 1}, []int {1}},
        {[]int {1, 2, 2, 3}, []int {2, 2}, []int {2, 2}},
    }
    for _, c := range cases {
        got := intersect(c.nums1, c.nums2)
        if intArrayEquals(got, c.want) == false {
            t.Errorf("intersect(%s, %s) == %s, want %s", c.nums1, c.nums2, got, c.want)
        }
    }
}