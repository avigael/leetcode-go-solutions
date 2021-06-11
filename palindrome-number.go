package main

import "fmt"

func isPalindrome(x int) bool {
	// Only positive numbers can be palindromes
    if x < 0 {
        return false
    }
    // Create copy of x and create b to store our flipped number
    a, b := x, 0
    // Remove last digit from a until there are no more digits
    for a > 0 {
    	// Add last digit of a to b which mirrors a
        b = b * 10 + a % 10
        // Remove last digit
        a /= 10
    }
    // Compare b (flipped number) to original
    return b == x
}

func main() {
    fmt.Println(isPalindrome(12321))
}