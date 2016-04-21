package leetcode

func integerBreak(n int) int {
    if n < 2 {
        return 0
    }
    
    if n == 2 {
        return 1 * 1
    }

    if n == 3 {
        return 2 * 1
    }
    
    m := n % 3

    switch {
    case m == 0:
        s := 1
        for i := 0; i < n / 3; i++ {
            s *= 3
        }
        return s
        
    case m == 1:
        s := 4
        for i := 0; i < (n - 4) / 3; i++ {
            s *= 3
        }
        return s

    case m == 2:
        s := 2
        for i := 0; i < (n - 2) / 3; i++ {
            s *= 3
        }
        return s
    }

    return 0
}
