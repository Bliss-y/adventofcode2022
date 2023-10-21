package main

import (
	"fmt"
	"os"
)

func LineReader(x []byte) func() ([]byte,int){
    content := x
    currentIndex := 0;
    return func()([]byte, int){
    if(currentIndex >= len(content)){
        return nil, currentIndex
    }
    current_string := make([]byte, 0, 20) 
    for ;content[currentIndex] != '\r' && currentIndex != len(content); currentIndex++{
       current_string = append(current_string, content[currentIndex]);
    }
    currentIndex+=2
    return current_string, currentIndex;
    }
}

func main() {
    inputfile := "input.txt"
    content, err := os.ReadFile(inputfile)
    if err != nil{
        fmt.Println("err reading file");
        return;
    }
    reader := LineReader(content);
    errorsum := 0
    for s, i := reader(); i < len(content); s,i = reader() {

        if(len(s) <1){
            break;
        }
        seperator := len(s)/2
        s1 := s[:seperator]
        s2 := s[seperator:]
        for _, c := range s1 {
            found := false
            for i, c2 := range s2 {
                if c == c2 {
                    s2[i] = '-'
                    found =true 
                }
            }
            if found {
            fmt.Printf("found %v %c \n", c, c);
            if c < 'a' {
                errorsum += int(c - 'A') + 27
            } else {
                errorsum +=  int(c - 'a') + 1;
            }
        }
        }
    }
    fmt.Print(errorsum)
}

