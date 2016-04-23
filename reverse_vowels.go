package leetcode

import "strings"

func reverseVowels(s string) string {
    size := len(s)
    if size <= 1 {
        return s
    }

    left := 0
    right := size - 1
    bytes := make([]byte, size)

    for i := left; i <= right; i++ {
        left = i + 1;
        found := false

        if isVowels(string(s[i])) {
            for j := right; j >= left; j-- {

                right = j - 1
                if isVowels(string(s[j])) {
                    bytes[i] = s[j]
                    bytes[j] = s[i]
                    found = true
                    break
                } else {
                    bytes[j] = s[j]
                }
            }

            if found == false {
                bytes[i] = s[right]
            }
        } else {
            bytes[i] = s[i]
        }
    }

    return string(bytes)
}

func isVowels(s string) bool {
    s = strings.ToLower(s)
    return s == "a" || s == "e" || s == "i" || s == "o" || s == "u" 
}