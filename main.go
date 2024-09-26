package main
import "fmt"

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    maxNum := max(10, 20)
    fmt.Println("Maximum is:", maxNum)
}