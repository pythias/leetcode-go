package leetcode

import "testing"

func TestIsPowerOfFour(t *testing.T) {
    cases := []struct {
        in int; 
        want bool
    }{
        {16, true},
        {8, false},
        {4, true},
        {1, true},
        {15, false},
    }
    for _, c := range cases {
        got := isPowerOfFour(c.in)
        if got != c.want {
            t.Errorf("isPowerOfFour('%d') == '%s', want '%s'", c.in, got, c.want)
        }
    }
}