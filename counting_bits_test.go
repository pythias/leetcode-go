package leetcode

import "testing"

func TestCountBits(t *testing.T) {
    cases := []struct {
        in int;
        want []int
    }{
        {1, []int {0, 1}},
        {7, []int {0, 1, 1, 2, 1, 2, 2, 3}},
        {8, []int {0, 1, 1, 2, 1, 2, 2, 3, 1}},
        {9, []int {0, 1, 1, 2, 1, 2, 2, 3, 1, 2}},
    }
    for _, c := range cases {
        got := countBits(c.in)
        if IntArrayEquals(got, c.want) == false {
            t.Errorf("countBits(%d) == %s, want %s", c.in, got, c.want)
        }
    }
}

func IntArrayEquals(a []int, b []int) bool {
    if len(a) != len(b) {
        return false
    }

    for i, v := range a {
        if v != b[i] {
            return false
        }
    }

    return true
}