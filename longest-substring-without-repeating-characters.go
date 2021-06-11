package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// Define variables
    sub := []rune{}
    max := 0
    // Iterate through string
    for _, c := range s {
    	// Check if letter already exists in substring
        if i, ok := find(sub, c); ok {
        	// If so resize to exclude said repeat
            sub = sub[i+1:]
        }
        // Add current letter to substring
        sub = append(sub, c)
        // Check size and record if larger than previous
        if len(sub) > max {
            max = len(sub)
        }
    }
    // Return largest size of substring at any point
    return max
}

// Helper function to find location of existing letters
func find(s []rune, x rune) (int, bool) {
	// Iterate through slice
    for i, c := range s {
    	// Check if letters match
        if c == x {
        	// Return index and ok
            return i, true
        }
    }
    // No match found
    return -1, false
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}