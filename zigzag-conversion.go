package main

import "fmt"

func convert(s string, numRows int) string {
    // A single row means the string does not change
    if numRows == 1 {
        return s
    }
    // i -> insert position for "ret"
    // n -> length of string
    // cycleLen -> number of characters skipped to create a row
    //             ex. if we have 3 rows we know there will be 3 characters down
    //                 and 3 diagonally (including start and end charcter) so
    //                 we take our row + diagonal and substract our start and
    //                 end to calculate how many characters we need to skip over
    //                 and we know row = diagonal and start = 1 character end =
    //                 1 charcter so 2*row - 1 -1 -> 2*row-2 <- cycleLen
    // ret -> array to build our string
    i, n, cycleLen := 0, len(s), 2 * numRows - 2
    ret := make([]byte, n)
    // Loop through the string
    for row := 0; row < numRows; row++ {
        // cur is the vertical charcter for a row, see diagram below
        for cur := 0; cur+row < n; cur += cycleLen {
            // Copy character to our array (these are only the verticals of our
            // zigzag)
            // A   G
            // B  F  <- A, B, C, D, & G are "verticals"
            // C E   <- E & F are "diagonals"
            // D
            ret[i] = s[cur+row]
            i++
            // After the first loop we need need additional logic which takes
            // care of the "diagonals" in our zigzag
            // current-cycle-position + cycleLen - row 
            if row != 0 && row != numRows-1 && cur+cycleLen-row < n {
                // Copy character to our array
                ret[i] = s[cur+cycleLen-row]
                i++
            }
        }
    }
    // Convert our zigzagged array to a string and return
    return string(ret)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}