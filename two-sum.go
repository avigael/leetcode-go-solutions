package main

import "fmt"

func twoSums(nums []int, target int) []int {
    // Create/Initialized Map[Key]Value
    m := make(map[int]int, len(nums))
    // Get index & value in nums
    for k, v := range nums {
        // Logic: If you want to get 2 int A & B that sum to equal X
        //        we know A + B = X has to be true. So X-B = A and
        //        X-A = B. With map we basically add all the numbers
        //        to the map with their index as their value. Then
        //        we know if we have A we just need to check if X-A
        //        exists in the map and make this our B.

        // A=v, X=target, if X-A exists then we have our B
        i, ok := m[target - v]

        // Yes?
        if ok {
            // Return indexes of current and matching element
            return []int{i, k}
        }
        // No? Add num to map with it's index in the array as its value
        m[v] = k
    }
    // Nothing found
    return []int{0, 0}
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    fmt.Println(twoSums(nums, target))
}