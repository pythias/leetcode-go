package leetcode

import (
    "bytes"
    "strconv"
)

func reverse(x int) int {
    negative := false
    if x < 0 {
        negative = true
        x = -x
    }

    s := strconv.Itoa(x)
    size := len(s)

    var buffer bytes.Buffer
    if negative {
        buffer.WriteByte(45)
    }
    
    for i := size - 1; i >= 0; i-- {
        buffer.WriteByte(s[i])
    }

    num, err := strconv.ParseInt(buffer.String(), 10, 32)
    if err != nil {
        return 0
    }

    return int(num)
}