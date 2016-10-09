package leetcode

import (
    "strings"
)

func romanToInt(s string) int {
    s = strings.Trim(s, " ")
    size := len(s)
    if size == 0 {
        return 0
    }

    last := 0
    num := 0
    var maps = map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    for i := 0; i < size; i++ {
        p := s[i:(i+1)] 

        current, exist := maps[p]
        if exist == false {
            return num
        }

        if current > last && last != 0 {
            num = num - last + (current - last) 
            last = 0
        } else {
            num += current
            last = current
        }
    }

    return num
}