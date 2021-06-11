package main

import "fmt"

func reverse(x int) int {
	// Problem ask to return 0 if beyond 2^31 and -2^31
    max := 1 << 31
    // Number will be stored in res
    res := 0
    // We divide by 10 which will remove the last digit of x, we loop until
    // there are no more numbers to remove
    for x != 0 {
    	// Modules of x returns the last digit of x in order to append it to our
    	// result we must multiple our result by 10 first then we can simply add
        res = res*10 + x % 10
        // Remove last digit of x
        x /= 10
        // Note: if x is negative, res will also be negative
    }
    // Check if result is not beyond max or min
    if res > max || res < -max {
        return 0
    }
    // Return result
    return res
}

func main() {
	fmt.Println(reverse(-123))
}