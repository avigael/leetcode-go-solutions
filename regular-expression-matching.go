package main

import "fmt"

// Logic:
// We will have a matrix where the rows represent the regular expression
// and the columns represent the string we are checking.
// We have to check the case were an empty string is valid. So we must add
// an extra row and column to our matrix. First we'll check and add the values
// for the empty row. We go down the columns up until we reach the position of
// we're at in the string. If a column is true but the column before produced a
// negative result then the current column is also false.
// Example Matrix: s = aa, p = a*  (- = nil)
//   - a *
// - T F T <- nil == nil = true, a != nil, * == nil = true because * = 0 or more
// a   T T <- a == a = true, * == a = true because * = 0 or more, prev true ok
// a     T <- a == * = true because * = 0, prev true ok -> final result true

func isMatch(s string, p string) bool {
    // Store lengths for speed and create matrix for dynamic programming
    lenS := len(s)
    lenP := len(p)
    dp := make([][]bool, lenS+1)
    // Initialize matrix
    for i := range dp {
        dp[i] = make([]bool, lenP+1)
    }
    // empty = empty
    dp[0][0] = true
    // only need to add [1][1] for first column
    if lenP > 0 && lenS > 0 {
        // do first character of string and expression match or does the
        // expression contain a "." meaning any character
        dp[1][1] = s[0] == p[0] || p[0] == '.'
    }
    // Checks repeatable
    for i := 2; i <= lenP; i++ {
        dp[0][i] = dp[0][i-2] && p[i-1] == '*'
    }
    // See Logic
    for i := 1; i <= lenS; i += 1 {
        for j := 2; j <= lenP; j += 1 {
            if p[j-1] != '*' {
                dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
            } else {
                dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
            }
        }
    }
    // Last entry should tell us if string is valid
    return dp[lenS][lenP]
}

func main() {
    fmt.Println(isMatch("aa","a*"))
}