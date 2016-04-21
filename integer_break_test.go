package leetcode

import "testing"

func TestIntegerBreak(t *testing.T) {
    cases := []struct {
        in, want int
    }{
        {2, 1},
        {8, 18},
        {10, 36},
        {15, 243},
    }
    for _, c := range cases {
        got := integerBreak(c.in)
        if got != c.want {
            t.Errorf("integerBreak(%d) == %d, want %d", c.in, got, c.want)
        }
    }
}