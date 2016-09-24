package leetcode

import "bytes"

func convert(s string, numRows int) string {
    size := len(s)
    if size <= numRows || numRows <= 1 {
        return s
    }

    var m [][]byte
    for i := 0; i < numRows; i++ {
        var line []byte
        m = append(m, line)
    }

    row := 0
    col := 0
    for i := 0; i < size; i++ {
        if (col) % (numRows - 1) == 0 {
            m[row] = append(m[row], s[i])
        } else {
            if (col + row) % (numRows - 1) == 0 {
                m[row] = append(m[row], s[i])
            } else {
                m[row] = append(m[row], 0)
                i--
            }
        }
        
        row++
        if row == numRows {
            row = 0
            col++
        }
    }

    var buffer bytes.Buffer
    for i := 0; i < len(m); i++ {
        l := m[i]
        for j := 0; j < len(l); j++ {
            if l[j] > 0 {
                buffer.WriteByte(l[j])
            }
        }
    }

    return buffer.String()
}