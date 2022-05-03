package main

//Imports:
import (
    "fmt";
    "math";
)

//Constants:
const MAX_LENGTH uint = 100;

//Type definitions:
type word struct {
    name string;
    id  int;
}

type buffer struct {
    n       uint;
    in      uint;
    out     uint;
    array   [MAX_LENGTH]word;
}

func main() {
    //Variables:
    var b buffer;
    w := word {
        name: "William",
        id: 0,
    }
    var w1 word;

    //Logic on a buffer:
    open(&b);
    deposit(&b, w);
    deposit(&b, w);
    //fetch(&b, &w1);
    prt(&b);
    fmt.Println(w1);
}

//Funcion that prints a buffer b:
func prt(b *buffer) {
    var i uint;
    for i = 0; i < b.n; i++ {
        fmt.Println(b.array[i]);
    }
}

//Function that initialize a buffer b, that is, it sets its zero values:
func open(b *buffer) {
    b.n = 0;
    b.in = 0;
    b.out = 0;
}

//Function that adds a word w to a buffer b:
func deposit(b *buffer, w word) {
    if(b.n < MAX_LENGTH) {
        b.n++;
        b.array[b.in] = w;
        b.in = uint(math.Mod(float64(b.out + 1), float64(MAX_LENGTH))); //This line makes the buffer a circular list.
    }
}

//Function that reads the available position of a buffer b to a word w:
func fetch(b *buffer, w *word) {
    if(b.n > 0) {
        b.n--;
        *w = b.array[b.out];
        b.out = uint(math.Mod(float64(b.out + 1), float64(MAX_LENGTH)));
    }
}
