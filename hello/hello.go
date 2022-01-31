package main

import( 
    "fmt"
    "math/rand"
    "math"
)
// short version "(x, y int)"
func add(x int, y int) int{
    return x + y
}

func swap(x, y string) (string, string){
    return y, x
}

// Named return values
// implicitly returns both x and y
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

// Variables with initializers
var i, j int = 1, 2

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("My favorite number is", rand.Intn(10))
    fmt.Println("Now you have %g problems.\n", math.Sqrt(69))
    fmt.Println("PI:",math.Pi)
    fmt.Println("result of function add(x,y): ", add(42,13))
    a, b := swap("hello", "world")
    fmt.Println("swapped (x, y): ", a, b)
}
