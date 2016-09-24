package leetcode

import "testing"

func TestCconvert(t *testing.T) {
    cases := []struct {
        s string;
        numRows int; 
        want string
    }{
        {"PAYPALISHIRING", 1, "PAYPALISHIRING"},
        {"PAYPALISHIRING", 2, "PYAIHRNAPLSIIG"},
        {"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
        {"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
        {"PAYPALISHIRING", 5, "PHASIYIRPLIGAN"},
        {"PAYPALISHIRING", 15, "PAYPALISHIRING"},
    }
    for _, c := range cases {
        got := convert(c.s, c.numRows)
        if got != c.want {
            t.Errorf("convert('%s', '%d') == '%s', want '%s'", c.s, c.numRows, got, c.want)
        }
    }
}