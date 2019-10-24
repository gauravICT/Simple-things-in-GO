// User can input n integer to find out prime number between 1 to n

package main

import (
    "fmt"
    "math"
)

// finding value is prime or not
func IsPrime(value int) bool {
    for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}

func main() {
    var n int
    fmt.Scanf("%d", &n)
    for i := 1; i <= n; i++ {
        if IsPrime(i) {
            fmt.Printf("%v ", i)
        }
    }

}
