package main

//Imports:
import "fmt"
import "unicode"

//Main funcion:
func main() {
    //Variables:
    x := 5
    y := 8
    z := 10

    letter0 := 'a'          //This is micro sign.
    var letter1 byte = 'B'  //This is a macro sign.

    //Logic:
    logic0 := (x == y)
    logic1 := (x < y) && (y < z)

    fmt.Println("p = ", logic0)
    fmt.Println("q = ", logic1)

    fmt.Println(letter0, letter1)
    fmt.Println(unicode.ToUpper(letter0))
}
