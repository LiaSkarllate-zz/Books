package main

//Imports:
import "fmt";

//Constants:
const MAX_LENGTH uint = 100;

//Type definitions:
type word struct {
    name string;
    id  int;
}

type sequence struct {
    pos     uint;
    length  uint;
    eof     bool;
    array   [MAX_LENGTH]word;
}

func main() {
    //Variables:]
    var s sequence; 
    w := word {
        name: "William",
        id: 0,
    }
    var w1 word;

    //Logic on a sequence:
    prt(&s);
    open(&s);
    write(&s, w);
    write(&s, w);
    reset(&s);
    read(&s, &w1);
    prt(&s);
    fmt.Println(w1);
}

//Funcion that prints a sequence s:
func prt(s *sequence) {
    var i uint;
    for i = 0; i < s.length; i++ {
        fmt.Println(s.array[i]);
    }
}

//Function that initialize a sequence s, that is, it sets its zero values:
func open(s *sequence) {
    s.length = 0;
    s.pos = 0;
    s.eof = false;
}

//Function that adds a word w to a sequence s:
func write(s *sequence, w word) {
    if(s.pos < MAX_LENGTH) {
        s.array[s.pos] = w;
        s.pos++;
        s.length = s.pos;
    }
}

//Function that resets the current position of a sequence s:
func reset(s *sequence) {
    s.pos = 0;
    s.eof = false;
}

//Function that reads the current position of a sequence s to a word w:
func read(s *sequence, w *word) {
    if(s.pos == MAX_LENGTH) {
        s.eof = true;
    } else {
        *w = s.array[s.pos];
        s.pos++;
    }
}
