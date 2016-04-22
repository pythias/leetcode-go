package leetcode

import "bytes"

func reverseString(s string) string {
    size := len(s)
    if size <= 1 {
        return s
    }

    var buffer bytes.Buffer
    for i := size - 1; i >= 0; i-- {
        buffer.WriteByte(s[i])
    }

    return buffer.String()
}
