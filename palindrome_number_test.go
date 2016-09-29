package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
    cases := []struct {
        in int; 
        want bool
    }{
        {161, true},
        {8, true},
        {11, true},
        {21, false},
        {15, false},
    }
    for _, c := range cases {
        got := isPalindrome(c.in)
        if got != c.want {
            t.Errorf("isPalindrome('%d') == '%s', want '%s'", c.in, got, c.want)
        }
    }
}