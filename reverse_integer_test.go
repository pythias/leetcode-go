package leetcode

import "testing"

func TestReverse(t *testing.T) {
    cases := []struct {
        in, want int
    }{
        {123, 321},
        {-123, -321},
    }
    for _, c := range cases {
        got := reverse(c.in)
        if got != c.want {
            t.Errorf("reverseString(%d) == %d, want %d", c.in, got, c.want)
        }
    }
}