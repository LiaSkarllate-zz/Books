package main

//Imports:
import "fmt"

func main() {
    //Definitions:
    type line [3]int
    type matrix [3]line

    //Variables:
    var chars1 [10]byte     //In Go, arrays index type are integers and base type, specified.
    var chars2 [10]rune

    var table matrix = [3]line{[3]int{1, 2, 3}, [3]int{3, 4, 5}, [3]int{5, 6, 7}}
    var row line = [3]int{50, 40, 30}

    var sum int = 0         //sum := 0
    var max int = row[0]

    //Logic:
    for i := 0; i < 3; i++ {
        sum = row[i] + sum
    }

    for i := 0; i < 3; i++ {
        if(row[i] > max) {
            max = row[i]
        }
    }

    fmt.Println(chars1)
    fmt.Println(chars2)
    fmt.Println(table)
    fmt.Println(row)
    fmt.Println(sum)
    fmt.Println(max)
}
