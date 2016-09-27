package leetcode

import "testing"

func TestReverse(t *testing.T) {
    cases := []struct {
        in string, 
        want int
    }{
        {" 0123 ", 123},
        {"-0123", -123},
        {"-12a3", -12},
        {"a123", 0},
        {"2147483647112", 2147483647},
        {"-2147483647112", -2147483648},
    }
    for _, c := range cases {
        got := myAtoi(c.in)
        if got != c.want {
            t.Errorf("myAtoi(%s) == %d, want %d", c.in, got, c.want)
        }
    }
}