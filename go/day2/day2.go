package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {
    inputfile := "input.txt"
    content, err := os.ReadFile(inputfile)
    if err != nil{
        fmt.Println("err reading file");
        return;
    }

    highest := 0;
    highest_next :=0;
    highest_last :=0;
    current_score := 0;
    s := [10]byte{}
    current_string := s[:0]
    for i:=0; i<len(content); i++{
       if(content[i] == '\r' || i == len(content)){
          if(len(current_string) == 0) {
            if(current_score > highest){
                highest_last = highest_next;
                highest_next = highest;
                highest = current_score;
            } else if(current_score > highest_next){
                highest_last = highest_next;
                highest_next = current_score;
            } else if(current_score > highest_last){
                highest_last = current_score;
            }
            current_score = 0;
          } else {
            x,err:= strconv.Atoi(string(current_string))
            if(err != nil){
                fmt.Printf("err: %v", err)
                return
            }
            current_score += x;
            current_string = s[:0]
          }
          i++;
          continue;
       }
       current_string = append(current_string, content[i]);
    }
    fmt.Println(highest + highest_next + highest_last);

}

