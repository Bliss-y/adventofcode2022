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
    uniquecounter :=0
    marker := -1
    for i,c := range content {
        for j:=0; j<uniquecounter; j++{
            //fmt.Printf("%c %v\n",c, uniquecounter)
            //fmt.Printf("%c %d", c , i)
           if c == content[i - j - 1] {
            uniquecounter = j;
            break;
           } 
        }
        uniquecounter += 1;
        if(uniquecounter == 4){
           marker = i+1; 
        break;
        }

    }
    println(marker)  

}
