package main

import (
	"fmt"
	"os"
    "strings"
    "strconv"
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
    num := 0;
    for s, i := reader(); i < len(content); s, i = reader(){
        if(len(s) <1) {
            break;
        }
        pairs := strings.Split(string(s), ",")
        first := strings.Split(pairs[0], "-")
        second := strings.Split(pairs[1], "-")
        ff,_ := strconv.Atoi(first[0])
        fs,_ := strconv.Atoi(first[1])
        sf,_ := strconv.Atoi(second[0])
        ss,_ := strconv.Atoi(second[1])
        if ff <= sf && fs >= ss{
            num++;
        } else if sf <= ff && ss >= fs{
            num ++ 
        }
    }
    fmt.Println(num)

}
