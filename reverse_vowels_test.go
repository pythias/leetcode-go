package leetcode

import "testing"

func TestReverseVowels(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"leetcode", "leotcede"},
        {"Ai", "iA"},
        {"let it go", "lot it ge"},
    }
    for _, c := range cases {
        got := reverseVowels(c.in)
        if got != c.want {
            t.Errorf("reverseVowels('%s') == '%s', want '%s'", c.in, got, c.want)
        }
    }
}