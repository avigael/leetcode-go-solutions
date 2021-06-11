package main

import "fmt"
import "math"

func longestPalindrome(s string) string {
    // Store max palindrome
    var max string
    // Iterate through every index of string
    for i := range s {
        // Expand outwards from center, also check special case with odd length
        max = maxPalindrome(s, i, i, max)
        max = maxPalindrome(s, i, i+1, max)
    }
    return max
}

func maxPalindrome(s string, i, j int, max string) string {
    // Stores current longest string
    var sub string
    // Keep bounds and check if letters match
    for i >= 0 && j < len(s) && s[i] == s[j] {
        // If so, assign slice as sub
        sub = s[i:j+1]
        // Move i and j pointers
        i--
        j++
    }
    // Return/replace longest string
    if len(max) < len(sub) {
        return sub
    }
    return max
}

func main() {
    fmt.Println(longestPalindrome("babad"))
}