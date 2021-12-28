package main

//Imports:
import "fmt"

func main() {
    //Variables:
    type currency string;
    type weekday int;

    const(
        Real currency = "Real"
        Pound = "Pound"
        Euro = "Euro"
    )

    const(
        Monday weekday = iota
        Tuesday = 2
        Wednesday = 3
        Thursday = 4
        Friday = 5
    )

    //Logic:
    var a currency;
    var b weekday;

    a = Pound;
    b = Tuesday;

    fmt.Println(a);
    fmt.Println(b);
}
