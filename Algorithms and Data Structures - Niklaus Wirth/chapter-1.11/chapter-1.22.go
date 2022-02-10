package main

//Imports:
import "fmt";

//Constants:
var MAX_LENGTH int = 100;

//Type definitions:
type word struct {
    name string;
    id  int;
}

type sequence struct {
    pos     uint;
    length  uint;
    eof     bool;
    array   []word;
}

func main() {
    word1 := word {
        name:   "William",
        id:     1,
    }

    sequence1 := sequence {
        pos: 0,
        length: 0,
        eof: false,
        array: []word{word1},
    }

    fmt.Println(word1);
    fmt.Println(sequence1);
}
