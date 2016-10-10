package leetcode

import "bytes"

func longestCommonPrefix(strs []string) string {
    minSize := -1
    for _, str := range strs {
        if len(str) < minSize || minSize == -1 {
            minSize = len(str)
        }
    }

    var buffer bytes.Buffer
    for i := 0; i < minSize; i++ {
        var b byte
        b = 0
        for _, str := range strs {
            if b > 0 {
                if b != str[i] {
                    return buffer.String()
                }
            } else {
                b = str[i]
            }
        }

        buffer.WriteByte(b)
    }

    return buffer.String()
}
