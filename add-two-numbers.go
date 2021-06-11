package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // Create Linked List
    result := ListNode{}
    // x=Output, p=L1 traversal, q=L2 traversal
    x, p, q := &result, l1, l2
    // Carry for nums > 9, val = total (l1[i] + l2[i])
    carry, val := 0, 0
    // Loop until end of list, if carry != 0 then our output will be 1 node
    // longer than our lists (Ex. 999 + 999 = 1998)
    for p != nil || q != nil || carry != 0 {
        // Add previous carry to our total
        val = carry
        // Check if L1 has ended
        if p != nil {
            // Add L1[i] to our total
            val += p.Val
            // Go to next node for L1
            p = p.Next
        }
        // Check if L2 has ended
        if q != nil {
            // Add L2[i] to our total
            val += q.Val
            // Go to next node for L2
            q = q.Next
        }
        // Determine if carry is needed for next loop
        carry = val / 10
        // Add computed value to list
        x.Next = &ListNode{
            Val:  val % 10,
            Next: nil,
        }
        // Move to next node
        x = x.Next
    }
    // Return result
    return result.Next
}

func main() {
    l1 := ListNode{
        Val: 2,
        Next: &ListNode{
            Val: 4,
            Next: &ListNode{
                Val:  3,
                Next: nil,
            },
        },
    }

    l2 := ListNode{
        Val: 5,
        Next: &ListNode{
            Val: 6,
            Next: &ListNode{
                Val:  4,
                Next: nil,
            },
        },
    }

    res := addTwoNumbers(&l1, &l2)
    fmt.Print("[")
    for res != nil {
        fmt.Printf(" %v ", res.Val)
        res = res.Next
    }
    fmt.Print("]\n")
}
