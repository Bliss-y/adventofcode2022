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
    s, i := reader()
    s2, _ := reader()
    s3, _ := reader()
    for ;i < len(content); {

        if(len(s) <1){
            break;
        }
        for _, c := range s {
            found := false
            for i, c2 := range s2 {
                if c == c2 {
                    s2[i] = '-'
                    for _, c3 := range s3 {
                        if c3 == c {
                            found = true;
                            break;
                        }
                    }
                    if(found) {
                        break;
                    }
                }
            }
            if found {
            if c < 'a' {
                errorsum += int(c - 'A') + 27
            } else {
                errorsum +=  int(c - 'a') + 1;
            }
            break;
            }
        }
    s, i = reader()
    s2, _ = reader()
    s3, _ = reader()
    }
    fmt.Print(errorsum)
}

