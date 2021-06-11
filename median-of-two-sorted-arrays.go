package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // Variables
    var result float64
    i, j := 0, 0
    len1, len2 := len(nums1), len(nums2)
    center := (len1+len2)/2
    arr := []int{}

    // Loop
    for {
        // Compare elements of nums1 and nums2 (Also checks if either is empty)
        if (len1 != 0 && len2 != 0) && nums1[i] < nums2[j] {
            // Add nums1 to larger array and move i to next position
            arr = append(arr, nums1[i])
            i++
        } else if len2 != 0 {
            // Else add nums2 to array and move j to next position
            arr = append(arr, nums2[j])
            j++
        }
        // Check to see if the end of the array has been reached
        if i == len1 {
            // If end of nums1 has been reached add the rest of nums2 to arr
            arr = append(arr, nums2[j:]...)
            break
        } else if j == len2 {
            // If end of nums2 has been reached add the rest of nums1 to arr
            arr = append(arr, nums1[i:]...)
            break
        }
    }

    // Get center element (extra logic for evenly sized arrays)
    if (len1 + len2) % 2 == 0 {
        num := float64(arr[center - 1] + arr[center])
        result = float64(num/2)
    } else {
        result = float64(arr[center])
    }
    // Return result
    return result
}

func main() {
    nums1 := []int{1,4,5}
    nums2 := []int{2,3,6}
    fmt.Println(findMedianSortedArrays(nums1, nums2))
}