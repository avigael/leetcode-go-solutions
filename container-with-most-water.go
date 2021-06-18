package main

import "fmt"

func maxArea(height []int) int {
    // Variables for storing current area and largest area calculated
    max, cur := 0, 0
    // To get the largest area we need the largest height and length so we start
    // at the ends of the array and work our way in
    left, right := 0, len(height) - 1
    // Loop until while we work our way inwards
    for left < right {
        // Since the problem involves water we want the shorter of the two
        // heights to calculate our area
        if height[left] > height[right] {
            // area = height * length
            cur = height[right] * (right - left);
            // Move inwards from right side
            right--
        } else {
            // area = height * length
            cur = height[left] * (right - left);
            // Move inwards from left side
            left++
        }
        // If the current calculation is larger than our current largest area we
        // replace our max with our current calculation
        if cur > max {
            max = cur
        }
    }
    // Return the largest area we calculated
    return max
}

func main() {
    height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
    fmt.Println(maxArea(height))
}
