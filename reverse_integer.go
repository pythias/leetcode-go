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

// uint8       the set of all unsigned  8-bit integers (0 to 255)
// uint16      the set of all unsigned 16-bit integers (0 to 65535)
// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

// int8        the set of all signed  8-bit integers (-128 to 127)
// int16       the set of all signed 16-bit integers (-32768 to 32767)
// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)