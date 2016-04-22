package leetcode

import "testing"

func TestReverseString(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"a", "a"},
        {"hello", "olleh"},
    }
    for _, c := range cases {
        got := reverseString(c.in)
        if got != c.want {
            t.Errorf("reverseString(%s) == %s, want %s", c.in, got, c.want)
        }
    }
}