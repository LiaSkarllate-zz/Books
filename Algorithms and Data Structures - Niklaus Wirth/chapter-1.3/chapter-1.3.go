package main

//Imports:
import "fmt"

//Global definitions:
type currency string;
type weekday int;

const(
    Real currency = "Real"
    Pound = "Pound"
    Euro = "Euro"
)

const(
    Monday weekday = iota
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday
)

//Main function:
func main() {
    //Logic:
    var a currency
    var b weekday

    a = Pound
    b = Tuesday

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(b + 1)
}
